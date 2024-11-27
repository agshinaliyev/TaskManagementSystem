package main

import (
	"fmt"
	"session-8/TaskManagementSystem/session-8/model/defining"
	"session-8/TaskManagementSystem/session-8/model/embedding"
	"session-8/TaskManagementSystem/session-8/model/implementing"
	"session-8/TaskManagementSystem/session-8/model/switching"
)

func main() {
	//task1
	circle := defining.Circle.Area(defining.Circle{Radius: 5})
	rectangle := defining.Rectangle.Area(defining.Rectangle{
		Length: 5,
		Width:  4,
	})
	fmt.Printf("Circle area : %v \n", circle)
	fmt.Printf("Rectangle area : %v \n", rectangle)
	fmt.Println("-----------------------------------")
	// task2
	dog := defining.Dog{}
	person := defining.Person{}
	fmt.Println("Dog says :", defining.MakeSound(dog))
	fmt.Println("Person says :", defining.MakeSound(person))
	fmt.Println("-----------------------------------")

	//task3
	CreditCard := implementing.CreditCard{}
	PayPal := implementing.Paypal{}
	implementing.Payment(CreditCard, 100)
	implementing.Payment(PayPal, 75.5)
	fmt.Println("-----------------------------------")

	//task4
	Email := implementing.EmailNotifier{}
	SMS := implementing.SMSNotifier{}
	implementing.SendMessage(Email, "Your item has been shipped")
	implementing.SendMessage(SMS, "Your item has been shipped")
	fmt.Println("-----------------------------------")
	//task5
	var i interface{}
	i = 42
	fmt.Printf("Value is of type %d \n", i)
	i = "GoLang"
	fmt.Printf("Value is of type %s \n", i)

	i = 3.14
	fmt.Printf("Value is of type %f \n", i)

	//task6
	var anyType interface{}
	anyType = 32
	switching.TypeSwitch(anyType)
	fmt.Println("-----------------------------------")
	//task7
	action := embedding.AutoBot{}
	action.Move()
	action.Say()

}
