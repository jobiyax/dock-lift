package main

import (
	"fmt"
	"log"

	"dock-lift-cli/internal/network"
)

func main() {
	// Connexion au serveur local
	message, err := network.Connect("localhost:9090")
	// Gestion d'erreur
	if err != nil {
		log.Fatal(err)
	}

	// Affichage du message reçu
	fmt.Println(message)
}
