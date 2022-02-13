package main

import (
	"fmt"
	"os"
	"strconv"
)

const digits = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func toBase10(num string, fromBase int) int {
	ans, _ := strconv.ParseInt(num, fromBase, 0)
	return int(ans)
}

func fromBase10(num int, toBase int) string { // I tried to understand how this function works, but its too dificult, there is lot of work with bits
	return strconv.FormatInt(int64(num), toBase)
}
func checkBase(base int) bool {
	if base < 2 || base > 36 {
		return false
	}
	return true
}
func checkNum(num string, fromBase int) bool {
	for _, r := range num {
		if byte(r) > digits[fromBase] { // is there a way to check if string is empty
			return false
		}
	}
	return true
}
func UIforEmptiness() {
	for {
		fmt.Println(`Choose an action:
		1) Convert number to base 10
		2) Convert number from base 10
		3) Exit`)
		var action uint
		fmt.Scan(&action)
		if action == 1 {
			var num string
			fmt.Print("Enter number: ")
			fmt.Scan(&num)
			var fromBase int
			fmt.Print("Enter base for number you have already entered: ")
			fmt.Scan(&fromBase)
			check := checkBase(fromBase)
			chhecknum := checkNum(num, fromBase)
			if check && chhecknum {
				ans1 := toBase10(num, fromBase)
				fmt.Println(ans1)
			} else {
				return
			}
		} else if action == 2 {
			var num int
			fmt.Print("Enter number: ")
			fmt.Scan(&num)
			var toBase int
			fmt.Print("Enter base in which you want to convert number: ")
			fmt.Scan(&toBase)
			check := checkBase(toBase)
			if check {
				ans := fromBase10(num, toBase)
				fmt.Println(ans)
			} else {
				return
			}
		} else if action == 3 {
			os.Exit(0)
		} else {
			return
		}
	}
}
func main() {
	UIforEmptiness()
}
