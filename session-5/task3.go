package main

func main() {

	x := 10
	y := 20

	println("Before swapping :", x, y)
	swap(&x, &y)
	println("After swapping:", x, y)

}

func swap(a *int, b *int) {
	x := *a
	*a = *b
	*b = x

}
