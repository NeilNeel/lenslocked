package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func homeHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","text/html; charset=utf-8")
	tpl, err := template.ParseFiles("./templates/home.gohtml")
	if err != nil {
		log.Printf("parsing template %v",err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = tpl.Execute(w, nil)
	if err!= nil{
		panic("Error")
	}
}

func contactHandler(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w, "<h1>Contact Page</h1><p>To get in touch, please email me at <a href=\"barvaliyaneel010@gmail.com\">barvaliyaneel010@gmail.com</a></p>")
}

func faqHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "<p>Q: Is there a free version?</p>")
	fmt.Fprintf(w, "<p>A: Yes! We offer a free trail for 30 days on any paid plans.</p>")
	fmt.Fprintf(w, "<hr/>")
	fmt.Fprintf(w, "<p>Q: What's the most memorable show you watched?</p>")
	fmt.Fprintf(w, "<p>A: Hmm, I think it would be Mr. Plankton</p>")
	fmt.Fprintf(w, "<hr/>")
	fmt.Fprintf(w, "<p>Q: How can I contact you?</p>")
	fmt.Fprintf(w, "<p>A: you can mail me at <a href=\"mailto: barvaliyaneel@gmail.com\">barvaliyaneel@gmail.com</a></p>")
}

func userHandler(w http.ResponseWriter, r *http.Request){
	userID := chi.URLParam(r, "userID")
	fmt.Fprintf(w, "Displaying the user %v", userID)
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