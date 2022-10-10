package service

import (
	"golang_unit_test/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryrepository = repository.Categoryrepositorymock{Mock: mock.Mock{}}
var categoryservice = CategoryService{Repository: &categoryrepository}

func TestCategoryService_GetNotFound(t *testing.T) {

	categoryrepository.Mock.On("Findbyid", "1").Return(nil)

	category, err := categoryservice.Get("1")

	assert.Nil(t, category)
	assert.NotNil(t, err)
}

// func TestCategoryService_GetSuccess(t *testing.T) {
	
// }
