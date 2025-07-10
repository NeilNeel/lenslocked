package views

import (
	"fmt"
	"html/template"
	"net/http"
)

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