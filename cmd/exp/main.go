package main

import (
	"html/template"
	"os"
)

type User struct{
	Name string
	Bio string
}

func main(){

	var user User

	user.Name = "Neel"
	user.Bio = "alert(\"Haha, you have been h4x0r3d!\")"

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