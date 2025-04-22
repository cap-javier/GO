package main

import (
	"fmt"
	"math"
) 
//llamamos a la libreria fmt para usar la funcion Print
//llamamos a la libreria math para usar la funcion pow

func main() {
	var investmentAmount float64 = 1000
	var expectedReturnRate float64 = 5.5
	var years float64 = 10

	var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	fmt.Println(futureValue)
}
