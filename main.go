package main

import "fmt"

func main() {
	for {
		choice := mainMenu()
		if choice == 1 {
			toBase10()
		} else if choice == 2 {
			fromBase10()
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

func toBase10(num string, fromBase int) int {}

func fromBase10(num int, toBase int) string {}
