package main

import "fmt"

func main() {

	var fruitList [4]string

	fruitList[0] = "Kiwi"
	fruitList[1] = "Lemon"
	fruitList[3] = "Peach"

	fmt.Println(fruitList)
	fmt.Println(len(fruitList))

	var vegList = [3]string{"Broccoli", "Celery", "Zucchini"}
	fmt.Println(vegList)
	fmt.Println(len(vegList))

}
