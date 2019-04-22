package validate

import (
	"reflect"
)

type fieldValidator func (reflect.Value) bool

func required(value reflect.Value) bool {
	switch value.Kind() {
		case reflect.Interface:
			return value.IsNil()
		case reflect.String:
			return value.String() == ""
	}

	return true
}

var Validators = map[string]fieldValidator {
	"required": required,
}