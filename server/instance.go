package server

import (
	"errors"
	"github.com/sirupsen/logrus"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"time"
)

const (
	logDir        = "logs"
	logFilename   = "logs_"
	logTimeFormat = "20060102_150405"
)

// TODO pass config files to server on startup
func StartServer(id int) error {
	logrus.WithField("id", id).Info("Starting server instance...")
	server := GetServerById(id)

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

	cmd := exec.Command(filepath.Join(os.Getenv("ACCWEB_SERVER_DIR"), os.Getenv("ACCWEB_SERVER_EXE")))
	cmd.Stdout = logfile
	cmd.Stderr = logfile
	//cmd.Dir = os.Getenv("ACCWEB_SERVER_DIR")

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

	if err := os.MkdirAll(filepath.Join(dir, logDir), 0770); err != nil {
		return nil, err
	}

	filename := logFilename + time.Now().Format(logTimeFormat) + "_" + strconv.Itoa(server.Id) + "_" + server.Settings.ServerName + ".log"
	logfile, err := os.Create(filepath.Join(dir, logDir, filename))

	return logfile, nil
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
	server := GetServerById(id)

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
