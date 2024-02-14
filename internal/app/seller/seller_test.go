package seller

import (
	"errors"
	"testing"

	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/pkg/dto"
	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/repository/mocks"
	"github.com/stretchr/testify/mock"
)

func TestRegisterSeller(t *testing.T) {
	userRepo := mocks.NewUserRepository(t)
	roleRepo := mocks.NewRoleRepository(t)
	sellerRepo := mocks.NewSellerRepository(t)

	service := NewService(sellerRepo, userRepo, roleRepo)

	tests := []struct {
		name            string
		input           dto.RegisterSellerRequest
		setup           func(sellerRepo *mocks.SellerRepository, userRepo *mocks.UserRepository, roleRepo *mocks.RoleRepository)
		isErrorExpected bool
	}{
		{
			name: "date check failed",
			input: dto.RegisterSellerRequest{
				RegisterUserRequest: dto.RegisterUserRequest{
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
				CompanyName:    "abc",
				CompanyAddress: "xyz",
			},
			setup: func(sellerRepo *mocks.SellerRepository, userRepo *mocks.UserRepository, roleRepo *mocks.RoleRepository) {
			},
			isErrorExpected: true,
		},
		{
			name: "email already present",
			input: dto.RegisterSellerRequest{
				RegisterUserRequest: dto.RegisterUserRequest{
					FirstName:   "vinay",
					LastName:    "chopda",
					Email:       "vinaychopda@gmail.com",
					Password:    "123",
					DateOfBirth: "01-01-2024",
					MobileNo:    "1289823498",
					Address:     "pune",
					City:        "pune",
					PostalCode:  872834,
				},
				CompanyName:    "abc",
				CompanyAddress: "xyz",
			},
			setup: func(sellerRepo *mocks.SellerRepository, userRepo *mocks.UserRepository, roleRepo *mocks.RoleRepository) {
				userRepo.On("IsUserWithEmailPresent", mock.Anything).Return(true).Once()
			},
			isErrorExpected: true,
		},
		{
			name: "user created",
			input: dto.RegisterSellerRequest{
				RegisterUserRequest: dto.RegisterUserRequest{
					FirstName:   "vinay",
					LastName:    "chopda",
					Email:       "vinaychopda@gmail.com",
					Password:    "123",
					DateOfBirth: "01-01-2024",
					MobileNo:    "1289823498",
					Address:     "pune",
					City:        "pune",
					PostalCode:  872834,
				},
				CompanyName:    "abc",
				CompanyAddress: "xyz",
			},
			setup: func(sellerRepo *mocks.SellerRepository, userRepo *mocks.UserRepository, roleRepo *mocks.RoleRepository) {
				userRepo.On("IsUserWithEmailPresent", mock.Anything).Return(false).Once()
				roleRepo.On("GetRoleId", mock.Anything).Return(1, errors.New("error")).Once()
				// sellerRepo.On("CreateCompany", mock.Anything).Return(1, nil).Once()
				// sellerRepo.On("CreateSeller", mock.Anything).Return(nil).Once()
			},
			isErrorExpected: true,
		},
		{
			name: "user created",
			input: dto.RegisterSellerRequest{
				RegisterUserRequest: dto.RegisterUserRequest{
					FirstName:   "vinay",
					LastName:    "chopda",
					Email:       "vinaychopda@gmail.com",
					Password:    "123",
					DateOfBirth: "01-01-2024",
					MobileNo:    "1289823498",
					Address:     "pune",
					City:        "pune",
					PostalCode:  872834,
				},
				CompanyName:    "abc",
				CompanyAddress: "xyz",
			},
			setup: func(sellerRepo *mocks.SellerRepository, userRepo *mocks.UserRepository, roleRepo *mocks.RoleRepository) {
				userRepo.On("IsUserWithEmailPresent", mock.Anything).Return(false).Once()
				roleRepo.On("GetRoleId", mock.Anything).Return(1, nil).Once()
				sellerRepo.On("CreateCompany", []interface{}{"abc", "xyz"}).Return(int64(1), errors.New("error")).Once()
				// sellerRepo.On("CreateSeller", mock.Anything).Return(nil).Once()
			},
			isErrorExpected: true,
		},
		{
			name: "user created",
			input: dto.RegisterSellerRequest{
				RegisterUserRequest: dto.RegisterUserRequest{
					FirstName:   "vinay",
					LastName:    "chopda",
					Email:       "vinaychopda@gmail.com",
					Password:    "123",
					DateOfBirth: "01-01-2024",
					MobileNo:    "1289823498",
					Address:     "pune",
					City:        "pune",
					PostalCode:  872834,
				},
				CompanyName:    "abc",
				CompanyAddress: "xyz",
			},
			setup: func(sellerRepo *mocks.SellerRepository, userRepo *mocks.UserRepository, roleRepo *mocks.RoleRepository) {
				userRepo.On("IsUserWithEmailPresent", mock.Anything).Return(false).Once()
				roleRepo.On("GetRoleId", mock.Anything).Return(1, nil).Once()
				sellerRepo.On("CreateCompany", []interface{}{"abc", "xyz"}).Return(int64(1), nil).Once()
				sellerRepo.On("CreateSeller", mock.Anything).Return(nil).Once()
			},
			isErrorExpected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup(sellerRepo, userRepo, roleRepo)

			// test service
			err := service.RegisterSeller(test.input)

			if (err != nil) != test.isErrorExpected {
				t.Errorf("Test Failed, expected error to be %v, but got err %v", test.isErrorExpected, err != nil)
			}
		})
	}
}

