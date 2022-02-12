package main

import (
	"fmt"
	"math"
	"strings"
)

func toBase10(num string, fromBase int) int {
	var base10, c, i int
	base36 := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	if 2 <= fromBase && fromBase <= 36 && len(num) > 0 {
		for _, v := range num {
			if strings.Contains(base36[:fromBase], string(v)) {
				if '0' <= v && v <= '9' {
					c = int(v) - 48
				} else if 'A' <= v && v <= 'Z' {
					c = int(v) - 55
				}
				base10 = base10 + c*int(math.Pow(float64(fromBase), float64(len(num)-1-i)))
				i++
			} else {
				base10 = -1
				break
			}
		}
	} else {
		base10 = -1
	}
	return base10
}

func main() {
	var num string
	fmt.Print("Enter number: ")
	fmt.Scan(&num)
	var fromBase int
	fmt.Print("Enter base: ")
	fmt.Scan(&fromBase)
	if toBase10(num, fromBase) != -1 {
		fmt.Println("The number in base 10:", toBase10(num, fromBase))
	} else {
		fmt.Println("ERROR: Incorrect base")
	}
}
