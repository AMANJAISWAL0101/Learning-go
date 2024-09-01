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

	fmt.Println("future Value is: ", futureValue)

	fmt.Println("future Real Value is :", futureRealValue)
}
