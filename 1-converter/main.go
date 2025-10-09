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
	currencies := []string{"USD", "EUR", "RUB"}
	originCurrency := getOriginCurrency(currencies)
	currencies = filterSlice(currencies, originCurrency)
	sum := getSum()
	targetCurrency := getTargetCurrency(currencies)

	return sum, originCurrency, targetCurrency
}


func convertCurrency(sum float64, originCurrency, targetCurrency string) float64 {
	const UsdToEur = 0.85
	const UsdToRub = 83.07
	const EurToRub = UsdToRub / UsdToEur

	if originCurrency == "USD" && targetCurrency == "RUB" {
		return sum * UsdToRub
	}

	if originCurrency == "USD" && targetCurrency == "EUR" {
		return sum * UsdToEur
	}

	if originCurrency == "EUR" && targetCurrency == "RUB" {
		return sum * EurToRub
	}

	if originCurrency == "EUR" && targetCurrency == "USD" {
		return sum / UsdToEur
	}

	if originCurrency == "RUB" && targetCurrency == "USD" {
		return sum / UsdToRub
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

func getOriginCurrency(currencies []string) string {
	var originCurrency string

	for {
		fmt.Printf("Введите исходную валюту %v: ", currencies)
		fmt.Scan(&originCurrency)

		if slices.Contains(currencies, originCurrency) {
			return originCurrency
		}
	}
}

func getSum() float64 {
	var sum float64

	for {
		fmt.Printf("Введите сумму: ")
		fmt.Scan(&sum)

		if !math.IsNaN(sum) && !math.IsInf(sum, 0) && sum > 0 {
			return sum
		}
	}
}

func getTargetCurrency(currencies []string) string {
	var targetCurrency string

	for {
		fmt.Printf("Введите целевую валюту %v: ", currencies)
		fmt.Scan(&targetCurrency)

		if slices.Contains(currencies, targetCurrency) {
			return targetCurrency
		}
	}
}