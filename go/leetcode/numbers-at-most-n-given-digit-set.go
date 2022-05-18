package main

import (
	"fmt"
)

func main() {
	fmt.Println(atMostNGivenDigitSet([]string{"9"}, 55))                       //1
	fmt.Println(atMostNGivenDigitSet([]string{"1"}, 834))                      //3
	fmt.Println(atMostNGivenDigitSet([]string{"3", "4", "5", "6"}, 64))        //18
	fmt.Println(atMostNGivenDigitSet([]string{"3", "4", "6", "7", "9"}, 4170)) // 280
	fmt.Println(atMostNGivenDigitSet([]string{"1", "7"}, 231))                 //10
	fmt.Println(atMostNGivenDigitSet([]string{"5", "7", "8"}, 59))             //6
}

func atMostNGivenDigitSet(digits []string, n int) int {
	dn, pre, now, num := 1, 0, 1, 0
	numstr := [10]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	for n > 0 {
		now, pre = 0, now
		for di := range digits {
			if numstr[n%10] > digits[di] {
				now += dn
			} else if numstr[n%10] == digits[di] {
				now += pre
			} else {
				break
			}
		}
		num += dn
		dn *= len(digits)
		n /= 10
	}
	num--
	return num + now
}
