package project

import (
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

// Gestion du service et de la concurrence
type Service struct {
	projects  []Project
	currentID int
	mu        sync.Mutex
}

func NewService() *Service {
	return &Service{
		projects:  []Project{},
		currentID: 1,
	}
}

func (s *Service) Create(name string) (Project, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	folderName := sanitizeName(name)

	project := Project{
		ID:        s.currentID,
		Name:      name,
		CreatedAt: time.Now(),
	}

	// Création du dossier de stockage
	path := filepath.Join("storage", folderName)
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return Project{}, err
	}

	s.currentID++
	s.projects = append(s.projects, project)

	return project, nil
}

func (s *Service) GetAll() []Project {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.projects
}

// Formatage du nom de dossier
func sanitizeName(name string) string {
	name = strings.ToLower(name)
	name = strings.ReplaceAll(name, " ", "_")
	return name
}
