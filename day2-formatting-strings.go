package main

import (
	"fmt"
	"math"
)

func main() {

	var investmentAmount, years, expectedReturnRate, inflationRate float64

	fmt.Println("InflationRate is: ")
	fmt.Scan(&inflationRate)

	fmt.Println("Time duration in years: ")
	fmt.Scan(&years)

	fmt.Println("expected return rate is: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Println("investment amount is: ")
	fmt.Scan(&investmentAmount)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)

	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	formattedFV := fmt.Sprintf("Future Value : %.1f\n", futureValue)

	formattedFRV := fmt.Sprintf("Future real value with inflation : %.1f\n", futureRealValue)

	fmt.Print(formattedFV, formattedFRV)

	fmt.Printf(`Future Value is: %.1f 
Future Value with inflation is: %.1f`, futureValue, futureRealValue)

}
