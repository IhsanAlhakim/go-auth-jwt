package validation

import (
	"errors"
	"reflect"
	"strings"
)

func CheckStructEmptyProperty(s any) error {
	var reflectValue = reflect.ValueOf(s)

	if reflectValue.Kind() != reflect.Struct {
		return errors.New("input must be a struct")
	}

	for i := 0; i < reflectValue.NumField(); i++ {
		if reflectValue.Field(i).String() == "" {
			return errors.New("field cannot be empty")
		}
	}
	return nil
}

func CheckStructWhitespaceProperty(s any) error {
	var reflectValue = reflect.ValueOf(s)

	if reflectValue.Kind() != reflect.Struct {
		return errors.New("input must be a struct")
	}

	for i := 0; i < reflectValue.NumField(); i++ {
		if strings.Contains(reflectValue.Field(i).String(), " ") {
			return errors.New("field cannot contain whitespace")
		}
	}
	return nil
}
