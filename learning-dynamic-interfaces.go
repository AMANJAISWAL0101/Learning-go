package main

import (
	"fmt"
)

func main(){
	printSomething(6)

}

func printSomething(value interface{}){
	intVal,ok := value.(int)
	
	if ok {
	  fmt.Printf("value is of type integer and its value is %v\n", intVal)
	  return
	}

	floatVal,ok := value.(float64)

	if ok {
		fmt.Print("value is of type float and its value is %v\n", floatVal)
	}

	stringVal,ok := value.(string)

	if ok {
		fmt.Println("value is of type string and its value is %v\n", stringVal)
	}
  
}
