package helpers

import (
	"bytes"
	"html/template"
	"log"
)

func TemplateToHTML(tmplFile, tmplName string, obj interface{}) string {

	var out bytes.Buffer

	t := template.Must(template.ParseFiles(tmplFile))
	if err := t.ExecuteTemplate(&out, tmplName, obj); err != nil {
		log.Fatal(err)
		return ""
	}

	return out.String()
}
