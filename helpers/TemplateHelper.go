package helpers

import (
	"bytes"
	"html/template"
	"log"
)

// TemplateToHTML - convert template into html
func TemplateToHTML(obj interface{}, tmplName string, pFuncMap *template.FuncMap, tmplFiles ...string) string {

	var out bytes.Buffer

	t := template.Must(template.New(tmplName).Funcs(*pFuncMap).ParseFiles(tmplFiles...))
	if err := t.ExecuteTemplate(&out, tmplName, obj); err != nil {
		log.Fatal(err)
		return ""
	}

	return out.String()
}
