package main

import "fmt"

func main(){
	resultInt := usingGeneric(1,6)
	resultFloat := usingGeneric(1.7,88.8)
	resultString := usingGeneric("obito", "uchiha")

	fmt.Printf("all the three results are %v\n %v\n %v\n",resultInt,resultFloat,resultString)



}

func usingGeneric[T int | float64 | string](a,b T) T {
	return a + b
}
