package main

import "fmt"

func intSeq() func() int { //函数 intSeq 的返回值是一个闭包，类型是 func() int
	i := 0
	return func() int {
		i++ //对于外部变量的捕获，直接使用，无需任何声明
		return i
	}
}

func closures() {

	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq() //重新调用 intSeq ，得到全新的闭包 newInts
	fmt.Println(newInts())
}
