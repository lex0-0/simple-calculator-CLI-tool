package main

import "fmt"

func main() {
	var number1, number2 float64
	var operator string
	fmt.Println(`                
	            |  _________________         
	            | |                          | |`)

	fmt.Print("Enter the first  number : ")
	fmt.Scan(&number1)

	fmt.Print("Enter the second number : ")
	fmt.Scan(&number2)

	fmt.Print("Enter the operator (+ - * / ) :")
	fmt.Scanln(&operator)

	switch operator {
	case "+":
		fmt.Printf("%.3f %s %.3f \n = %.3f", number1, operator, number2, number1+number2)

	case "-":
		fmt.Printf("%.3f %s %.3f \n = %.3f", number1, operator, number2, number1-number2)

	case "*":
		fmt.Printf("%.3f %s %.2f \n = %.3f", number1, operator, number2, number1*number2)
	case "/":
		if number2 == 0.0 {
			fmt.Println("Divide by zero situation")
		} else {
			fmt.Printf("%.3f %s %.3f \n = %.3f", number1, operator, number2, number1/number2)

		}
	default:
		fmt.Println("Invalid operator")
	}

	fmt.Println(`
						| |________________________| |)
						|     ___ ___ ___   ___      |)
						|    | 7 | 8 | 9 | | + |     |)
						|    |___|___|___| |___|     |)
						|    | 4 | 5 | 6 | | - |     |)
						|    |___|___|___| |___|     |)
						|    | 1 | 2 | 3 | | x |     |)
						|    |___|___|___| |___|     |)
						|    | . | 0 | = | | / |     |)
						|    |___|___|___| |___|     |)
						|____________________________|`)
}
