package main

import (
	"errors"
	"fmt"
)

// Go 的函数:多返回值、命名返回值、可变参数、函数即值、方法、defer。

// 1) 多返回值 —— C 用出参/errno,C++ 可用 std::pair,Go 是语法级支持
func div(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("除数不能为 0")
	}
	return a / b, nil
}

// 2) 命名返回值 + 裸 return(慎用,降低可读性)
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return // 等价于 return x, y
}

// 3) 可变参数
func sum(nums ...int) int {
	total := 0
	for _, n := range nums { // nums 是 []int 切片
		total += n
	}
	return total
}

// 4) 函数是一等值,可以赋值、传参、闭包
func apply(f func(int, int) int, a, b int) int {
	return f(a, b)
}

func counter() func() int {
	i := 0
	return func() int { // 闭包捕获 i
		i++
		return i
	}
}

// 5) 方法:带接收者的函数;值接收者 vs 指针接收者
type point struct{ x, y int }

func (p point) dist() int { // 值接收者:操作副本
	return p.x*p.x + p.y*p.y
}

func (p *point) move(dx, dy int) { // 指针接收者:就地修改
	p.x += dx
	p.y += dy
}

func funcDemo() {
	// 1) 多返回值,第二个是错误,不是全局 errno
	if q, err := div(10, 2); err == nil {
		fmt.Println("10/2 =", q)
	}
	if _, err := div(1, 0); err != nil {
		fmt.Println("错误处理:", err)
	}

	// 2) 命名返回值
	a, b := split(100)
	fmt.Println("split:", a, b)

	// 3) 可变参数;整体传切片要加 ...
	parts := []int{1, 2, 3}
	fmt.Println("sum =", sum(1, 2, 3), sum(parts...))

	// 4) 函数值与闭包
	mul := func(x, y int) int { return x * y }
	fmt.Println("apply =", apply(mul, 6, 7))
	next := counter()
	fmt.Println("闭包:", next(), next(), next())

	// 5) 方法
	p := point{3, 4}
	fmt.Println("dist =", p.dist())
	p.move(1, 1)
	fmt.Println("move 后:", p, p.dist())

	// 6) 匿名函数立即执行(IIFE 风格)
	func(msg string) {
		fmt.Println("IIFE:", msg)
	}("hello")

	// 7) defer:后进先出,常用于资源清理
	for i := 0; i < 2; i++ {
		defer fmt.Println("defer", i) // 参数在调用时求值,执行在函数返回前
	}
	fmt.Println("funcDemo 结束")
}
