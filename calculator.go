package main

import (
	"fmt"
	"os"
)

func main() {
	var number1, number2 int
	var choice int

	fmt.Println("Введите число от 1 до 10")
	fmt.Scan(&number1)
	if 1 <= number1 {

	} else {
		fmt.Println("число не подходит")
		os.Exit(0)
	}
	//fmt.Print(number1)
	if number1 <= 10 {

	} else {
		fmt.Println("число не подходит")
		os.Exit(0)
	}

	fmt.Println("Введите число 1 до 10")
	fmt.Scan(&number2)
	if 1 <= number2 {

	} else {
		fmt.Println("число не подходит")
		os.Exit(0)
	}

	if number2 <= 10 {

	} else {
		fmt.Println("число не подходит")
		os.Exit(0)
	}

	var x int
	fmt.Println("number 1 = ", number1, "\nnumber 2 =", number2)
	fmt.Println(" choice 1: сумма двух чисел")
	fmt.Println(" choice 2: разность двух чисел")
	fmt.Println(" choice 3: умножение двух чисел")
	fmt.Println(" choice 4: деление двух чисел")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		x = number1 + number2
		fmt.Printf("сумма двух чисел: %d", x)
	case 2:
		x = number1 - number2
		fmt.Printf("разность двух чисел: %d", x)
	case 3:
		x = number1 * number2
		fmt.Printf("умножение двух чисел: %d", x)
	case 4:
		x = number1 / number2
		fmt.Printf("деление двух чисел: %d", x)
	default:
		fmt.Println("неверная операция")
	}
}
