package pages

import (
	"errors"
	"html/template"
)

var funcMap = template.FuncMap{
	"dict": createDictionary,
}

func createDictionary(values ...interface{}) (map[string]interface{}, error) {
	if len(values)%2 != 0 {
		return nil, errors.New("invalid dict call")
	}

	dict := make(map[string]interface{}, len(values)/2)

	for i := 0; i < len(values); i += 2 {
		key, ok := values[i].(string)

		if !ok {
			return nil, errors.New("dict keys must be strings")
		}

		dict[key] = values[i+1]
	}

	return dict, nil
}
