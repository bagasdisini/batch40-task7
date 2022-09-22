package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	route := mux.NewRouter()

	route.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

	route.HandleFunc("/", home).Methods("GET")
	route.HandleFunc("/contact", contact).Methods("GET")
	route.HandleFunc("/detail-project", blogDetail).Methods("GET")
	route.HandleFunc("/add-project", formAddBlog).Methods("GET")
	route.HandleFunc("/add-blog", addBlog).Methods("POST")

	fmt.Println("Server berjalan di port 8080")

	http.ListenAndServe("localhost:8080", route)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	var template, error = template.ParseFiles("views/index.html")

	if error != nil {
		w.Write([]byte(error.Error()))
		return
	}

	template.Execute(w, nil)
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	var template, error = template.ParseFiles("views/contact.html")

	if error != nil {
		w.Write([]byte(error.Error()))
		return
	}

	template.Execute(w, nil)
}

func blogDetail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	var template, error = template.ParseFiles("views/detail-project.html")

	if error != nil {
		w.Write([]byte(error.Error()))
		return
	}

	template.Execute(w, nil)
}

func formAddBlog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	var template, error = template.ParseFiles("views/add-project.html")

	if error != nil {
		w.Write([]byte(error.Error()))
		return
	}

	template.Execute(w, nil)
}

func addBlog(w http.ResponseWriter, r *http.Request) {
	error := r.ParseForm()
	if error != nil {
		log.Fatal(error)
	}

	fmt.Println("Title : " + r.PostForm.Get("projectName"))
	fmt.Println("Start Date : " + r.PostForm.Get("startDate"))
	fmt.Println("End Date : " + r.PostForm.Get("endDate"))
	fmt.Println("Deskripsi : " + r.PostForm.Get("deskripsi"))

	http.Redirect(w, r, "/add-project", http.StatusMovedPermanently)
}
