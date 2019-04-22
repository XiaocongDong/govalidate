package validate

import (
	"fmt"
	"strings"
	"reflect"
)

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

		fmt.Printf("Name of field: %s, value: %v, tag: %s\n", fieldName, fieldVal, fieldTag)
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
	validationsStr := strings.Trim(strings.TrimLeft(tag, "validate:"), "\"")
	validations := strings.Split(validationsStr, ",")

	for _, validation := range validations {
		ok := Validators[validation](val)

		if !ok {
			return &FieldError{
				tag,
				validation,
				fmt.Sprint(val),
			}
		}
	}

	return nil
}
