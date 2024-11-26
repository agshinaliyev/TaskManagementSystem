package main

import (
	"fmt"
	"reflect"
)

func main() {
	type Person struct {
		Name string
		Age  int
	}
	p := Person{
		Name: "Ali",
		Age:  30,
	}

	InspectStruct(p)

}
func InspectStruct(val any) {
	of := reflect.TypeOf(val)
	value := reflect.ValueOf(val)
	for i := 0; i < of.NumField(); i++ {
		field := of.Field(i)
		fmt.Println("Field name :", field.Name, "Type ", field.Type, "Value", value.Field(i))
	}
}
