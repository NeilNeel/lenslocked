package main

import (
	"html/template"
	"os"
)

func main(){

	user := struct{
		Name string
	}{
		Name: "neel",
	}

	// var tmpl = "hello.gohtml"
	tmpl, err := template.ParseFiles("hello.gohtml")
	if err!= nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, user)
	if err!= nil {
		panic(err)
	}
}