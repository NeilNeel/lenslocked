package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func executeTemplate(w http.ResponseWriter, filePath string){
	w.Header().Set("Content-Type","text/html; charset=utf-8")
	// load the template
	tpl, err := template.ParseFiles(filePath)
	if err!= nil{
		log.Printf("parsing the template %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	err = tpl.Execute(w, nil)
	if err!=nil{
		panic("Error")
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request){
	executeTemplate(w, "./templates/home.gohtml")
}

func contactHandler(w http.ResponseWriter,r *http.Request){
	executeTemplate(w, "./templates/contact.gohtml")
}

func faqHandler(w http.ResponseWriter, r *http.Request){
	executeTemplate(w, "./templates/faq.gohtml")
}

func userHandler(w http.ResponseWriter, r *http.Request){
	// userID := chi.URLParam(r, "userID")
	// fmt.Fprintf(w, "Displaying the user %v", userID)

	tpl, err := template.ParseFiles("./templates/user.gohtml")
	if err!= nil{
		log.Printf("Loading the page %v",err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	err = tpl.Execute(w, chi.URLParam(r, "userID"))
	if err!=nil{
		panic(err)
	}
}


func main(){
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	
	r.Group(func(r chi.Router){
		r.Use(middleware.Logger)
		r.Get("/user/{userID}",userHandler)
	})

	r.NotFound(func (w http.ResponseWriter, r *http.Request)  {
		http.Error(w, "Page Not Found", http.StatusNotFound)
	})
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}