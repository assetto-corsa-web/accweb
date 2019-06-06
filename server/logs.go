package server

import (
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
)

const (
	maxLogSize = 256000
)

func GetServerLogs(id int) (string, error) {
	server := GetServerById(id, true)

	if server == nil {
		return "", ServerNotFound
	}

	dir, _, err := getConfigDirectoryAndID(server.Id)

	if err != nil {
		return "", err
	}

	logDirPath := filepath.Join(dir, logDir)
	logDir, err := ioutil.ReadDir(logDirPath)

	if err != nil {
		return "", err
	}

	filename, err := findLatestLogfileName(logDirPath, logDir)

	if err != nil {
		return "", err
	}

	if filename == "" {
		return "", nil
	}

	logrus.WithField("filename", filename).Debug("Loading server log file")
	file, err := os.Open(filepath.Join(logDirPath, filename))

	if err != nil {
		return "", err
	}

	defer func() {
		if err := file.Close(); err != nil {
			logrus.WithError(err).Error("Error closing log file")
		}
	}()
	start, err := getLogStartBytes(file)

	if err != nil {
		return "", err
	}

	content, err := getLogContent(file, start)

	if err != nil {
		return "", err
	}

	return content, nil
}

func findLatestLogfileName(logDirPath string, logDir []os.FileInfo) (string, error) {
	logFilenames := make([]string, 0)

	for _, file := range logDir {
		if !file.IsDir() && filepath.Ext(filepath.Join(logDirPath, file.Name())) == logExt {
			logFilenames = append(logFilenames, file.Name())
		}
	}

	if len(logFilenames) == 0 {
		return "", nil
	}

	sort.Strings(logFilenames)
	return logFilenames[len(logFilenames)-1], nil
}

func getLogStartBytes(file *os.File) (int64, error) {
	info, err := file.Stat()

	if err != nil {
		return 0, err
	}

	start := info.Size() - maxLogSize + 1

	if start < 0 {
		start = 0
	}

	return start, nil
}

func getLogContent(file *os.File, start int64) (string, error) {
	buffer := make([]byte, maxLogSize)
	n, err := file.ReadAt(buffer, start)

	if err != nil && err != io.EOF {
		return "", err
	}

	if n == 0 {
		return "", nil
	}

	return string(buffer[:n-1]), nil
}
