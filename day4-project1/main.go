package main

import (
	"fmt"
	"errors"
)
        
func main(){
   title,content, err := getNoteData()

   if err!=nil{
	   fmt.Pritnln(err)
	   return 
   }
	}

func getNoteData() (string,string,error) {

   title, err := getUserInput("Note Title:")
        if err!=nil {
                fmt.Println(err)
                return "", "", err
        }

        content, err := getUserInput("Note Content:")

        if err!=nil{
                fmt.Println(err)
                return "", "", err
        }

	return title,content,err
}

fucn getUserInput(prompt string){
	fmt.Print(prompt)
	var value string
	fmt.Scanln(&value)

	if(value == "")
	{
		return "",errors.New("Invalid Input")
	}

	return value, nil

}
