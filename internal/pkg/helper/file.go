package helper

import (
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

var (
	utf16Encoding = unicode.UTF16(unicode.LittleEndian, unicode.UseBOM)

	// ErrFileNotFound file not found
	ErrFileNotFound = errors.New("file not found")
	// ErrInvalidJsonFileFormat Contains all errors reported during json decode.
	ErrInvalidJsonFileFormat = errors.New("invalid json file")
)

func LoadFromPath(baseDir, filename string, obj interface{}) error {
	path := filepath.Join(baseDir, filename)
	f, err := os.Open(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return WrapErrors(ErrFileNotFound, err)
		}

		return err
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	r := transform.NewReader(f, utf16Encoding.NewDecoder().Transformer)

	if err := json.NewDecoder(r).Decode(obj); err != nil {
		if _, errSeek := f.Seek(0, io.SeekStart); errSeek != nil {
			return errSeek
		}

		// trying decode non ut16 content
		if err2 := json.NewDecoder(f).Decode(obj); err2 != nil {
			return WrapErrors(ErrInvalidJsonFileFormat, err2, err)
		}
	}

	return nil
}

func SaveToPath(baseDir, filename string, obj interface{}) error {
	data, err := json.Marshal(obj)
	if err != nil {
		return err
	}

	encodedData, err := utf16Encoding.NewEncoder().Bytes(data)
	if err != nil {
		return err
	}

	path := filepath.Join(baseDir, filename)
	if err := ioutil.WriteFile(path, encodedData, 0655); err != nil {
		return err
	}

	return nil
}

func CheckMd5Sum(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return "", WrapErrors(ErrFileNotFound, err)
		}

		return "", err
	}

	hash := md5.New()
	if _, err := hash.Write(data); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

func Exists(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	}

	return true
}

func Copy(srcFile, dstFile string) error {
	out, err := os.Create(dstFile)
	if err != nil {
		return err
	}

	defer out.Close()

	in, err := os.Open(srcFile)
	defer func() { _ = in.Close() }()

	if err != nil {
		return err
	}

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}

	return nil
}

// CreateIfNotExists creates a new directory if not exists
func CreateIfNotExists(dir string, perm os.FileMode) error {
	if Exists(dir) {
		return nil
	}

	if err := os.MkdirAll(dir, perm); err != nil {
		return fmt.Errorf("failed to create directory: '%s', error: '%s'", dir, err.Error())
	}

	return nil
}

func CopyDirectory(scrDir, dest string) error {
	entries, err := ioutil.ReadDir(scrDir)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		sourcePath := filepath.Join(scrDir, entry.Name())
		destPath := filepath.Join(dest, entry.Name())

		fileInfo, err := os.Stat(sourcePath)
		if err != nil {
			return err
		}

		switch fileInfo.Mode() & os.ModeType {
		case os.ModeDir:
			if err := CreateIfNotExists(destPath, 0755); err != nil {
				return err
			}
			if err := CopyDirectory(sourcePath, destPath); err != nil {
				return err
			}
		default:
			if err := Copy(sourcePath, destPath); err != nil {
				return err
			}
		}

		if err := os.Chmod(destPath, entry.Mode()); err != nil {
			return err
		}

	}
	return nil
}
