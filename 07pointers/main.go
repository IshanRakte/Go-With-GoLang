package main

import "fmt"

func main() {

	var ptr *int
	fmt.Println(ptr) //<nil>

	myNum := 12

	var ptr2 = &myNum
	fmt.Println(ptr2)  //actual address(0xc0000120c0)
	fmt.Println(*ptr2) // 12

	*ptr2 = *ptr2 + 2
	fmt.Println(myNum) //14

}
