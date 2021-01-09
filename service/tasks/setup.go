package service

import (
	"todogolist/repository"
)

//Setup the repository db.
func Setup(dryRun bool, repository repository.Repository) (string, error) {
	resp, err := repository.Setup(dryRun)

	if err != nil {
		return "", err
	}

	return resp, nil

}
