package main

import (
	"fmt"
)

// isPrime checks if a number is prime
func isPrime(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func main() {

	generatedNumber := make([]int, 100)
	for i := 1; i <= 100; i++ {
		generatedNumber[i-1] = i
	}

	// Can just for loop from 100 to 1 but the task said need to generate 100 numbers then print them in reverse order.
	for i := len(generatedNumber) - 1; i >= 0; i-- {
		num := generatedNumber[i]
		str := ""

		// Do not print prime numbers.
		if isPrime(num) {
			continue
		}

		// Replace numbers divisible by 3 with the text "Foo."
		if num%3 == 0 {
			str += "Foo"
		}

		// Replace numbers divisible by 5 with the text "Bar."
		if num%5 == 0 {
			str += "Bar"
		}

		if str == "" {
			str = fmt.Sprintf("%d", num)
		}

		// Last comma not trimmed
		fmt.Print(str, ", ")
	}
}
