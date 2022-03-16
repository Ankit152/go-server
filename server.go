package main

import (
	"fmt"
	"net/http"
	"text/template"
	"time"
)

// welcome struct for user name and time
type User struct {
	Name string
	Time string
}

// main
func main() {

	// struct declaration and initialisation
	user := User{"Gopher", time.Now().Format(time.Stamp)}

	// creating the template after parsing the html file
	templates := template.Must(template.ParseFiles("template/index.html"))

	// used for handling all the static files that are required inside the server
	http.Handle("/static/", //final url can be anything
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static"))))

	// taking the name as input from forms
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//Takes the name from the URL query e.g ?name=Martin, will set user.Name = Martin.
		if name := r.FormValue("name"); name != "" {
			user.Name = name
		}
		//If errors show an internal server error message
		//I also pass the welcome struct to the welcome-template.html file.
		if err := templates.ExecuteTemplate(w, "index.html", user); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	fmt.Println("Listening")
	fmt.Println(http.ListenAndServe(":8080", nil))
}
