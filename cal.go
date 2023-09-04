package main

import "fmt"

func main() {
	var number1, number2 float64
	var operator string
	fmt.Println(" _____________________       ")
	fmt.Println("|  _________________         |")
	fmt.Println("| |                          | |")

	fmt.Print("Enter the first  number : ")
	fmt.Scan(&number1)

	fmt.Print("Enter the second number : ")
	fmt.Scan(&number2)

	fmt.Print("Enter the operator (+ - * / ) :")
	fmt.Scanln(&operator)

	switch operator {
	case "+":
		fmt.Printf("%.3f %s %.3f = %.3f", number1, operator, number2, number1+number2)

	case "-":
		fmt.Printf("%.3f %s %.3f = %.3f", number1, operator, number2, number1-number2)

	case "*":
		fmt.Printf("%.3f %s %.2f = %.3f", number1, operator, number2, number1*number2)
	case "/":
		if number2 == 0.0 {
			fmt.Println("Divide by zero situation")
		} else {
			fmt.Printf("%.3f %s %.3f = %.3f", number1, operator, number2, number1/number2)

		}
	default:
		fmt.Println("Invalid operator")
	}
	fmt.Println("| |                        | |")
	fmt.Println("| |________________________| |")
	fmt.Println("|     ___ ___ ___   ___      |")
	fmt.Println("|    | 7 | 8 | 9 | | + |     |")
	fmt.Println("|    |___|___|___| |___|     |")
	fmt.Println("|    | 4 | 5 | 6 | | - |     |")
	fmt.Println("|    |___|___|___| |___|     |")
	fmt.Println("|    | 1 | 2 | 3 | | x |     |")
	fmt.Println("|    |___|___|___| |___|     |")
	fmt.Println("|    | . | 0 | = | | / |     |")
	fmt.Println("|    |___|___|___| |___|     |")
	fmt.Println("|____________________________|")
}
