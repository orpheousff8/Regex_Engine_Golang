package main

import (
	"log"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "", log.Lmsgprefix)
	var array = []int64{6, 8, 11}
	for _, num := range array {
		logger.Printf("Factorial of a number %d is %d\n", num, factorial(num))
	}
}

func factorial(num int64) int64 {
	res := int64(1)
	for i := int64(1); i <= num; i++ {
		res *= i
	}
	return res
}
