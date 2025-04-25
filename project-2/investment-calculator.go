package main

import (
	"fmt" // usaremos print y tambien pillaremos el imput del usuario gracias a este paquete
	"math"
)

// declaramos variables globales
const inflationRate = 2.5

func main() {
	// declaramos variables y las definimos
	var investmentAmount float64
	var years float64
	var expectedReturnRate float64

	//fmt.Print("Investment amount: ")
	outputText("Investment amount: ")
	fmt.Scan(&investmentAmount)

	//fmt.Print("Investment years: ")
	outputText("Investment years: ")
	fmt.Scan(&years)

	//fmt.Print("Expected return rate: ")
	outputText("Expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	//llamamos a nuestra funcion de calcular
	futureValue, futureRealValue := calculateFutureValue(investmentAmount, expectedReturnRate, years)

	formattedFV := fmt.Sprintf("Future Value: %.2f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future Value (adjusted for inflation): %.2f\n", futureRealValue)

	fmt.Print(formattedFV, formattedRFV)

	//fmt.Println("Future Value:", futureValue)
	//fmt.Printf(`Future Value: %.2f\nFuture Valu

	//e (adjusted for inflation): %.2f`, futureValue, futureRealValue)
	//fmt.Println("Future Value (adjusted for inflation):", futureRealValue)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValue(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv = fv / math.Pow(1+inflationRate/100, years)
	return
}
