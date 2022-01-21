package server

import (
	"errors"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"time"
	"strings"

	"github.com/assetto-corsa-web/accweb/cfg"
	"github.com/sirupsen/logrus"
)

const (
	logDir             = "logs"
	logFilename        = "logs_"
	logTimeFormat      = "20060102_150405"
	logExt             = ".log"
	cfgDir             = "cfg"
	accServerFile      = "accServer.exe"
	serverLogDir       = "log"
	serverErrorLogDir  = "error"
	serverLogFile      = "server.log"
	serverErrorLogFile = "error.log"
)

func StartServer(id int) error {
	logrus.WithField("id", id).Info("Starting server instance...")
	server := GetServerById(id, true)

	if server == nil {
		logrus.Error("Server not found")
		return ServerNotFound
	}

	if server.PID != 0 {
		logrus.Error("Server running already")
		return errors.New("Server running")
	}

	logfile, err := createLogFile(server)

	if err != nil {
		logrus.WithError(err).Error("Error creating log file")
		return err
	}

	if err := copyServerFiles(server.Id); err != nil {
		logrus.WithError(err).Error("Error copying server files")
		return err
	}

	if err := copyCfgFiles(server.Id); err != nil {
		logrus.WithError(err).Error("Error copying configuration files")
		return err
	}

	command := "." + string(filepath.Separator) + cfg.Get().ACC.ServerExe
	args := make([]string, 0)

	if runtime.GOOS == "linux" {
		command = "wine"
		args = append(args, cfg.Get().ACC.ServerExe)
	}

	serverExecutionPath := filepath.Join(cfg.Get().ConfigPath, strconv.Itoa(server.Id))
	cmd := exec.Command(command, args...)
	cmd.Stdout = logfile
	cmd.Stderr = logfile
	cmd.Dir = serverExecutionPath
	logrus.Warn(serverExecutionPath)

	if err := cmd.Start(); err != nil {
		logrus.WithError(err).Error("Error starting server")
		return err
	}

	server.start(cmd)
	setServer(server)
	
	if runtime.GOOS == "windows" {
		if server.Configuration.Affinity != "" {
			logrus.Info("Parsing Affinity...")
			numCPU := runtime.NumCPU()
			affinityMask := make([]byte, numCPU)
			for i := range affinityMask {
				affinityMask[i] = '0'
			}
			
			cores := strings.Split(server.Configuration.Affinity, ",")
			for _, core := range cores {
				coreInt64, coreErr := strconv.ParseInt(core, 10, 0)
				coreInt := int(coreInt64)
				if coreErr == nil {
					if coreInt > 0 && coreInt <= numCPU {
						affinityMask[len(affinityMask)-int(coreInt)] = '1'
					} else {
						logrus.WithField("Core", core).Warning("Core out of Range")
					}
				} else {
					logrus.WithField("Core", core).WithError(coreErr).Error("Core Parse Error")
				}
			}
			logrus.Info("Parsed Affinity, generating Hex...")
			affinityMaskInt, affinityMaskErr := strconv.ParseInt(string(affinityMask), 2, 0)
			if affinityMaskErr == nil {
				affinityArgs := []string{"$Process = Get-Process -Id " + strconv.Itoa(server.PID) + "; $Process.ProcessorAffinity=" + strconv.FormatInt(affinityMaskInt, 10)}
				affinityCmd := exec.Command("PowerShell", affinityArgs...)
				affinityCmd.Stdout = logfile
				affinityCmd.Stderr = logfile
				affinityErr := affinityCmd.Start()
				if affinityErr != nil {
					logrus.WithError(affinityErr).Error("Error set affinity")
				} else {
					logrus.WithField("Command", affinityCmd.String()).Info("affinity set")
				}
			} else {
				logrus.WithError(affinityMaskErr).Error("AffinityMask Parse Error")
			}
		}

		logrus.WithField("Prio", server.Configuration.Priority).Info("Prio get")
		prioArgs := []string{"process", "where", "ProcessId=" + strconv.Itoa(server.PID), "call", "setpriority", strconv.Itoa(server.Configuration.Priority)}
		prioCmd := exec.Command("wmic", prioArgs...)
		prioCmd.Stdout = logfile
		prioCmd.Stderr = logfile
		prioErr := prioCmd.Start()
		if prioErr != nil {
			logrus.WithError(prioErr).Error("Error set priority")
		} else {
			logrus.WithField("PID", prioCmd.Process.Pid).Info("Prio set")
			logrus.WithField("Command", prioCmd.String()).Info("Prio set")
		}
		
		logrus.Info("Add Firewall Rules")
		fwTcpArgs := []string{"advfirewall", "firewall", "add", "rule", "name=\"ACCSERVER_" + strconv.Itoa(server.Id) + "\"", "dir=in", "action=allow", "protocol=TCP", "localport=" + strconv.Itoa(server.Configuration.TcpPort)}
		fwTcpCmd := exec.Command("netsh", fwTcpArgs...)
		fwTcpCmd.Stdout = logfile
		fwTcpCmd.Stderr = logfile
		fwTcpErr := fwTcpCmd.Start()
		if fwTcpErr != nil {
			logrus.WithError(fwTcpErr).Error("Error opening TCP Port")
		} else {
			logrus.Info("Opened TCP Port " + strconv.Itoa(server.Configuration.TcpPort))
			logrus.WithField("Command", fwTcpCmd.String()).Info("Opened TCP Port")
		}
		
		fwUdpArgs := []string{"advfirewall", "firewall", "add", "rule", "name=\"ACCSERVER_" + strconv.Itoa(server.Id) + "\"", "dir=in", "action=allow", "protocol=UDP", "localport=" + strconv.Itoa(server.Configuration.UdpPort)}
		fwUdpCmd := exec.Command("netsh", fwUdpArgs...)
		fwUdpCmd.Stdout = logfile
		fwUdpCmd.Stderr = logfile
		fwUdpErr := fwUdpCmd.Start()
		if fwUdpErr != nil {
			logrus.WithError(fwUdpErr).Error("Error opening UDP Port")
		} else {
			logrus.Info("Opened UDP Port " + strconv.Itoa(server.Configuration.UdpPort))
			logrus.WithField("Command", fwUdpCmd.String()).Info("Opened UDP Port")
		}
	}
	
	logrus.WithField("PID", server.PID).Info("Server started")
	go observeProcess(server, logfile)

	return nil
}

