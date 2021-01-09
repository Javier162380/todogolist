package repository

import (
	"os"
	"todogolist/models"

	"gopkg.in/yaml.v2"
)

type repositoryConfiguration struct {
	DbURI     string `yaml:"dburi"`
	DbDialect string `yaml:"dbdialect"`
}

//Repository main interface for separting db logic with cli logic
type Repository interface {
	Setup(dryRun bool) (string, error)
	GetTask(taskFilterDomain models.TaskFilterDomain) ([]models.TaskDomainResponse, error)
	CreateTask(taskDomains []models.TaskDomain, debugMode bool) (string, error)
	ListTask(taskFilterListDomain models.TaskFilterListDomain) ([]models.TaskDomainResponse, error)
	DeleteTask(taskDeleteDomain models.TaskDeleteDomain) (string, error)
}

func loadConfiguration(configPath string) (*repositoryConfiguration, error) {

	config := &repositoryConfiguration{}

	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	d := yaml.NewDecoder(file)

	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}

//NewRepository generates a new repository.
func NewRepository(repotype string, configPath string) Repository {

	switch repotype {
	case "ent":
		repo := newEntRepository(configPath)
		return repo

	case "memory":
		repo := newMemoryRepository()

		return repo

	default:
		repo := newEntRepository(configPath)
		return repo
	}
}
