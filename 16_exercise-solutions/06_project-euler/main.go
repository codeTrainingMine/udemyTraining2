package main

import (
	"fmt"
	"strings"
)

func main()  {
	var accum int
	for i := 1; i<= 1000; i++ {
		accum += convertLetters(i)
	}
	fmt.Println("final count", accum)
}

func convertLetters(i int) int {
	var s string
	var tens, ones int
	num := i

	if num >= 1000 {
		s = "one-thousand"
	} else {
		if num >= 100 {
			hundreds := num / 100
			hundredsLetters := mapLetters(hundreds) + "-hundred"
			s += hundredsLetters
			num -= hundreds * 100
		}

		tens = num / 10 * 10
		ones = num % 10

		if tens == 10 {
			tens = 0
			ones = num
		}

		if s != "" && (tens > 0 || ones > 0) {
			s += "-and"
		}

		if tens > 0 {
			tensLetters := mapLetters(tens)
			if s != "" {
				s += "-"
			}
			s += tensLetters
		}

		if ones > 0 {
			onesLetters := mapLetters(ones)
			if s != "" {
				s += "-"
			}
			s += onesLetters
		}
	}

	clean := strings.Replace(s, "-", "", -1)
	charlen := len(clean)
	fmt.Println(i, s, num, tens, ones, "'" + clean + "'", charlen)
	return charlen
}

func mapLetters(i int) string  {
	var s string
	switch i {
	case 1: s = "one"
	case 2: s = "two"
	case 3: s = "three"
	case 4: s = "four"
	case 5: s = "five"
	case 6: s = "six"
	case 7: s = "seven"
	case 8: s = "eight"
	case 9: s = "nine"
	case 10: s = "ten"
	case 11: s = "eleven"
	case 12: s = "twelve"
	case 13: s = "thirteen"
	case 14: s = "fourteen"
	case 15: s = "fifteen"
	case 16: s = "sixteen"
	case 17: s = "seventeen"
	case 18: s = "eighteen"
	case 19: s = "nineteen"
	case 20: s = "twenty"
	case 30: s = "thirty"
	case 40: s = "forty"
	case 50: s = "fifty"
	case 60: s = "sixty"
	case 70: s = "seventy"
	case 80: s = "eighty"
	case 90: s = "ninety"
	}
	return s
}


/*
Number letter counts
Problem 17
If the numbers 1 to 5 are written out in words: one, two, three, four, five, then there are 3 + 3 + 5 + 4 + 4 = 19 letters used in total.

If all the numbers from 1 to 1000 (one thousand) inclusive were written out in words, how many letters would be used?


NOTE: Do not count spaces or hyphens. For example, 342 (three hundred and forty-two) contains 23 letters and 115 (one hundred and fifteen) contains 20 letters. The use of "and" when writing out numbers is in compliance with British usage.

==========================

final count 21124

Congratulations, the answer you gave to problem 17 is correct.

You are the 112419th person to have solved this problem.
 */