package main

import "fmt"

func operations() {
	var i string
	fmt.Print("pick an operation, add, minus, multiply, divide: ")
	fmt.Scan(&i)
	if i == "add" {
		add()
	} else if i == "minus" {
		subtract()
	} else if i == "multiply" {
		multiply()
	} else if i == "divide" {
		divide()
	} else {
		fmt.Println("invalid option please try again")
		operations()
	}
}

func processMore(result *int, operation string) {
	var next int
	var choice string

	for {
		fmt.Print("Do you want to add another number? (yes/no): ")
		fmt.Scan(&choice)

		if choice == "yes" {
			fmt.Print("Type out the next number: ")
			fmt.Scan(&next)

			switch operation {
			case "add":
				*result += next
			case "subtract":
				*result -= next
			case "multiply":
				*result *= next
			case "divide":
				if next == 0 {
					fmt.Println("Error: Division by zero is not allowed. Skipping this number.")
					continue
				}
				*result /= next
			}
		} else if choice == "no" {
			break
		} else {
			fmt.Println("Invalid input. Please type 'yes' or 'no'.")
		}
	}
}

func add() {
	var i, j int
	fmt.Print("type out the first number: ")
	fmt.Scan(&i)
	fmt.Print("type out the second number: ")
	fmt.Scan(&j)
	sum := i + j
	processMore(&sum, "add")
	fmt.Println("the answer is : ", sum)
}

func subtract() {
	var i, j int
	fmt.Print("Type out the first number: ")
	fmt.Scan(&i)
	fmt.Print("Type out the second number: ")
	fmt.Scan(&j)

	result := i - j
	processMore(&result, "subtract")

	fmt.Println("The answer is:", result)
}

func multiply() {
	var i, j int
	fmt.Print("Type out the first number: ")
	fmt.Scan(&i)
	fmt.Print("Type out the second number: ")
	fmt.Scan(&j)

	result := i * j
	processMore(&result, "multiply")

	fmt.Println("The answer is:", result)
}

func divide() {
	var i, j int
	fmt.Print("Type out the first number: ")
	fmt.Scan(&i)
	fmt.Print("Type out the second number: ")
	fmt.Scan(&j)

	if j == 0 {
		fmt.Println("Error: Division by zero is not allowed. Restarting division.")
		divide()
		return
	}

	result := i / j
	processMore(&result, "divide")

	fmt.Println("The answer is:", result)
}

func main() {
	fmt.Println("Hello, welcome to my calculator app")
	operations()
}
