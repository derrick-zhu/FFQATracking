package helpers

import (
	"bytes"
	"html/template"
	"log"
)

// TemplateToHTML - convert template into html
func TemplateToHTML(tmplFile, tmplName string, funcMap template.FuncMap, obj interface{}) string {

	var out bytes.Buffer

	t := template.Must(template.New(tmplName).Funcs(funcMap).ParseFiles(tmplFile))
	if err := t.ExecuteTemplate(&out, tmplName, obj); err != nil {
		log.Fatal(err)
		return ""
	}

	return out.String()
}
