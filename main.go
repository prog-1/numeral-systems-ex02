package main

import (
	"fmt"
	"math"
	"strings"
)

const base36 = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func isValidNumber(num string, base int) bool {
	if base < 1 || base > 36 {
		return false
	}
	if base >= 1 && base <= 36 {
		for _, v := range num {
			if !strings.Contains(base36[:base], string(v)) {
				return false
			}
		}
	}
	return true
}

func toBase10(num string, fromBase int) (convert int) {
	length := len(num) - 1
	for i := 0; i <= length; i++ {
		number := strings.Index(base36, string(num[i]))
		convert += number * int(math.Pow(float64(fromBase), float64(length-i)))
	}
	return convert
}

func fromBase10(num int, toBase int) (convert string) {
	if num == 0 {
		return "0"
	}
	var z int
	var number float64
	for number <= float64(num) {
		z++
		number = math.Pow(float64(toBase), float64(z))
	}
	number /= float64(toBase)
	for z > 0 {
		z--
		convert += string(base36[int(num/int(number))])
		num %= int(number)
		number /= float64(toBase)
	}
	return convert
}

func mainMenu() (choice int) {
	fmt.Println(`Choose your action:
1) Convert a number to base 10
2) Convert a number from base 10
3) Quit`)
	fmt.Scanln(&choice)
	return choice
}

func main() {
	var num string
	var number, base int
	for {
		choice := mainMenu()
		if choice == 1 {
			fmt.Print("Enter base:")
			fmt.Scanln(&base)
			fmt.Println("Enter number:")
			fmt.Scanln(&num)
			if isValidNumber(num, base) {
				fmt.Printf("%v in base%v is %v in base10\n", num, base, toBase10(num, base))
			} else {
				fmt.Println("Incorrect base or number")
			}
		} else if choice == 2 {
			fmt.Print("Enter number:")
			fmt.Scanln(&number)
			fmt.Println("Enter base:")
			fmt.Scanln(&base)
			if isValidNumber(num, base) {
				fmt.Printf("%v in base10 is %v in base%v\n", number, fromBase10(number, base), base)
			} else {
				fmt.Println("Incorrect base or number")
			}
		} else if choice == 3 {
			break
		} else {
			fmt.Println("ERR: wrong choice", choice)
		}
	}
}
