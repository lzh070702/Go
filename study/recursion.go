package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1) //递归实现阶乘
}

func recursion() {
	fmt.Println(fact(7))

	var fib func(n int) int

	fib = func(n int) int { // fib 是一个闭包
		if n < 2 {
			return n
		}

		return fib(n-1) + fib(n-2) // 递归实现斐波那契数
	}

	fmt.Println(fib(7))
}
