package main

import "fmt"

func main() {
	fmt.Println("___ Конвертер валют ___")
	const UsdToEur = 0.85
	const UsdToRub = 83.07
	const EurToRub = UsdToRub / UsdToEur
	sum, originCurrency, targetCurrency := getUserInput()
	fmt.Print(sum, originCurrency, targetCurrency)
	fmt.Print(EurToRub)
}

func getUserInput() (float64, string, string) {
	var sum float64
	var originCurrency string
	var targetCurrency string
	fmt.Printf("Введите сумму: ")
	fmt.Scan(&sum)
	fmt.Printf("Введите исходную валюту: ")
	fmt.Scan(&originCurrency)
	fmt.Printf("Введите целевую валюту: ")
	fmt.Scan(&targetCurrency)
	return sum, originCurrency, targetCurrency
}