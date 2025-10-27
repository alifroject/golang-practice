package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func main() {

	var student1 string = "John" //type is string
	a, b, c := 1, "hello", 2
	var (
		d int
		e string = "Golang"
	)

	const pi float32 = 3.14

	fname := "Jhon"

	result := true

	//struct
	p := Person{Name: "Alex", Age: 21}

	fmt.Println("Hallo world")
	fmt.Println(student1)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(pi)
	fmt.Printf("the number is %d", c)
	fmt.Printf("the number is %d \n", c)
	fmt.Printf("Your first name is %s", fname)
	fmt.Printf("Float number is %.2f \n", pi)
	fmt.Printf("Float number is %t \n", result)
	fmt.Printf("Default is : %v \n", p)
	fmt.Printf("Include field name is : %+v \n", p)
	fmt.Printf("Include representation name is : %#v \n", p)
	fmt.Printf("Include type of value is : %T \n", p)

	//array
	var array1 = [...]int{1, 2, 3}
	fmt.Println("Values are ", array1)
	var array2 = [...]string{"Hallo", "World", "I", "am happy"}
	fmt.Println("Values are => ", array2)

	//slice
	var slice1 = []int{1, 2, 3, 4, 5}
	fmt.Println("Values in slice are ", slice1)
	//create slice

	var slice2 = []int{1, 2, 3, 4, 5}
	arrSlice2 := slice2[2:4]
	fmt.Println("Values slice are ", arrSlice2)

	//slice and append
	var newValue = append(slice2, 6, 7, 8)
	fmt.Println("Values are ", newValue)
	//change value
	newValue[2] = 20
	fmt.Println("New values are ", newValue)


	//cap => is remaining space
	numbers := []int{1, 2, 3, 4, 5, 6}
	sub := numbers[1:4] // [2 3 4]

	fmt.Println("Sub-slice:", sub)
	fmt.Println("Length:", len(sub))
	fmt.Println("Capacity:", cap(sub))


	var (
		sum1 = 100 + 100
		sum2 = sum1 - (sum1 / 2)
		sum3 = sum2 + sum1
	)
	fmt.Println("Value is", sum3)



	//comparison
	fmt.Println(5 > 8)
	fmt.Println(5 <= 8)
	fmt.Println(5 != 8)
	fmt.Println(5 < 8)
}
