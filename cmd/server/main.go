// server.go
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var motADeviner string
var lettresDevinees []bool
var displayWord string

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/init", initGame).Methods("GET")
	r.HandleFunc("/deviner/{lettre}", devinerLettre).Methods("POST")

	fmt.Print("Entrez le mot à faire deviner : ")
	fmt.Scanln(&motADeviner)

	for range motADeviner {
		displayWord += "_"
	}

	lettresDevinees = make([]bool, len(motADeviner))

	log.Fatal(http.ListenAndServe(":8080", r))
}

func initGame(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Nouvelle partie initiée par le serveur !\nMot à deviner : %s", displayWord)
}

func devinerLettre(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	lettre := params["lettre"][0]

	for i, _ := range motADeviner {
		if motADeviner[i] == lettre {
			lettresDevinees[i] = true
		}
	}

	displayWord = ""
	nbLettresDevinees := 0
	for i, char := range motADeviner {
		if lettresDevinees[i] {
			displayWord += string(char)
			nbLettresDevinees++
		} else {
			displayWord += "_"
		}
	}

	if nbLettresDevinees == len(motADeviner) {
		fmt.Fprintf(w, "Mot %s deviné ! Bien joué !", displayWord)
	} else {
		fmt.Fprintf(w, "Mot à deviner : %s", displayWord)
	}
}
