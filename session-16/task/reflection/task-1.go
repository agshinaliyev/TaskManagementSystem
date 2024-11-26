package main

import (
	"fmt"
	"reflect"
)

func main() {

	IdentifyTypeAndKind("hello")
}

func IdentifyTypeAndKind(value any) {

	reflectValue := reflect.ValueOf(value)
	reflectType := reflect.TypeOf(value)
	kind := reflectValue.Kind()
	fmt.Println("Value:", reflectValue, "Type:", reflectType, "Kind:", kind)
}
