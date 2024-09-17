package main

import (
	"fmt"
	"os"
	"strconv"
)

// mathFunc represents a math function.
type mathFunc func(x, y int) int

// Add adds two numbers and returns the result.
//
// For more info, see: https://www.mathsisfun.com/numbers/addition.html
func Add(x, y int) int {
	return x + y
}

// Subtract subtracts two numbers and returns the result.
//
// For more info, see: https://www.mathsisfun.com/numbers/subtraction.html
func Subtract(x, y int) int {
	return x + y
}

func main() {
	if len(os.Args) != 4 {
		fmt.Println("usage: ./quickmaths <op> <x> <y>")
		os.Exit(0)
	}
	op := os.Args[1]
	x, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Printf("%s is not a valid int\n", os.Args[2])
		os.Exit(0)
	}
	y, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Printf("%s is not a valid int\n", os.Args[3])
		os.Exit(0)
	}
	ops := map[string]mathFunc{
		"add": Add,
		"sub": Subtract,
	}
	fmt.Println(ops[op](x, y))
}
