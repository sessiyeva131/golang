// Homework 0: Hello Go!
package main

import "fmt"

func main() {
	// Feel free to use the main function for testing your functions

	fmt.Println("Hello, а¤¦аҐЃа¤Ёа¤їа¤Їа¤ѕ!")

	fmt.Println("Checking Fizzbuzz function:")
	fmt.Println(Fizzbuzz(25))
	fmt.Println(Fizzbuzz(9))
	fmt.Println(Fizzbuzz(15))

	fmt.Println("Checking IsPrime function:")
	fmt.Println(IsPrime(11))
	fmt.Println(IsPrime(9))

	fmt.Println("Checking IsPalindrome function:")
	fmt.Println(IsPalindrome("qazaq"))
	fmt.Println(IsPalindrome("qr"))
}

// Fizzbuzz is a classic introductory programming problem.
// If n is divisible by 3, return "Fizz"
// If n is divisible by 5, return "Buzz"
// If n is divisible by 3 and 5, return "FizzBuzz"
// Otherwise, return the empty string
func Fizzbuzz(n int) string {
	// TODO
	var output string = ""
	if n%3 == 0 && n%5 == 0 {
		output = "FizzBuzz"
	} else if n%5 == 0 {
		output = "Buzz"
	} else if n%3 == 0 {
		output = "Fizz"
	}
	return output
}

// IsPrime checks if the number is prime.
// You may use any prime algorithm, but you may NOT use the standard library.
func IsPrime(n int) bool {
	// TODO
	isPrime := true
	if n == 0 || n == 1 {
		isPrime = false
	} else {
		for i := 2; i <= n/2; i++ {
			if n%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime == true {
			return true
		}
	}
	return false
}

// IsPalindrome checks if the string is a palindrome.
// A palindrome is a string that reads the same backward as forward.
func IsPalindrome(s string) bool {
	// TODO
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}
