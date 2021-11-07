package main

import "fmt"

func adder() func(value int) int {
	sum := 0

	return func(value int) int {
		sum += value
		return sum
	}
}

type iAdder func(value int) (int, iAdder)

func adder2(base int) iAdder {
	return func(value int) (int, iAdder) {
		return base + value, adder2(base + value)
	}
}

func main() {
	adder := adder()
	for i := 0; i < 10; i++ {
		println(adder(i))
	}

	fmt.Println()

	a := adder2(0)
	for i := 0; i < 10; i++ {
		var s int
		s, a = a(i)
		fmt.Println(s)
	}
}
