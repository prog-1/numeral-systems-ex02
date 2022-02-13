package main

import (
	"fmt"
	"math"
	"strings"
)

const base36 = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func isValidNumber(num string, base int) bool {
	for _, v := range num {
		if !strings.Contains(base36[:base], string(v)) {
			return false
		}
	}
	return true
}

func toBase10(num string, fromBase int) (out int) {
	length := len(num) - 1
	for i := 0; i <= length; i++ {
		number := strings.Index(base36, string(num[i]))
		out += number * int(math.Pow(float64(fromBase), float64(length-i)))
	}
	return
}

func fromBase10(num int, toBase int) (out string) {
	var digitsinnumber int
	var start float64
	if num == 0 {
		return "0"
	}
	for ; start <= float64(num); digitsinnumber++ {
		start = math.Pow(float64(toBase), float64(digitsinnumber))
	}
	digitsinnumber--
	start = start / float64(toBase)
	for ; digitsinnumber > 0; digitsinnumber-- {
		out += string(base36[int(num/int(start))])
		num = num % int(start)
		start = start / float64(toBase)
	}
	return out
}

func main() {
	var base int
	var mode int
	for {
		fmt.Println("Choose one of the options")
		fmt.Println(`
		1) From base2-base36 to base10
		2) From base10 to base2-base36
		3) Close
		`)
		fmt.Scanln(&mode)
		switch mode {
		case 1:
			var num string
			fmt.Print("Enter base: ")
			fmt.Scanln(&base)
			fmt.Printf("Enter number in base%v that will be converted in base10: ", base)
			fmt.Scanln(&num)
			if isValidNumber(num, base) {
				fmt.Printf("*%v in base%v is %v in base10*\n", num, base, toBase10(num, base))
			} else {
				fmt.Println("Number is entered incorrect")
			}
		case 2:
			var num int
			fmt.Print("Enter number: ")
			fmt.Scanln(&num)
			fmt.Printf("Enter base in which will be converted number %v: ", num)
			fmt.Scanln(&base)
			fmt.Printf("*%v in base10 is %v in base%v*\n", num, fromBase10(num, base), base)
		case 3:
			return
		default:
			fmt.Println("Invalid choise")
		}

	}
}
