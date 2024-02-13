package user

import (
	"testing"

	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/pkg/dto"
	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/repository/mocks"
	"github.com/stretchr/testify/mock"
)

func TestRegisterUser(t *testing.T) {
	userRepo := mocks.NewUserRepository(t)
	roleRepo := mocks.NewRoleRepository(t)
	service := NewService(userRepo, roleRepo)

	tests := []struct {
		name            string
		input           dto.RegisterUserRequest
		setup           func(userMock *mocks.UserRepository, roleMock *mocks.RoleRepository)
		isErrorExpected bool
	}{
		{
			name: "date check failed",
			input: dto.RegisterUserRequest{
				FirstName:   "vinay",
				LastName:    "chopda",
				Email:       "vinaychopda@gmail.com",
				Password:    "123",
				DateOfBirth: "87234-23-123",
				MobileNo:    "1289823498",
				Address:     "pune",
				City:        "pune",
				PostalCode:  872834,
			},
			setup:           func(userMock *mocks.UserRepository, roleRepo *mocks.RoleRepository) {},
			isErrorExpected: true,
		},
		{
			name: "email already present",
			input: dto.RegisterUserRequest{
				FirstName:   "vinay",
				LastName:    "chopda",
				Email:       "vinaychopda@gmail.com",
				Password:    "123",
				DateOfBirth: "01-01-2023",
				MobileNo:    "1289823498",
				Address:     "pune",
				City:        "pune",
				PostalCode:  872834,
			},
			setup: func(userMock *mocks.UserRepository, roleRepo *mocks.RoleRepository) {
				userMock.On("IsUserWithEmailPresent", mock.Anything).Return(true).Once()
			},
			isErrorExpected: true,
		},
		{
			name: "user created",
			input: dto.RegisterUserRequest{
				FirstName:   "vinay",
				LastName:    "chopda",
				Email:       "vinaychopda@gmail.com",
				Password:    "123",
				DateOfBirth: "01-01-2023",
				MobileNo:    "1289823498",
				Address:     "pune",
				City:        "pune",
				PostalCode:  872834,
			},
			setup: func(userMock *mocks.UserRepository, roleRepo *mocks.RoleRepository) {
				userMock.On("IsUserWithEmailPresent", mock.Anything).Return(false).Once()
				roleRepo.On("GetRoleId", mock.Anything).Return(1, nil).Once()
				userMock.On("CreateUser", mock.Anything).Return(nil).Once()
			},
			isErrorExpected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup(userRepo, roleRepo)

			// test service
			err := service.RegisterUser(test.input, "buyer")

			if (err != nil) != test.isErrorExpected {
				t.Errorf("Test Failed, expected error to be %v, but got err %v", test.isErrorExpected, err != nil)
			}
		})
	}
}
