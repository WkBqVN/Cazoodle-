package utils

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

func ConvertStringToInt(number string) (int, error) {
	result, err := strconv.Atoi(number)
	if err != nil {
		return -1, err
	}
	return result, nil
}

func ConvertMapToStruct[T any](data map[string]T) (interface{}, error) {
	structType := reflect.StructOf(getStructFields(data))
	instance := reflect.New(structType).Elem()
	for key, value := range data {
		field := instance.FieldByName(capitalizeFirstLetter(key))
		if !field.IsValid() {
			return nil, errors.New("invalid field")
		}
		if field.CanSet() {
			fieldValue := reflect.ValueOf(value)
			field.Set(fieldValue)
		} else {
			return nil, errors.New("can't not set field value")
		}
	}
	return instance.Interface(), nil
}

// structFields creates []reflect.StructField based on map keys
func getStructFields[T any](data map[string]T) []reflect.StructField {
	var fields []reflect.StructField
	for key, value := range data {
		fields = append(fields, reflect.StructField{
			Name: capitalizeFirstLetter(key),
			Type: reflect.TypeOf(value),
			Tag:  reflect.StructTag(fmt.Sprintf(`json:"%s"`, key)),
		})
	}
	return fields
}

// capitalizeFirstLetter capitalizes the first letter of a string
func capitalizeFirstLetter(s string) string {
	if len(s) == 0 {
		return s
	}
	return fmt.Sprintf("%s%s", string(s[0]-32), s[1:])
}
