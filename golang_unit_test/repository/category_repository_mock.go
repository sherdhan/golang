package repository

import (
	"golang_unit_test/entity"

	"github.com/stretchr/testify/mock"
)

type Categoryrepositorymock struct {
	Mock mock.Mock
}

func (repository *Categoryrepositorymock) Findbyid(id string) *entity.Category {
	arguments := repository.Mock.Called(id)
	if arguments.Get(0) == nil {
		return nil
	}
	category := arguments.Get(0).(entity.Category)
	return &category
}
