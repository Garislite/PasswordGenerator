package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("🔐 PASSWORD GENERATOR")

	var length int
	var useLower, useUpper, useNumbers, useSymbols bool

	fmt.Print("Password length: ")
	fmt.Scan(&length)

	for {
		fmt.Print("Use lowercase letters? (true/false): ")
		fmt.Scan(&useLower)
		fmt.Print("Use uppercase letters? (true/false): ")
		fmt.Scan(&useUpper)
		fmt.Print("Use numbers? (true/false): ")
		fmt.Scan(&useNumbers)
		fmt.Print("Use symbols? (true/false): ")
		fmt.Scan(&useSymbols)

		if useLower || useUpper || useNumbers || useSymbols {
			break
		}
		fmt.Println("You must select at least one character type")
	}

	lower := "abcdefghijklmnopqrstuvwxyz"
	upper := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers := "0123456789"
	symbols := "!@#$%^&*"

	allChars := ""
	if useLower {
		allChars += lower
	}
	if useUpper {
		allChars += upper
	}
	if useNumbers {
		allChars += numbers
	}

	password := ""
	for i := 0; i < length; i++ {
		if useSymbols && rand.Intn(100) < 20 {
			symbolIndex := rand.Intn(len(symbols))
			password += string(symbols[symbolIndex])
		} else {
			randomIndex := rand.Intn(len(allChars))
			password += string(allChars[randomIndex])
		}
	}

	fmt.Println("Generated password:", password)
}
