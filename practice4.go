package main

import "fmt"

type Speaker interface {
	Speak()
}

//define type
type Human struct {
	name string
}
type Dog struct {
	name string
}

func (h Human) Speak() {
	fmt.Println(h.name, "says hallo")
}

func (h Dog) Speak() {
	fmt.Println(h.name, "says woff")
}

func makeItSpeak(s Speaker) {
	s.Speak()
}

func main() {
	h := Human{name: "Alex"}
	d := Dog{name: "Alex"}

	makeItSpeak(h)
	makeItSpeak(d)
}
