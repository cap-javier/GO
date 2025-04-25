package main

import (
	"fmt"
)

func main() {
	// declaramos variables
	// var revenue float64
	// var expenses float64
	// var tax_rate float64

	// pedimos al usuario que introduzca datos
	revenue := inputUser("Introduce el ingreso anual: ")
	expenses := inputUser("Introduce el gasto anual: ")
	tax_rate := inputUser("Introduce el tipo de impuesto: ")

	ebt, eat, ratio := operaciones(revenue, expenses, tax_rate)

	//mandamos datos a pantalla
	fmt.Printf("%.1f\n", ebt)
	fmt.Printf("%.1f\n", eat)
	fmt.Printf("%.3f\n", ratio)
}

func inputUser(infoText string) float64 {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	return userInput
}

func operaciones(revenue, expenses, tax_rate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	eat := ebt * (1 - tax_rate/100)
	ratio := ebt / eat
	return ebt, eat, ratio
}