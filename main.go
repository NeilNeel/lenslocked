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
	
	//sample data for home named template
	tpl := views.Must(views.ParseFS("home.gohtml"))
	r.Get("/", controller.StaticHandler(tpl))

	tpl = views.Must(views.ParseFS("contact.gohtml"))
	r.Get("/contact", controller.StaticHandler(tpl))

	tpl = views.Must(views.ParseFS("faq.gohtml"))
	r.Get("/faq", controller.StaticHandler(tpl))

	r.NotFound(func (w http.ResponseWriter, r *http.Request)  {
		http.Error(w, "Page Not Found", http.StatusNotFound)
	})
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}