package main

import (
	"fmt"
	"math"
	"time"
)

func decimalToBinary(num int) (r []rune) {
	var binary []int

	for num != 0 {
		binary = append(binary, num%2)
		num = num / 2
	}
	if len(binary) == 0 {
		return
	} else {
		for i := len(binary) - 1; i >= 0; i-- {
			r = append(r, rune(binary[i]+48))
		}

	}
	return r
}
func fromBase10(n int, k int) string {
	if k == 2 {
		return string(decimalToBinary(n))
	}
	var a []rune
	for n > 0 {
		var v int
		var tmp, tmp2 int
		if n <= k {
			v = 1
			for tmp2 = 1; n > v*(tmp2); tmp2++ {
				fmt.Println("2  ", v*tmp2)

			}
			var h rune
			if tmp2 >= 10 {
				h = rune(tmp2 + 55)
			} else {
				h = rune(tmp2 + 48)
			}
			a = append(a, h)
			return string(a)
		}
		for n > int(math.Pow(float64(k), float64(tmp))) {
			v = int(math.Pow(float64(k), float64(tmp)))
			fmt.Println(v)
			time.Sleep(1 * time.Second)
			tmp++
		}
		for tmp2 = 1; n > v*(tmp2+1); {
			tmp2++
		}
		v = v * tmp2
		var h rune
		if tmp2 >= 10 {
			h = rune(tmp2 + 55)
		} else {
			h = rune(tmp2 + 48)
		}

		a = append(a, h)
		fmt.Println(a)
		n = n - v
		fmt.Println(n)

	}
	return string(a)
}
func main() {
	fmt.Print("enter number:")
	var n int
	var k int
	fmt.Scanln(&n)
	fmt.Print("Enter P: ")
	fmt.Scanln(&k)
	fmt.Println(fromBase10(n, k))
}
