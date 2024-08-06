package main

import(
	"fmt"
	"time"
)

type User struct{
	firstName string
	lastName string
	dob string
	createdAt time.Time
}

func main(){
  userFirstName := getValue("Enter your first name")
  userLastName := getValue("Enter your last name")
  userDOB := getValue("Enter your DOB")

  var appUser User

  appUser = User{
	  firstName: userFirstName,
	  lastName: userLastName,
	  dob: userDOB,
	  createdAt: time.Now(),
	  
  }
  
  outputValue(appUser)

}

func outputValue(u User) {
	fmt.Println(u)
}

func getValue(promptString string) string{
	fmt.Println(promptString)
	var value string 
	fmt.Scan(&value)
	return value

}
