package main

import (
	"flag"
	"fmt"
	"strconv"
)

var a int

func parseFlags() {
	flag.Parse()
	if flag.NArg() != 1 {
		fmt.Println("Usage: palindrome [a]")
		noOfArgs := flag.NArg()
		fmt.Println("Example: reverse 12. Given arguments: ", noOfArgs)
		return
	}
	var err error
	a, err = strconv.Atoi(flag.Arg(0))
	if err != nil {
		fmt.Println("Invalid a")
		return
	}
}
func reverse(num, rem int) int {
	if num == 0 {
		return rem
	}
	rem = rem*10 + num%10
	return reverse(num/10, rem)
}
func main() {
	parseFlags()
	if a < 0 {
		fmt.Println("a must be non-negative")
		return
	}
	rev := reverse(a, 0)
	if rev == a {
		fmt.Printf("%d is a Palindrome\n", a)
	} else {
		fmt.Printf("%d is not a palindrome\n", a)
	}
}
