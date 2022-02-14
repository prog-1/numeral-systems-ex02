package main

import (
	"fmt"
	"math"
)

func CurrentPower(toBase int, power int) int {
	return int(math.Pow(float64(toBase), float64(power)))
}

func fromBase10(num int, toBase int) string {
	var baseB string
	if 2 <= toBase && toBase <= 36 {
		k, power := 1, -1 // because "for" makes 1 extra move
		for ; k <= num; k = k * toBase {
			power++
		}

		l := CurrentPower(toBase, power)
		var times int
		base36 := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		for power >= 0 {
			if l <= num {
				l = l + CurrentPower(toBase, power)
				times++
			} else {
				// if toBase >= 10 {
				baseB = baseB + string(base36[times]) // it works like that too
				// } else {
				//	   baseB = baseB + fmt.Sprint(times)
				// }
				times = 0
				num = num - (l - CurrentPower(toBase, power)) // because we have an extra "int(math.Pow(float64(toBase), float64(power)))" here
				power--
				l = CurrentPower(toBase, power)

			}
		}
	} else {
		baseB = "-1"
	}
	return baseB
}

func main() {
	var num int
	fmt.Print("Enter number: ")
	fmt.Scan(&num)
	var toBase int
	fmt.Print("Enter base: ")
	fmt.Scan(&toBase)
	if fromBase10(num, toBase) != "-1" {
		fmt.Printf("The number in base %d: %v\n", toBase, fromBase10(num, toBase))
	} else {
		fmt.Println("ERROR: Incorrect base")
	}
}
