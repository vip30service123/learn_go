package main

import (
	"fmt"
	"hello/practices"
	"math/rand"
	"time"
)

func main() {
	currentTime := time.Now()

	randomNumber := rand.Float32()

	fmt.Println(currentTime)

	fmt.Println(randomNumber)

	random_keywords := practices.RandomKeywords()
	fmt.Println(random_keywords)
}
