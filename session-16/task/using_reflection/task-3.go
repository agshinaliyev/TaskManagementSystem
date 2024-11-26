package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
	City string
}

func SetFieldValue(value interface{}, fieldName string, newValue interface{}) {
	of := reflect.ValueOf(value)
	if of.Kind() == reflect.Ptr {
		of = of.Elem()
		field := of.FieldByName(fieldName)
		if field.IsValid() {
			field.Set(reflect.ValueOf(newValue))
		}
	}
}
func main() {

	person := Person{
		Name: "ali",
		Age:  42,
		City: "baku",
	}
	SetFieldValue(&person, "City", "ganja")
	fmt.Println(&person)

}
