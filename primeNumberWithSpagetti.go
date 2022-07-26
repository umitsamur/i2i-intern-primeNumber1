package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var number int
	var isPrime bool

	f, err := os.Open("a.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		isPrime = true
		number, err = strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Casting problem. Please enter integer data")
		}
		if number < 2 {
			fmt.Println("Number must be greater than 1")
		} else {
			for i := 2; i <= number/2; i++ {
				if number%i == 0 {
					isPrime = false
				}
			}
			if isPrime == true {
				fmt.Println("Number is prime number")
			} else {
				fmt.Println("Number is not prime number")
			}
		}
	}

	f.Close()

}
