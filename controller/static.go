package controller

import (
	"github/NeilNeel/lenslocked/views"
	"net/http"
)

func StaticHandler(tpl views.Template) http.HandlerFunc{
	return func(w http.ResponseWriter,r *http.Request){
		tpl.Execute(w,nil)
	}
}