package switching

import "fmt"

func TypeSwitch(any interface{}) {
	switch any.(type) {

	case int:
		fmt.Printf("Type is int : %v\n", any)
	case string:
		fmt.Printf("Type is string :%v\n", any)
	case float64:
		fmt.Printf("Type is float64 :%v\n", any)
	case bool:
		fmt.Printf("Type is bool: %v\n", any)
	default:
		fmt.Printf("Unknown type")
	}
}
