package main

import (
	"fmt"
	"log"

	"dock-lift-cli/config"
	"dock-lift-cli/internal/network"
)

func main() {
	// Chargement de la configuration
	cfg, err := config.LoadConfig("config.json")
	if err != nil {
		log.Fatal(err)
	}

	// Récupération de l'adresse
	address := cfg.GetAddress()

	// Tentative de connexion
	message, err := network.Connect(address)
	if err != nil {
		log.Fatal(err)
	}

	// Affichage des informations de connexion
	fmt.Println("Connecté à :", address)
	fmt.Println(message)
}
