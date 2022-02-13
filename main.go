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

func toBase10(num string, fromBase int) int {
	var number int
	pow := len(num) - 1
	for indexBase36 := range num {
		for indexBase10, v := range base36 {
			if byte(v) == num[indexBase36] {
				number += int(float64(indexBase10) * math.Pow(float64(fromBase), float64(pow)))
				pow--
			}
		}
	}
	return number
}

func fromBase10(num int, toBase int) string {
	if num == 0 {
		return "0"
	}
	var x float64
	var number string
	digitsInNumber := 1
	for ; x <= float64(num); digitsInNumber++ {
		x = math.Pow(float64(toBase), float64(digitsInNumber))
	}
	digitsInNumber--
	x /= float64(toBase)
	for ; digitsInNumber != 0; digitsInNumber-- {
		number += string(base36[num/int(x)])
		num %= int(x)
		x /= float64(toBase)
	}
	return number
}

func main() {
	var num string
	var fromBase, toBase int
	fmt.Print("Enter number: ")
	fmt.Scan(&num)
	fmt.Print("Enter from which base to convert the number: ")
	fmt.Scan(&fromBase)
	fmt.Print("Enter to which base to convert the number: ")
	fmt.Scan(&toBase)
	// Fun fact: Theoretically, there is no maximum value for the base of a numeral system. In reality, we don't need very large bases. The maximum of most commonly used base is 360(trecentosexagesimal), used for degrees for angle.
	// List of commonly used numeral systems: https://en.wikipedia.org/wiki/List_of_numeral_systems
	if fromBase < 1 || toBase < 1 {
		fmt.Println("ERR: non-existent base")
	} else if fromBase > 36 || toBase > 36 {
		fmt.Println("The program does not support this base.")
	} else if isValidNumber(num, fromBase) {
		number := toBase10(num, fromBase)
		num = fromBase10(number, toBase)
		fmt.Println("Converted number: ", num)
	} else {
		fmt.Printf("%s is not a valid number for base %v", num, fromBase)
	}
}
