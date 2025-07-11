package main

import (
	"fmt"
	"github/NeilNeel/lenslocked/controller"
	"github/NeilNeel/lenslocked/views"
	"net/http"

	"github.com/go-chi/chi/v5"
)



func main(){
	r := chi.NewRouter()

	tpl, err:=views.Parse("./templates/home.gohtml")
	if err!=nil{
		panic(err)
	}
	r.Get("/", controller.StaticHandler(tpl))

	tpl, err = views.Parse("./templates/contact.gohtml")
	if err!=nil{
		panic(err)
	}
	r.Get("/contact", controller.StaticHandler(tpl))

	tpl, err = views.Parse("./templates/faq.gohtml")
	if err!=nil{
		panic(err)
	}
	r.Get("/faq", controller.StaticHandler(tpl))

	r.NotFound(func (w http.ResponseWriter, r *http.Request)  {
		http.Error(w, "Page Not Found", http.StatusNotFound)
	})
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}