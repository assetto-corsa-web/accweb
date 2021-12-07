package file

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
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

func LoadFromPath(baseDir, filename string, config interface{}) error {
	path := filepath.Join(baseDir, filename)
	f, err := os.Open(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return fmt.Errorf("%e - %w", ErrFileNotFound, err)
		}

		return err
	}

	r := transform.NewReader(f, utf16Encoding.NewDecoder().Transformer)

	if err := json.NewDecoder(r).Decode(config); err != nil {
		if _, errSeek := f.Seek(0, io.SeekStart); errSeek != nil {
			return errSeek
		}

		// trying decode non ut16 content
		if err2 := json.NewDecoder(f).Decode(config); err2 != nil {
			e := fmt.Errorf("decode error: %e. utf16 decode error: %w", err2, err)
			return fmt.Errorf("%e / %w", ErrInvalidJsonFileFormat, e)
		}
	}

	return nil
}
