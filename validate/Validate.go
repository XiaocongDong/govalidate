package validate

import (
	"fmt"
	"strings"
	"reflect"
)

type Validate struct {
}

type FieldError struct {
	tag string
	param string
	field string
}

func (f *FieldError) Tag () string {
	return f.tag
}

func (f *FieldError) Param () string {
	return f.param
}

func (f *FieldError) Field () string {
	return f.field
}

func (v *Validate) Struct (s interface{}) []*FieldError {
	t := reflect.TypeOf(s)

	val := reflect.Indirect(reflect.ValueOf(s))
	numOfField := val.NumField()
	errors := []*FieldError{}

	for i := 0; i < numOfField; i++ {
		field := t.Field(i)
		fieldName := field.Name
		fieldTag := field.Tag
		fieldVal := val.FieldByName(fieldName)

		// fmt.Printf("Name of field: %s, value: %v, tag: %s", fieldName, fieldVal, fieldTag)
		fieldError := validateField(string(fieldTag), fieldVal)

		if fieldError != nil {
			errors = append(errors, fieldError)
		}
	}

	if len(errors) == 0 {
		return nil
	} else {
		return errors
	}
}

func validateField (tag string, val reflect.Value) *FieldError {
	validationsStr := strings.TrimLeft(tag, "validate:")
	validations := strings.Split(validationsStr, ",")

	for _, validation := range validations {
		switch validation {
			case "required":
				fmt.Println("It is required")
		}
	}

	return nil
}
