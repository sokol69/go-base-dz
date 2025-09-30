package main

import "fmt"

func main() {
	const UsdToEur = 0.85
	const UsdToRub = 83.07
	const EurToRub = UsdToRub / UsdToEur

	fmt.Print(EurToRub)
}