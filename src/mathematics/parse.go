package mathematics

import (
	"errors"
	"strings"
	"text/template"
)

func parseTemplate(tplStr string, data interface{}) (string, error) {
	tpl := template.Must(template.New("t1").Parse(tplStr))

	var buf strings.Builder
	err := tpl.Execute(&buf, data)
	if err != nil {
		return "", errors.New("internal system error: template execute error")
	}
	return buf.String(), nil
}
