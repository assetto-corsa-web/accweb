package instance

import (
	"os"
	"path"
	"slices"

	"github.com/assetto-corsa-web/accweb/internal/pkg/helper"
)

func (s *Instance) GetResultList() ([]string, error) {
	files, err := os.ReadDir(path.Join(s.Path, "results"))
	if err != nil {
		return nil, err
	}

	var resultFiles []string
	for _, f := range files {
		if f.IsDir() {
			continue
		}

		resultFiles = append(resultFiles, f.Name())
	}

	slices.Sort(resultFiles)

	return resultFiles, nil
}

func (s *Instance) GetResultContent(idx int) ([]byte, error) {
	files, err := s.GetResultList()
	if err != nil {
		return nil, err
	}

	if idx < 0 || idx >= len(files) {
		return nil, os.ErrNotExist
	}

	data, err := os.ReadFile(path.Join(s.Path, "results", files[idx]))
	if err != nil {
		return nil, err
	}

	return helper.DecodeBytes(data)
}
