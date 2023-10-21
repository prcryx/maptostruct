package main

import (
	"fmt"
	"reflect"
)

type Config struct {
	Port   string `config:"Port"`
	DbUser string `config:"DbUser"`
	DbPass string `config:"DbPass"`
}

type InputMap map[string]string

func (inputMap *InputMap) IsNil() bool {
	return len(*inputMap) > 0
}

func main() {
	config := new(Config)
	myMap := map[string]string{
		"Port":   "8080",
		"DbUser": "postgres",
		"DbPass": "prcryx123",
	}
	myMap1 := new(map[string]string)
	Convert(myMap, config)
	Convert(*myMap1, config)

	fmt.Println(*config)
}

func Convert(input InputMap, out interface{}) error {
	if ok := input.IsNil(); !ok {
		return fmt.Errorf("nil input")
	}
	outVal := reflect.ValueOf(out).Elem()
	fmt.Println(outVal)
	if outVal.Kind() != reflect.Struct {
		return fmt.Errorf("out has to be struct")
	}
	for i := 0; i < outVal.NumField(); i++ {
		field := outVal.Type().Field(i)
		key := field.Tag.Get("config")
		fmt.Println(key)

		if val, ok := input[key]; ok {
			fieldValue := outVal.Field(i)
			fieldType := fieldValue.Type()

			if fieldType.Kind() == reflect.String {
				fieldValue.SetString(val)
			} else {
				errorMsg := fmt.Sprintf("Field %s is not a string type", field.Name)
				return fmt.Errorf(errorMsg)
			}
		}
	}
	return nil
}
