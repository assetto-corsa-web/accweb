package instance

import (
	"os"
	"path"
	"testing"

	"github.com/assetto-corsa-web/accweb/internal/pkg/helper"

	"github.com/stretchr/testify/assert"
)

func TestLoadServerFromPath(t *testing.T) {
	dir, err := os.MkdirTemp("", "config")
	assert.NoError(t, err)

	err = helper.CopyDirectory(path.Join("test_data", "accweb_old_config_dir"), dir)
	assert.NoError(t, err)

	defer os.RemoveAll(dir)

	s, err := LoadServerFromPath(dir)
	assert.NoError(t, err)
	assert.Equal(t, path.Base(dir), s.GetID())
	assert.Equal(t, "test server name", s.AccCfg.Settings.ServerName)

	ok, err := s.CheckServerExeMd5Sum()
	assert.True(t, ok)
	assert.NoError(t, err)
	assert.Equal(t, "2e0a681866ab68a386ab00de5fc3c126", s.Cfg.Md5Sum)
}