func TestGetSellerList(t *testing.T) {
	userRepo := mocks.NewUserRepository(t)
	roleRepo := mocks.NewRoleRepository(t)
	sellerRepo := mocks.NewSellerRepository(t)

	service := NewService(sellerRepo, userRepo, roleRepo)

	tests := []struct {
		name            string
		setup           func(sellerRepo *mocks.SellerRepository, userRepo *mocks.UserRepository, roleRepo *mocks.RoleRepository)
		isErrorExpected bool
	}{
		{
			name: "role error",
			setup: func(sellerRepo *mocks.SellerRepository, userRepo *mocks.UserRepository, roleRepo *mocks.RoleRepository) {
				roleRepo.On("GetRoleId", mock.Anything).Return(1, errors.New("error")).Once()
			},
			isErrorExpected: true,
		},
		{
			name: "role error",
			setup: func(sellerRepo *mocks.SellerRepository, userRepo *mocks.UserRepository, roleRepo *mocks.RoleRepository) {
				roleRepo.On("GetRoleId", mock.Anything).Return(1, nil).Once()
				sellerRepo.On("GetAllSellers", 1).Return([]dto.SellerResponseObject{}, errors.New("error")).Once()
			},
			isErrorExpected: true,
		},
		{
			name: "role error",
			setup: func(sellerRepo *mocks.SellerRepository, userRepo *mocks.UserRepository, roleRepo *mocks.RoleRepository) {
				roleRepo.On("GetRoleId", mock.Anything).Return(1, nil).Once()
				sellerRepo.On("GetAllSellers", 1).Return([]dto.SellerResponseObject{}, nil).Once()
			},
			isErrorExpected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup(sellerRepo, userRepo, roleRepo)

			// test service
			_, err := service.GetSellerList()

			if (err != nil) != test.isErrorExpected {
				t.Errorf("Test Failed, expected error to be %v, but got err %v", test.isErrorExpected, err != nil)
			}
		})
	}
}

func TestDeleteSellerList(t *testing.T) {
	userRepo := mocks.NewUserRepository(t)
	roleRepo := mocks.NewRoleRepository(t)
	sellerRepo := mocks.NewSellerRepository(t)

	service := NewService(sellerRepo, userRepo, roleRepo)

	tests := []struct {
		name            string
		setup           func(sellerRepo *mocks.SellerRepository, userRepo *mocks.UserRepository, roleRepo *mocks.RoleRepository)
		isErrorExpected bool
	}{
		{
			name: "role error",
			setup: func(sellerRepo *mocks.SellerRepository, userRepo *mocks.UserRepository, roleRepo *mocks.RoleRepository) {
				sellerRepo.On("DeleteSeller", 1).Return(errors.New("error")).Once()
			},
			isErrorExpected: true,
		},
		{
			name: "role error",
			setup: func(sellerRepo *mocks.SellerRepository, userRepo *mocks.UserRepository, roleRepo *mocks.RoleRepository) {
				sellerRepo.On("DeleteSeller", 1).Return(nil).Once()
			},
			isErrorExpected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup(sellerRepo, userRepo, roleRepo)

			// test service
			err := service.DeleteSeller(1)

			if (err != nil) != test.isErrorExpected {
				t.Errorf("Test Failed, expected error to be %v, but got err %v", test.isErrorExpected, err != nil)
			}
		})
	}
}
