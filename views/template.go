package views

import (
	"fmt"
	"github/NeilNeel/lenslocked/templates"
	"html/template"
	"net/http"
)

func Must(t Template, err error) Template{
	if err!=nil{
		panic(err)
	}
	return t
}

// ParseFS is a function that parses a template from a file system
func ParseFS(filePath string) (Template, error) {
	tpl, err := template.ParseFS(templates.FS, filePath)
	if err!=nil{
		return Template{}, fmt.Errorf("rendring the error template :%v",err)
	}
	return Template{
		HtmlTpl: tpl,
	}, err
}
// old way of parsing templates, we are using the new way of parsing templates
func Parse(filePath string) (Template, error){
	tpl, error := template.ParseFiles(filePath)
	if error!=nil{
		return Template{}, fmt.Errorf("rendring the error template :%v",error)
	}
	return Template{
		HtmlTpl: tpl,
	}, error
}

type Template struct {
	HtmlTpl *template.Template
}

func (t Template) Execute(w http.ResponseWriter, data interface{}){
	w.Header().Set("Content-Type","text/html; charset=utf-8")
	err := t.HtmlTpl.Execute(w, data)
	if err!=nil{
		panic("Error")
	}
}