package dto

import (
	"errors"
	"fmt"
	"reflect"
)

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

func (req *RegisterUserRequest) Validate() error {
	fmt.Println("in validate")

	err := ValidateStruct(*req)
	if err != nil {
		return err
	}

	fmt.Println("request validate successfully")
	return nil
}

func ValidateStruct(req interface{}) error {
	v := reflect.ValueOf(req)

	if v.Kind() != reflect.Struct {
		return errors.New("Input is not a struct")
	}

	for i := 0; i < v.NumField(); i++ {
		fieldValue := v.Field(i)
		fieldType := fieldValue.Type()

		if fieldType.Kind() == reflect.String {
			if fieldValue.String() == "" {
				return errors.New(fmt.Sprintf("Field '%s' is not present or has invalid value", v.Type().Field(i).Name))
			}
		}

		if fieldType.Kind() == reflect.Int {
			// Check if the int field is zero
			if fieldValue.Int() == 0 {
				return errors.New(fmt.Sprintf("Field '%s' is not present or has invalid value", v.Type().Field(i).Name))
			}
		}
	}

	return nil
}
