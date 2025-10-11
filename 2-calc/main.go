package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("___Calc___")

	operation, values := getUserInput()
	result := calculateResult(operation, values)

	fmt.Println(result)
}

func getUserInput() (string, []int) {
	operation := getOperation()
	values := getValues()

	return operation, values
}

func getOperation() string {
	operations := []string{"AVG", "SUM", "MED"}
	var operation string

	for {
		fmt.Printf("Enter operation %v: ", operations)
		fmt.Scan(&operation)

		if slices.Contains(operations, operation) {
			return operation
		}
	}
}

func getValues() []int {
	var inputStr string

	fmt.Printf("Enter values: ")
	fmt.Scan(&inputStr)

	strSlice := strings.Split(inputStr, ",")
	numberSlice := make([]int, 0, len(strSlice))

	for _, s := range strSlice {
		num, err := strconv.Atoi(s)

		if err != nil {
			fmt.Printf("Error converting '%s' to int: %v\n", s, err)
			return nil
		}

		numberSlice = append(numberSlice, num)
	}

	return numberSlice
}

func calculateResult(operation string, values []int) int {
	switch {
	case operation == "AVG":
		return calculateAvg(values)
	case operation == "SUM":
		return calculateSum(values)
	case operation == "MED":
		return calculateMed(values)
	default:
		return 0
	}
}

func calculateSum(values []int) int {
	res := 0

	for _, value := range values {
		res += value
	}

	return res
}

func calculateAvg(values []int) int {
	sum := calculateSum(values)
	avg := sum / len(values)

	return avg
}

func calculateMed(values []int) int {
	slices.Sort(values)

	mid := len(values) / 2

	if len(values) % 2 == 0 {
		midSum := values[mid - 1] + values[mid]
		return midSum / 2
	} else {
		return values[mid]
	}
}