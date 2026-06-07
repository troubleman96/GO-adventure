package main

import ( 
	"fmt"
	"strings"
)

func main(){
	firstName := "user"
	secondName := "iss"
	fullName := firstName + " " + secondName

	fmt.Println(fullName)

	fmt.Println(strings.ToUpper(fullName))
}