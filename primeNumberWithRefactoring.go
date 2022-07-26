package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func isPrime(number int) (message string) {
	var isPrime bool
	isPrime = true
	for i := 2; i <= number/2; i++ {
		if number%i == 0 {
			isPrime = false
		}
	}
	if isPrime == true {
		message = "Number is prime number"
	} else {
		message = "Number is not prime number"
	}
	return message
}

func check(err error) {
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
}

func stringToInt(text string) (number int) {
	number, err := strconv.Atoi(text)
	if err != nil {
		fmt.Println("Casting problem. Please enter integer data")
	}
	return number
}

func main() {
	var number int

	f, err := os.Open("a.txt")
	check(err)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		number = stringToInt(scanner.Text())
		if number < 2 {
			fmt.Println("Number must be greater or equal than 2")
		} else {
			fmt.Println(isPrime(number))
		}
	}

	f.Close()

}