package tests

import (
	"richard-project-back/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_mock_test(t *testing.T) {
	controller := gomock.NewController(t)
	helloWorldRepositoryMock := mocks.NewMockHelloWorldRepository(controller)
	helloWorldRepositoryMock.EXPECT().GetHelloWorld().Return("Hello World")

	result := helloWorldRepositoryMock.GetHelloWorld()
	assert.Equal(t, "Hello World", result)
}

func Test_Product_mock_test(t *testing.T) {
	controller := gomock.NewController(t)
	prodRepositoryMock := mocks.NewMockProductRepositoy(controller)
	prodRepositoryMock.EXPECT().Test().Return("Good")

	result := prodRepositoryMock.Test()
	assert.Equal(t, "Good", result)
}
