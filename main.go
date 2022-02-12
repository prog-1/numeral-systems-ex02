package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func Isletter(b byte) bool {
	return b >= 65
}
func tobase10(s string, p int) (a uint64) {
	c := 1

	switch p > 0 {
	case (p <= 10):
		for i := len(s) - 1; i >= 0; i-- {
			a = a + uint64((int(s[i])-48)*c)
			c = c * p
		}

	case p > 10:
		for i := len(s) - 1; i >= 0; i-- {
			if !Isletter(s[i]) {
				a = a + uint64((int(s[i])-48)*c)
				c = c * p
			} else {
				a = a + uint64((int(s[i])-55)*c)
				c = c * p

			}
		}

	}
	return a

}
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
func menu() int {
	fmt.Println(`Please choose you action:
1) frombase10
2) tobase10
3) Stop programm 
Enter number(1-3):`)
	var a uint
	fmt.Scan(&a)
	switch a {
	case 1:
		fmt.Println("enter number to convert:")
		var n int
		var k int
		fmt.Scanln(&n)
		fmt.Print("Enter toBase: ")
		fmt.Scanln(&k)
		fmt.Println(fromBase10(n, k))
	case 2:
		var k int
		fmt.Println("enter number to convert:")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		s := scanner.Text()
		fmt.Println("Enter base of your number: ")
		fmt.Scanln(&k)
		fmt.Println(tobase10(s, k))
	case 3:
		return 4
	default:
		fmt.Println("Incorect value")

	}
	return 0
}
func main() {
	menu()
}
