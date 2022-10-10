package repository

import "golang_unit_test/entity"

type CategoryRepository interface {
	Findbyid(id string) *entity.Category
}
