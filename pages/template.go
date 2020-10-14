package pages

import (
	"github.com/assetto-corsa-web/accweb/config"
	"github.com/emvi/logbuch"
	"html/template"
	"io"
	"os"
	"path/filepath"
	"strings"
)

const (
	templateDir = "template/"
)

var (
	tpl *template.Template
)

// LoadTemplate loads the template files.
func LoadTemplate() {
	tpl = template.New("").Funcs(funcMap)
	err := filepath.Walk(templateDir, func(path string, info os.FileInfo, err error) error {
		if strings.Contains(path, ".html") {
			if _, err := tpl.ParseFiles(path); err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		logbuch.Fatal("Error loading templates", logbuch.Fields{"err": err})
	}
}

func executeTemplate(w io.Writer, template string, data interface{}) error {
	if config.Get().Server.HotReload {
		LoadTemplate()
	}

	if err := tpl.ExecuteTemplate(w, template, data); err != nil {
		logbuch.Error("Error rendering template", logbuch.Fields{"err": err, "template": template})
		return err
	}

	return nil
}
