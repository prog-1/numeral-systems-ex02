package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	for {
		choice := mainMenu()
		if choice == 1 {
			fmt.Println("Enter the number:")
			var num string
			fmt.Scan(&num)
			fmt.Println("Enter the base B:")
			var b int
			fmt.Scan(&b)
			fmt.Println("Converted number is", toBase10(num, b))
			break
		} else if choice == 2 {
			fmt.Println("Enter the number:")
			var num int
			fmt.Scan(&num)
			fmt.Println("Enter the base B:")
			var b int
			fmt.Scan(&b)
			fmt.Println("Converted number is", fromBase10(num, b))
			break
		} else if choice == 3 {
			break
		} else {
			fmt.Println("ERR: wrong choice", choice)
		}
	}
}

func mainMenu() (choice int) {
	fmt.Println(`Choose your action:
1) Program converts a number in base B to a number in base 10.
2) Program converts a number in base 10 to a number in a given base B.
3) Quit`)
	fmt.Scanln(&choice)
	return choice
}

const dictionary = "0123456789abcdefghijklmnopqrstuvwxyz" // from 1 homework

func toBase10(num string, fromBase int) (v int) {
	for i := 0; i <= len(num)-1; i++ {
		number := strings.Index(dictionary, string(num[i]))
		//string.Index --> https://pkg.go.dev/strings#Index
		v += number * int(math.Pow(float64(fromBase), float64(len(num)-1-i)))
	}
	return v
}

func fromBase10(num int, toBase int) string {
	
}
