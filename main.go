package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	var m int
	var n int
	var output string

	args := os.Args[1:]

	if len(args) != 2 {
		err1 := "This program only takes two argument as inputs; m and n, respectively"
		err2 := "Where m is the number you want to test whether it is the product of the power of n"
		printLn(err1)
		printLn(err2)
		return
	}

	for _, arg := range args {
		if !isNumeric(arg) {
			err1 := "Your input contains INVALID characters"
			err2 := "This program only takes positive numbers as input"
			printLn(err1)
			printLn(err2)
			return
		}
	}

	m = atoi(args[0])
	n = atoi(args[1])

	if isPower(m, n) {
		output = "true"
		printLn(output)
	} else {
		output = "false"
		printLn(output)
	}
}

func isNumeric(str string) bool {
	var status bool
	for _, ch := range str {
		if ch >= '0' && ch <= '9' {
			status = true
		} else {
			status = false
			break
		}
	}
	return status
}

func printLn(str string) {
	for _, ch := range str {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}

func atoi(str string) int {
	var result int

	for _, ch := range str {
		result = result*10 + int(ch-'0')
	}
	return result
}

func isPower(m, n int) bool {
	var status bool

	for m%n == 0 {
		m /= n

		if m == 1 {
			status = true
		}
	}

	return status
}
