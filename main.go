package main

import (
	"fmt"
	"math"
	"strings"
)

const base36 = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func toBase10(num string, fromBase int) (res int) {
	// This will work if there is no character that take more than 1 byte
	length := len(num) - 1
	for i := 0; i <= length; i++ {
		number := strings.Index(base36, string(num[i]))
		res += number * int(math.Pow(float64(fromBase), float64(length-i)))
	}
	return
}

func main() {
	var num string
	var base int
	fmt.Print("Enter number and base: ")
	fmt.Scan(&num, &base)
	fmt.Printf("%v(base: %v) in base 10 is %v", num, base, toBase10(num, base))
}
