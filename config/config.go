package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Domain string `json:"domain"`
	IP     string `json:"ip"`
	Port   string `json:"port"`
}

// Retourne l'adresse à utiliser
func (c Config) GetAddress() string {
	if c.Domain != "" {
		// Port 443 pour le domaine
		return fmt.Sprintf("%s:443", c.Domain)
	}

	// Formatage IP et port
	return fmt.Sprintf("%s:%s", c.IP, c.Port)
}

// Charge le fichier de configuration
func LoadConfig(path string) (Config, error) {
	var cfg Config

	// Lecture du fichier
	file, err := os.ReadFile(path)
	if err != nil {
		return cfg, fmt.Errorf("Erreur lors de la lecture du fichier : %w", err)
	}

	// Décodage du JSON
	if err := json.Unmarshal(file, &cfg); err != nil {
		return cfg, fmt.Errorf("Erreur lors de l'analyse du json : %w", err)
	}

	// Port par défaut
	if cfg.Port == "" {
		cfg.Port = "9090"
	}

	// Validation des données
	if cfg.IP == "" && cfg.Domain == "" {
		return cfg, fmt.Errorf("Configuration invalide : domaine ou ip requis")
	}

	return cfg, nil
}
