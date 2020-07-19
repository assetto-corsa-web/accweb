package pages

import (
	"github.com/assetto-corsa-web/accweb/config"
	"github.com/emvi/logbuch"
	"html/template"
	"io"
)

var (
	tpl *template.Template
)

// LoadTemplate loads the template files.
func LoadTemplate() {
	var err error
	tpl, err = template.New("").ParseGlob("template/*")

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
