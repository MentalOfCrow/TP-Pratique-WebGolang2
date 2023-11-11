package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var (
	counter int
)

func main() {
	http.HandleFunc("/", changeHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":8080", nil)
}

func changeHandler(w http.ResponseWriter, r *http.Request) { // Le site Fonction go run . puis localhost:8080 puis refresh la page pour voir la modification
	// meme si des fois il y a 1 faute vous pouvez quand meme le lancer
	fmt.Println("Requête reçue")
	fmt.Println("Compteur actuel : ", counter)
	message := getMessage()

	fmt.Println("Message obtenu : ", message)

	tmpl, err := template.ParseFiles("template/template.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Message string
		Counter int
	}{
		Message: message,
		Counter: counter,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func getMessage() string {

	counter++
	fmt.Println(counter)
	if counter%2 == 0 {
		return "Le nombre de visites est pair!"
	}
	return "Le nombre de visites est impair!"

}
