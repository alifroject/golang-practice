package main

import (
	"fmt"
)

func myFunct() {
	fmt.Println("Hallo world")
}

func funcAdd(num1 int, num2 int) int {
	return num1 * num2
}

func funValue(x int, y string, z int) (result string, target1 string) {

	if x < z {
		target1 = y + "is less than z "

	} else {
		target1 = y + "is greater than z"

	}

	return
}

func main() {
	if 20 > 18 {
		fmt.Println("20 is greater than 18")
	}

	x := 20
	var y int = 21

	if x > y {
		fmt.Println("is right")
	} else {
		fmt.Println("is wrong")

	}

	//else if

	time := 22
	if time < 10 {
		fmt.Println("Good morning.")
	} else if time < 20 {
		fmt.Println("Good day.")
	} else {
		fmt.Println("Good evening.")
	}

	var day int = 6

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	}

	// go loop

	//function
	myFunct()
	result := funcAdd(2, 4)
	fmt.Println("Hasil\n", result)

	//define 3 return values

	_, b := funValue(10, "I think this one ", 50)
	fmt.Println(b)

}
