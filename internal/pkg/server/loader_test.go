package server

import (
	"io/ioutil"
	"path"
	"testing"
)

func TestLoadServerFromPath(t *testing.T) {
	configDir := path.Join("test_data", "configs")

	dir, err := ioutil.ReadDir(configDir)
	if err != nil {
		t.Fatal(err)
	}

	for _, entry := range dir {
		if !entry.IsDir() {
			continue
		}

		baseDir := path.Join(configDir, entry.Name())
		_, err := LoadServerFromPath(baseDir)
		if err != nil {
			t.Fatal(err)
		}
	}
}