func createLogFile(server *ServerSettings) (*os.File, error) {
	dir, _, err := getConfigDirectoryAndID(server.Id)

	if err != nil {
		return nil, err
	}

	if err := os.MkdirAll(filepath.Join(dir, logDir), 0755); err != nil {
		return nil, err
	}

	var re = regexp.MustCompile(`[^a-zA-Z0-9_-]+`)
	serverName := re.ReplaceAllString(server.Settings.ServerName, "_")

	filename := logFilename + time.Now().Format(logTimeFormat) + "_" + strconv.Itoa(server.Id) + "_" + serverName + logExt
	logfile, err := os.Create(filepath.Join(dir, logDir, filename))

	return logfile, nil
}

func copyServerFiles(id int) error {
	accServerSource := filepath.Join(cfg.Get().ACC.ServerPath, accServerFile)
	accServerTarget := filepath.Join(cfg.Get().ConfigPath, strconv.Itoa(id), accServerFile)

	if err := copyFile(accServerSource, accServerTarget, 0755); err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Join(cfg.Get().ConfigPath, strconv.Itoa(id), serverLogDir), 0755); err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Join(cfg.Get().ConfigPath, strconv.Itoa(id), serverLogDir, serverErrorLogDir), 0755); err != nil {
		return err
	}

	if err := ioutil.WriteFile(filepath.Join(cfg.Get().ConfigPath, strconv.Itoa(id), serverLogDir, serverLogFile), nil, 0755); err != nil {
		return err
	}

	if err := ioutil.WriteFile(filepath.Join(cfg.Get().ConfigPath, strconv.Itoa(id), serverLogDir, serverErrorLogDir, serverErrorLogFile), nil, 0755); err != nil {
		return err
	}

	return nil
}

