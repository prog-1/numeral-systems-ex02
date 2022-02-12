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

func findStartpoint(num int, base int) (power int) {
	for math.Pow(float64(base), float64(power)) < float64(num) {
		power++
	}
	return power - 1
}

func fromBase10(num int, toBase int) (result string) {
	if num == 0 {
		return "0"
	}
	fmt.Println(findStartpoint(num, toBase))
	for power := findStartpoint(num, toBase); power >= 0; power-- {
		devideBy := int(math.Pow(float64(toBase), float64(power)))

		result += string(base36[int(num/devideBy)])
		num = num % devideBy
	}
	return
}

func main() {
	for {
		var choice int
		fmt.Print(`
1) Convert number to base 10
2) Convert number from base 10
3) Exit
`)
		fmt.Scan(&choice)
		switch choice {
		case 1:
			var num string
			var base int
			fmt.Print("Enter number and base: ")
			fmt.Scan(&num, &base)
			fmt.Printf("%v(base: %v) in base 10 is %v", num, base, toBase10(num, base))
		case 2:
			var num int
			var base int
			fmt.Print("Enter number and base: ")
			fmt.Scan(&num, &base)
			fmt.Printf("%v(base: 10) in base %v is %v", num, base, fromBase10(num, base))
		case 3:
			return
		default:
			fmt.Println("Enter a valid option")
		}
	}
}
