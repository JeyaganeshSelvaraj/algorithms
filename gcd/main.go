package main

import (
	"flag"
	"fmt"
	"strconv"
)

var a, b int

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	res := gcd(b, a%b)
	fmt.Printf("gcd(%d, %d) = %d\n", a, b, res)
	return res
}
func parseFlags() {
	flag.Parse()
	if flag.NArg() != 2 {
		fmt.Println("Usage: gcd [a] [b]")
		noOfArgs := flag.NArg()
		fmt.Println("Example: gcd 12 6. Given arguments: ", noOfArgs)
		return
	}
	var err error
	a, err = strconv.Atoi(flag.Arg(0))
	if err != nil {
		fmt.Println("Invalid a")
		return
	}
	b, err = strconv.Atoi(flag.Arg(1))
	if err != nil {
		fmt.Println("Invalid b")
		return
	}
}
func main() {
	parseFlags()
	if a < 0 || b < 0 {
		fmt.Println("a and b must be non-negative")
		return
	}
	parseFlags()
	if a < 0 || b < 0 {
		fmt.Println("a and b must be non-negative")
		return
	}
	fmt.Printf("GCD of %d, %d is %d\n", a, b, gcd(a, b))
}
