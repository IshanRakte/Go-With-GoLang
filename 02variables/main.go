package main

import "fmt"

//Capital L in LoginToken acts as a public variable
const LoginToken string = "gibberish"

func main() {
	var username string = "Ishan"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedin bool = false
	fmt.Println(isLoggedin)
	fmt.Printf("Variable is of type: %T \n", isLoggedin)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var bigFloat float64 = 455.372937838378373
	fmt.Println(bigFloat)
	fmt.Printf("Variable is of type: %T \n", bigFloat)

	//default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	var anotherVariable2 string
	fmt.Println(anotherVariable2)
	fmt.Printf("Variable is of type: %T \n", anotherVariable2)

	// Implicit Type
	var website = "google.co.in"
	fmt.Println(website)

	// no var style (walrus operator :=)
	// no var style is allowed only inside a method
	numberofUser := 300000
	fmt.Println(numberofUser)

	fmt.Println(LoginToken)
}
