package server

import (
	"fmt"
	"github.com/emvi/logbuch"
	"html/template"
	"reflect"
	"strings"
)

type selectOption struct {
	value string
	label string
}

// Form generates the HTML form for given server configuration.
func Form(cfg *Server) template.HTML {
	var sb strings.Builder
	sb.WriteString(iterateStruct(cfg.ConfigurationJSON, "", -1))
	return template.HTML(sb.String())
}

func iterateStruct(s interface{}, group string, index int) string {
	var sb strings.Builder
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)

	if t.Kind() != reflect.Struct {
		return ""
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		name := field.Tag.Get("json") // use json as name for form field
		label := field.Tag.Get("label")
		options := field.Tag.Get("options")

		if label == "" {
			continue
		}

		value := v.Field(i)

		switch field.Type.Kind() {
		case reflect.Int:
			sb.WriteString(input(group, index, "number", label, name, getOptions(options), value))
		case reflect.Bool:
			sb.WriteString(checkbox(group, index, label, name, value))
		case reflect.String:
			sb.WriteString(input(group, index, "text", label, name, getOptions(options), value))
		case reflect.Struct:
			if group != "" {
				group += "."
			}

			sb.WriteString(fmt.Sprintf(`<div class="group"><div class="group-title">%s<i class="fas fa-chevron-down"></i></div><div class="group-content">%s</div></div>`,
				label, iterateStruct(value.Interface(), fmt.Sprintf("%s%s", group, name), -1)))
		case reflect.Slice:
			for j := 0; j < value.Len(); j++ {
				sb.WriteString("<fieldset>")
				sb.WriteString(iterateStruct(value.Index(j).Interface(), group, j))
				sb.WriteString("</fieldset>")
			}
		default:
			logbuch.Error("Unknown field type", logbuch.Fields{"type": field.Type.Kind().String()})
		}
	}

	return sb.String()
}

func input(group string, index int, t, label, name string, options []selectOption, value reflect.Value) string {
	// regular input field
	if len(options) == 0 {
		return fmt.Sprintf(`<label>%s<input type="%s" name="%s%s%s" value="%v" /></label>`, label, t, group, getIndex(index, group), name, value)
	}

	// select field with options
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf(`<label>%s<select type="%s" name="%s%s%s">`, label, t, group, getIndex(index, group), name))
	strValue := fmt.Sprintf("%v", value) // allow string comparison

	for _, opt := range options {
		selected := ""

		if strValue == opt.value {
			selected = "selected"
		}

		sb.WriteString(fmt.Sprintf(`<option value="%s" %s>%s</option>`, opt.value, selected, opt.label))
	}

	sb.WriteString("</select></label>")
	return sb.String()
}

func checkbox(group string, index int, label, name string, value reflect.Value) string {
	checked := ""

	if value.Bool() {
		checked = "checked"
	}

	return fmt.Sprintf(`<label><input type="checkbox" name="%s%s%s" %s />%s</label>`, group, getIndex(index, group), name, checked, label)
}

func getOptions(tag string) []selectOption {
	options := make([]selectOption, 0)
	parts := strings.Split(tag, ",")

	for _, part := range parts {
		valueLabel := strings.Split(part, ":")

		if len(valueLabel) == 2 {
			options = append(options, selectOption{
				value: strings.TrimSpace(valueLabel[0]),
				label: strings.TrimSpace(valueLabel[1]),
			})
		}
	}

	return options
}

func getIndex(i int, group string) string {
	if i > -1 {
		return fmt.Sprintf("[%d].", i)
	}

	if group == "" {
		return ""
	}

	return "."
}
