package main

import (
	"fmt"
	"math"
	"slices"
)

func main() {
	fmt.Println("___ Конвертер валют ___")

	sum, originCurrency, targetCurrency := getUserInput()
	result := convertCurrency(sum, originCurrency, targetCurrency)

	fmt.Println(result)
}

func getUserInput() (float64, string, string) {
	var originCurrency string
	var sum float64
	var targetCurrency string
	currencies := []string{"USD", "EUR", "RUB"}

	for {
		fmt.Printf("Введите исходную валюту %v: ", currencies)
		fmt.Scan(&originCurrency)

		if slices.Contains(currencies, originCurrency) {
			currencies = filterSlice(currencies, originCurrency)
			break
		}
	}

	for {
		fmt.Printf("Введите сумму: ")
		fmt.Scan(&sum)

		if !math.IsNaN(sum) && !math.IsInf(sum, 0) && sum > 0 {
			break
		}
	}

	for {
		fmt.Printf("Введите целевую валюту %v: ", currencies)
		fmt.Scan(&targetCurrency)

		if slices.Contains(currencies, targetCurrency) {
			break
		}
	}

	return sum, originCurrency, targetCurrency
}


func convertCurrency(sum float64, originCurrency, targetCurrency string) float64 {
	const UsdToEur = 0.85
	const UsdToRub = 83.07
	const EurToRub = UsdToRub / UsdToEur

	if originCurrency == "USD" && targetCurrency == "RUB" {
		return sum * UsdToRub
	}

	if originCurrency == "EUR" && targetCurrency == "RUB" {
		return sum * EurToRub
	}

	if originCurrency == "USD" && targetCurrency == "EUR" {
		return sum * UsdToEur
	}

	if originCurrency == "RUB" && targetCurrency == "USD" {
		return sum / UsdToEur
	}

	if originCurrency == "RUB" && targetCurrency == "EUR" {
		return sum / EurToRub
	}

	return 0
}

func filterSlice(slice []string, target string) []string {
	result := []string{}

	for i := 0; i < len(slice); i++ {
		if string(slice[i]) != target {
			result = append(result, string(slice[i]))
		}
	}

	return result
}