func copyCfgFiles(id int) error {
	sourceDir := filepath.Join(cfg.Get().ConfigPath, strconv.Itoa(id))
	targetDir := filepath.Join(cfg.Get().ConfigPath, strconv.Itoa(id), cfgDir)

	if err := os.MkdirAll(targetDir, 0777); err != nil {
		return err
	}

	if err := copyFile(filepath.Join(sourceDir, configurationJsonName), filepath.Join(targetDir, configurationJsonName), 0755); err != nil {
		return err
	}

	if err := copyFile(filepath.Join(sourceDir, settingsJsonName), filepath.Join(targetDir, settingsJsonName), 0755); err != nil {
		return err
	}

	if err := copyFile(filepath.Join(sourceDir, eventJsonName), filepath.Join(targetDir, eventJsonName), 0755); err != nil {
		return err
	}

	if err := copyFile(filepath.Join(sourceDir, eventRulesJsonName), filepath.Join(targetDir, eventRulesJsonName), 0755); err != nil {
		return err
	}

	if err := copyFile(filepath.Join(sourceDir, entrylistJsonName), filepath.Join(targetDir, entrylistJsonName), 0755); err != nil {
		return err
	}

	if err := copyFile(filepath.Join(sourceDir, bopJsonName), filepath.Join(targetDir, bopJsonName), 0755); err != nil {
		return err
	}

	if err := copyFile(filepath.Join(sourceDir, assistRulesJsonName), filepath.Join(targetDir, assistRulesJsonName), 0755); err != nil {
		return err
	}

	return nil
}

func copyFile(source, target string, perm os.FileMode) error {
	logrus.WithFields(logrus.Fields{"source": source, "target": target}).Debug("Copying file")
	sourceFile, err := ioutil.ReadFile(source)

	if err != nil {
		logrus.WithFields(logrus.Fields{"err": err, "file": source}).Error("Error reading source file")
		return err
	}

	if err := ioutil.WriteFile(target, sourceFile, perm); err != nil {
		logrus.WithFields(logrus.Fields{"err": err, "file": target}).Error("Error writing target file")
		return err
	}

	return nil
}

func observeProcess(server *ServerSettings, logfile *os.File) {
	// close log file when server stops
	defer func() {
		if err := logfile.Close(); err != nil {
			logrus.WithError(err).Error("Error closing server log file")
		}
	}()

	// wait for shutdown or crash
	if err := server.Cmd.Wait(); err != nil {
		exitErr, ok := err.(*exec.ExitError)

		if ok {
			logrus.WithField("err", exitErr.Error()).Error("Server stopped with an error")
		} else {
			logrus.WithError(err).Error("Error when server stopped")
		}
	}

	// reset PID and cmd so that server can be started again
	logrus.WithField("id", server.Id).Info("Server stopped")
	server.stop()
	setServer(server)
}

func StopServer(id int) error {
	logrus.WithField("id", id).Info("Stopping server instance...")
	server := GetServerById(id, true)

	if server == nil {
		return ServerNotFound
	}

	if server.PID == 0 {
		return nil
	}

	if err := server.Cmd.Process.Signal(os.Interrupt); err != nil {
		logrus.WithError(err).Error("Error stopping instance sending an interrupt, trying to kill it next")

		if err := server.Cmd.Process.Kill(); err != nil {
			logrus.WithError(err).Error("Error stopping instance sending kill signal")
			return err
		}
	}
	
	logfile, _ := createLogFile(server)
	
	if runtime.GOOS == "windows" {		
		logrus.Info("Remove Firewall Rules")
		fwTcpArgs := []string{"advfirewall", "firewall", "del", "rule", "name=\"ACCSERVER_" + strconv.Itoa(server.Id) + "\""}
		fwTcpCmd := exec.Command("netsh", fwTcpArgs...)
		fwTcpCmd.Stdout = logfile
		fwTcpCmd.Stderr = logfile
		fwTcpErr := fwTcpCmd.Start()
		if fwTcpErr != nil {
			logrus.WithError(fwTcpErr).Error("Error closing TCP Port")
		} else {
			logrus.WithField("Command", fwTcpCmd.String()).WithField("Port", server.Configuration.TcpPort).Info("Closed TCP Port")
		}
		
		fwUdpArgs := []string{"advfirewall", "firewall", "del", "rule", "name=\"ACCSERVER_" + strconv.Itoa(server.Id) + "\""}
		fwUdpCmd := exec.Command("netsh", fwUdpArgs...)
		fwUdpCmd.Stdout = logfile
		fwUdpCmd.Stderr = logfile
		fwUdpErr := fwUdpCmd.Start()
		if fwUdpErr != nil {
			logrus.WithError(fwUdpErr).Error("Error closing UDP Port")
		} else {
			logrus.Info("Closed UDP Port " + strconv.Itoa(server.Configuration.UdpPort))
			logrus.WithField("Command", fwUdpCmd.String()).Info("Closed UDP Port")
		}
	}

	return nil
}
