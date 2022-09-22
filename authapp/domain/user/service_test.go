package user_test

import (
	"authapp/domain/user"
	"authapp/domain/user/impl"
	"authapp/domain/user/mocks"
	"authapp/dto/request"
	authMocks "authapp/pkg/auth/mocks"
	"errors"

	"testing"

	"github.com/stretchr/testify/assert"
)

type userServiceTest struct {
	MockRepo    *mocks.Repository
	MockService *mocks.Service
	MockAuth    *authMocks.Service
	Service     user.Service
}

var svcTest userServiceTest

func init() {
	mockRepo := new(mocks.Repository)
	mockService := new(mocks.Service)
	mockAuth := new(authMocks.Service)

	svcTest = userServiceTest{
		MockRepo:    mockRepo,
		MockService: mockService,
		MockAuth:    mockAuth,
		Service: impl.NewService(
			mockRepo,
			mockAuth,
		),
	}
}

func TestCreateUserIfNotAny(test *testing.T) {
	req := request.CreateUserRequest{
		Phonenumber: "123456",
		Name:        "User-123",
		Role:        "ADMIN",
	}

	u := &user.User{
		Name:        "User-123",
		Phonenumber: "123456",
		Password:    "ASDF",
		Role:        user.RoleAdmin,
	}

	userWithId := &user.User{
		ID:          1,
		Name:        "User-123",
		Phonenumber: "123456",
		Password:    "QWER",
		Role:        user.RoleAdmin,
	}

	test.Run("error role not valid", func(t *testing.T) {
		req := request.CreateUserRequest{}
		expectedErr := errors.New("role does not exist")

		res, err := svcTest.Service.CreateUserIfNotAny(req)

		assert.Nil(t, res)
		assert.NotNil(t, err)
		assert.Equal(t, expectedErr, err)
	})

	test.Run("error users already exist", func(t *testing.T) {
		expectedErr := errors.New("err Database")

		svcTest.MockAuth.On("GeneratePassword", 4).Return("QWER").Once()
		svcTest.MockAuth.On("EncryptPassword", "QWER").Return("ASDF").Once()
		svcTest.MockRepo.On("GetUserByPhonenumber", "123456").Return(nil, expectedErr).Once()

		resAct, resErr := svcTest.Service.CreateUserIfNotAny(req)

		assert.True(t, svcTest.MockRepo.AssertNotCalled(t, "Persist", u))
		assert.True(t, svcTest.MockAuth.AssertExpectations(t), "mock method from mock auth not called as expected")
		assert.True(t, svcTest.MockRepo.AssertExpectations(t), "mock method from mock repo not called as expected")
		assert.Nil(t, resAct)
		assert.Equal(t, expectedErr, resErr)

	})

	test.Run("error users already exist", func(t *testing.T) {
		svcTest.MockAuth.On("GeneratePassword", 4).Return("QWER").Once()
		svcTest.MockAuth.On("EncryptPassword", "QWER").Return("ASDF").Once()
		svcTest.MockRepo.On("GetUserByPhonenumber", "123456").Return(userWithId, nil).Once()

		expectedErr := errors.New("user with this phonenumber already exist")
		resAct, resErr := svcTest.Service.CreateUserIfNotAny(req)

		assert.True(t, svcTest.MockRepo.AssertNotCalled(t, "Persist", u))
		assert.True(t, svcTest.MockAuth.AssertExpectations(t), "mock method from mock auth not called as expected")
		assert.True(t, svcTest.MockRepo.AssertExpectations(t), "mock method from mock repo not called as expected")
		assert.Nil(t, resAct)
		assert.Equal(t, expectedErr, resErr)

	})

	test.Run("error users already exist", func(t *testing.T) {
		expectedErr := errors.New("err Database")

		svcTest.MockAuth.On("GeneratePassword", 4).Return("QWER").Once()
		svcTest.MockAuth.On("EncryptPassword", "QWER").Return("ASDF").Once()
		svcTest.MockRepo.On("GetUserByPhonenumber", "123456").Return(nil, nil).Once()
		svcTest.MockRepo.On("Persist", u).Return(nil, expectedErr).Once()

		resAct, resErr := svcTest.Service.CreateUserIfNotAny(req)

		assert.True(t, svcTest.MockAuth.AssertExpectations(t), "mock method from mock auth not called as expected")
		assert.True(t, svcTest.MockRepo.AssertExpectations(t), "mock method from mock repo not called as expected")
		assert.Nil(t, resAct)
		assert.Equal(t, expectedErr, resErr)

	})

	test.Run("positive result", func(t *testing.T) {
		svcTest.MockAuth.On("GeneratePassword", 4).Return("QWER").Once()
		svcTest.MockAuth.On("EncryptPassword", "QWER").Return("ASDF").Once()
		svcTest.MockRepo.On("GetUserByPhonenumber", "123456").Return(nil, nil).Once()
		svcTest.MockRepo.On("Persist", u).Return(userWithId, nil).Once()
		res, err := svcTest.Service.CreateUserIfNotAny(req)

		assert.True(t, svcTest.MockAuth.AssertExpectations(t), "mock method from mock auth not called as expected")
		assert.True(t, svcTest.MockRepo.AssertExpectations(t), "mock method from mock repo not called as expected")
		assert.Nil(t, err)
		assert.Equal(t, userWithId, res)

	})

}
