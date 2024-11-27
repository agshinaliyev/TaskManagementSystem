package implementing

import "fmt"

type PaymentProcessor interface {
	ProcessPayment(amount float64)
}

type CreditCard struct {
}

type Paypal struct {
}

func (c CreditCard) ProcessPayment(amount float64) {
	fmt.Println("Processing credit card payment of", amount)
}
func (p Paypal) ProcessPayment(amount float64) {
	fmt.Println("Processing PayPal payment of ", amount)

}

func Payment(processor PaymentProcessor, amount float64) {
	processor.ProcessPayment(amount)
}
