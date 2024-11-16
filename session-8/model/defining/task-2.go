package defining

type Speaker interface {
	Speak() string
}
type Dog struct {
}

type Person struct {
}

func (d Dog) Speak() string {

	return "Woof!"
}
func (p Person) Speak() string {
	return "Hello!"
}
func MakeSound(s Speaker) string {
	return s.Speak()
}
