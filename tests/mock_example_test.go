package tests

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"richard-project-back/mocks"
	"testing"
)

func Test_mock_test(t *testing.T) {
	controller := gomock.NewController(t)
	helloWorldRepositoryMock := mocks.NewMockHelloWorldRepository(controller)
	helloWorldRepositoryMock.EXPECT().GetHelloWorld().Return("Hello World")

	result := helloWorldRepositoryMock.GetHelloWorld()
	assert.Equal(t, "Hello World", result)
}
