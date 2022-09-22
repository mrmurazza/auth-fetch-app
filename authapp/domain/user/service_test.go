package user_test

import (
	"authapp/domain/user"
	"authapp/domain/user/impl"
	"authapp/domain/user/mocks"
	"authapp/dto/request"
	"testing"
)

type itemServiceTest struct {
	MockRepo    *mocks.Repository
	MockService *mocks.Service
	Service     user.Service
}

var svcTest itemServiceTest

func init() {
	mockRepo := new(mocks.Repository)
	mockService := new(mocks.Service)

	svcTest = itemServiceTest{
		MockRepo:    mockRepo,
		MockService: mockService,
		Service: impl.NewService(
			mockRepo,
		),
	}
}

func TestCreateUserIfNotAny(t *testing.T) {
	req := request.CreateUserRequest{}

	it := user.User{
		Name: "User-123",
	}

	t.Run("positive result", func(t *testing.T) {
		svcTest.MockRepo.On("GetUserByUsername", it).Return(true, nil).Once()
		svcTest.MockRepo.On("Persist", it).Once()
		svcTest.Service.CreateUserIfNotAny(req)
	})

}
