package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")
	for i := 1; i < 6000; i++ {
		if i%2 != 0 {
			if isOddComposite(i) {
				//fmt.Println("input", i)
				isNotGoldbachNumber(i)
			}
		}
	}
}
func isOddComposite(i int) bool {
	for j := 2; j < i/2; j++ {
		if i%j == 0 {
			return true
		}
	}
	return false
}

func isPrime(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func isNotGoldbachNumber(num int) {
	var isNumber bool
	for i := 1; i < num; i++ {
		n1 := 2 * i * i
		tmp := num - n1
		if tmp > 0 {
		//	fmt.Println("Tmp is", tmp)
			if isPrime(tmp) {
				isNumber = false
				break
			} else {
				isNumber = true
				continue
			}
		}
	}
	if isNumber {
		fmt.Println("num  is", num)
	}
}
