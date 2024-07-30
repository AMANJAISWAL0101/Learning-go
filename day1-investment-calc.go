package main

import(
        "fmt"
	"math"
)

func main(){
  
  const inflationRate = 4.8
  var   investmentAmount float64 = 10000000
  years := 10.0
  var expectedReturnRate float64 = 16

  futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)

  futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

  fmt.Println("future Value is: ", futureValue)

  fmt.Println("future Real Value is :", futureRealValue)
}
