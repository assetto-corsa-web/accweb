package server

import (
	"errors"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"time"
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

	command := "." + string(filepath.Separator) + os.Getenv("ACCWEB_SERVER_EXE")
	args := make([]string, 0)

	if runtime.GOOS == "linux" {
		command = "wine"
		args = append(args, os.Getenv("ACCWEB_SERVER_EXE"))
	}

	serverExecutionPath := filepath.Join(os.Getenv("ACCWEB_CONFIG_PATH"), strconv.Itoa(server.Id))
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

	filename := logFilename + time.Now().Format(logTimeFormat) + "_" + strconv.Itoa(server.Id) + "_" + server.Settings.ServerName + logExt
	logfile, err := os.Create(filepath.Join(dir, logDir, filename))

	return logfile, nil
}

func copyServerFiles(id int) error {
	accServerSource := filepath.Join(os.Getenv("ACCWEB_SERVER_DIR"), accServerFile)
	accServerTarget := filepath.Join(os.Getenv("ACCWEB_CONFIG_PATH"), strconv.Itoa(id), accServerFile)

	if err := copyFile(accServerSource, accServerTarget, 0755); err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Join(os.Getenv("ACCWEB_CONFIG_PATH"), strconv.Itoa(id), serverLogDir), 0755); err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Join(os.Getenv("ACCWEB_CONFIG_PATH"), strconv.Itoa(id), serverLogDir, serverErrorLogDir), 0755); err != nil {
		return err
	}

	if err := ioutil.WriteFile(filepath.Join(os.Getenv("ACCWEB_CONFIG_PATH"), strconv.Itoa(id), serverLogDir, serverLogFile), nil, 0755); err != nil {
		return err
	}

	if err := ioutil.WriteFile(filepath.Join(os.Getenv("ACCWEB_CONFIG_PATH"), strconv.Itoa(id), serverLogDir, serverErrorLogDir, serverErrorLogFile), nil, 0755); err != nil {
		return err
	}

	return nil
}

func copyCfgFiles(id int) error {
	sourceDir := filepath.Join(os.Getenv("ACCWEB_CONFIG_PATH"), strconv.Itoa(id))
	targetDir := filepath.Join(os.Getenv("ACCWEB_CONFIG_PATH"), strconv.Itoa(id), cfgDir)

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

	if err := copyFile(filepath.Join(sourceDir, entrylistJsonName), filepath.Join(targetDir, entrylistJsonName), 0755); err != nil {
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

	return nil
}
