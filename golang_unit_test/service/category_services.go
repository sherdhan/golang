package service

import (
	"errors"
	"golang_unit_test/entity"
	"golang_unit_test/repository"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}

func (service CategoryService) Get(id string) (*entity.Category, error) {
	category := service.Repository.Findbyid(id)
	if category == nil {
		return category, errors.New("category tidak ada")
	} else {
		return category, nil
	}
}
