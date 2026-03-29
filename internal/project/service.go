package project

import "time"

// Stockage et ID global
var projects []Project
var currentID = 1

func Create(name string) Project {
	project := Project{
		ID:        currentID,
		Name:      name,
		CreatedAt: time.Now(),
	}

	currentID++
	// Ajout à la liste
	projects = append(projects, project)

	return project
}

func GetAll() []Project {
	return projects
}
