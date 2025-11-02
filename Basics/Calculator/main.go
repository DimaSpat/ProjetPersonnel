package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: go run main.go [add|subtract|multiply|divide] num1 num2")
		return
	}

	operation := os.Args[1]
	num1, _ := strconv.ParseFloat(os.Args[2], 64)
	num2, _ := strconv.ParseFloat(os.Args[3], 64)

	switch operation {
	case "add":
		fmt.Printf("Result: %f\n", Add(num1, num2))
	case "subtract":
		fmt.Printf("Result: %f\n", Subtract(num1, num2))
	case "multiply":
		fmt.Printf("Result: %f\n", Multiply(num1, num2))
	case "divide":
		result, err := Divide(num1, num2)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("Result: %f\n", result)
		}
	default:
		fmt.Println("Invalid operation. Use [add|subtract|multiply|divide].")
	}
}

func Add(a, b float64) float64 {
	return a + b
}

func Subtract(a, b float64) float64 {
	return a - b
}

func Multiply(a, b float64) float64 {
	return a * b
}

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("divide by zero")
	} else {
		return a / b, nil
	}
}
