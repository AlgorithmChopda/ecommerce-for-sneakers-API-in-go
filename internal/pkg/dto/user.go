package dto

import (
	"fmt"
	"reflect"
	"regexp"
	"time"

	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/pkg/apperrors"
	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/pkg/constants"
)

// Buyer type
type RegisterUserRequest struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	DateOfBirth string `json:"date_of_birth"`
	MobileNo    string `json:"mobile_no"`
	Address     string `json:"address"`
	City        string `json:"city"`
	PostalCode  int    `json:"postal_code"`
}

type LoginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (req *RegisterUserRequest) Validate() error {
	err := ValidateStruct(*req)
	if err != nil {
		return err
	}

	_, err = regexp.Match(constants.EMAIL_REGEX, []byte(req.Email))
	if err != nil {
		return apperrors.EmptyError{Message: "invalid email format"}
	}

	if len(req.MobileNo) != 10 {
		return apperrors.EmptyError{Message: "invalid mobile no"}
	}

	err = ValidatePassword(req.Password)
	if err != nil {
		return err
	}

	return nil
}

func (req *LoginUserRequest) Validate() error {
	err := ValidateStruct(*req)
	if err != nil {
		return err
	}
	_, err = regexp.Match(constants.EMAIL_REGEX, []byte(req.Email))
	if err != nil {
		return apperrors.EmptyError{Message: "invalid email format"}
	}

	return nil
}

// Token
type JwtToken struct {
	Id   int
	Role int
}

// Seller Type
type RegisterSellerRequest struct {
	RegisterUserRequest
	CompanyName    string `json:"company_name"`
	CompanyAddress string `json:"company_address"`
}

type SellerResponseObject struct {
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	Email          string    `json:"email"`
	DateOfBirth    time.Time `json:"date_of_birth"`
	MobileNumber   int       `json:"mobile_no"`
	Address        string    `json:"address"`
	City           string    `json:"city"`
	PostalCode     int       `json:"postal_code"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	CompanyName    string    `json:"company_name"`
	CompanyAddress string    `json:"company_address"`
}

func (req *RegisterSellerRequest) Validate() error {
	err := ValidateStruct(*req)
	if err != nil {
		return err
	}
	return nil
}

// TODO move to different file
func ValidateStruct(req interface{}) error {
	v := reflect.ValueOf(req)

	if v.Kind() != reflect.Struct {
		return apperrors.EmptyError{Message: "Input is not a struct"}
	}

	for i := 0; i < v.NumField(); i++ {
		fieldValue := v.Field(i)
		fieldType := fieldValue.Type()

		if fieldType.Kind() == reflect.String {
			if fieldValue.String() == "" {
				return apperrors.EmptyError{Message: fmt.Sprintf("Field '%s' is not present or has invalid value", v.Type().Field(i).Name)}
			}
		}

		if fieldType.Kind() == reflect.Int {
			// Check if the int field is zero
			if fieldValue.Int() == 0 {
				return apperrors.EmptyError{Message: fmt.Sprintf("Field '%s' is not present or has invalid value", v.Type().Field(i).Name)}
			}
		}

		if fieldType.Kind() == reflect.Struct {
			if err := ValidateStruct(fieldValue); err != nil {
				return err
			}
		}
	}

	return nil
}

type UserResponseObject struct {
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	Email        string    `json:"email"`
	DateOfBirth  time.Time `json:"date_of_birth"`
	MobileNumber string    `json:"mobile_no"`
	Address      string    `json:"address"`
	City         string    `json:"city"`
	Role         string    `json:"role"`
	PostalCode   int       `json:"postal_code"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type UserRegisterResponseObject struct {
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	DateOfBirth  string `json:"date_of_birth"`
	MobileNumber string `json:"mobile_no"`
	Address      string `json:"address"`
	City         string `json:"city"`
	PostalCode   int    `json:"postal_code"`
}

func ValidatePassword(password string) error {
	if len(password) < 8 {
		return apperrors.EmptyError{Message: "password must be at least 8 characters long"}
	}

	if ok, _ := regexp.MatchString(`[A-Z]`, password); !ok {
		return apperrors.EmptyError{Message: "password must contain at least one uppercase letter"}
	}

	if ok, _ := regexp.MatchString(`[a-z]`, password); !ok {
		return apperrors.EmptyError{Message: "password must contain at least one lowercase letter"}
	}
	if ok, _ := regexp.MatchString(`\d`, password); !ok {
		return apperrors.EmptyError{Message: "password must contain at least one digit"}
	}

	return nil
}

type SaleResponse struct {
	Month string `json:"month"`
	Sale  string `json:"sale"`
}
