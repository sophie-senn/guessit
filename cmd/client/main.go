// client.go
package main

import (
	"bufio"
	"fmt"
	"net/http"
)

const srv = "http://localhost:8080"

func main() {
	fmt.Printf("Entrez l'adresse du serveur (par default %s): ", srv)
	var serverAddr string
	fmt.Scanln(&serverAddr)
	if serverAddr == "" {
		serverAddr = srv
	}

	fmt.Println("Nouvelle partie ! Attendez que le serveur initialise le jeu.")

	// Initialisation de la partie
	initGame(serverAddr)

	// Début de la boucle pour deviner les lettres
	for {
		fmt.Print("Entrez une lettre pour deviner le mot : ")
		var lettre string
		fmt.Scanln(&lettre)

		if lettre == "exit" {
			break
		}

		if len(lettre) == 0 {
			continue
		}

		devinerLettre(serverAddr, lettre[0])
	}
}

func initGame(serverAddr string) {
	resp, err := http.Get(serverAddr + "/init")
	if err != nil {
		fmt.Println("Erreur lors de l'initialisation de la partie :", err)
		return
	}
	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Erreur lors de la lecture de la réponse du serveur :", err)
	}
}

func devinerLettre(serverAddr string, lettre byte) {
	resp, err := http.Post(serverAddr+"/deviner/"+string(lettre), "", nil)
	if err != nil {
		fmt.Println("Erreur lors de la devinette de la lettre :", err)
		return
	}
	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Erreur lors de la lecture de la réponse du serveur :", err)
	}
}
