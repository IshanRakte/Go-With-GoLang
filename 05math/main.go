package main

import (
	"fmt"
	"math/big"

	// "math/rand"
	"crypto/rand"
	// "time"
)

func main() {
	//random number
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(5))

	//random from cryto
	myRandomNum, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myRandomNum)
}
