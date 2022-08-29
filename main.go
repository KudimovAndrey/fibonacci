package main

import "fmt"

func main() {
	result := 0
	for i := 0; i <= 10; i++ {
		result = secondFibonacci(i)
		fmt.Printf("secondfibonacci(%d) is: %d\n", i, result)
	}

	f := firstFibonacci()
	for i := 1; i <= 10; i++ {
		fmt.Printf("firstFibonacci(%d) is: %d\n", i, f())
	}
}

func firstFibonacci() func() int {
	x := 0
	y := 1
	return func() int {
		x, y = y, x+y
		return y
	}
}

func secondFibonacci(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = secondFibonacci(n-1) + secondFibonacci(n-2)
	}
	return
}
