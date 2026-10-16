package main

import "fmt"

// Go 只有一种循环:for。没有 while / do-while / foreach / repeat。

func forDemo() {
	// 1) 三段式 for,与 C 相同,但小括号必须省略
	for i := 0; i < 3; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	// 2) 条件式 for = while;没有 do-while,想"至少执行一次"用 break
	n := 0
	for n < 3 {
		n++
	}
	fmt.Println("n =", n)

	// 3) 无限循环写 for {},不是 while(true)
	count := 0
	for {
		count++
		if count >= 3 {
			break // 必须显式 break,否则死循环
		}
	}

	// 4) for i := range 整数 (Go 1.22+),C/C++ 里没有的写法
	sum := 0
	for i := range 5 { // i = 0..4
		sum += i
	}
	fmt.Println("sum =", sum)

	// 5) range 遍历 slice / map / string
	fruits := []string{"apple", "banana"}
	for i, v := range fruits {
		fmt.Printf("  [%d] %s\n", i, v)
	}
	for i := range fruits { // 只要下标
		_ = i
	}
	for _, v := range fruits { // 只要元素,_ 表示丢弃
		_ = v
	}

	// 6) range 值是副本,改它不影响原切片;要修改用下标
	for i := range fruits {
		fruits[i] += "!"
	}
	fmt.Println(fruits)

	// 7) 逗号 OK:一次拿到 (值, 是否成功)
	elems := map[string]int{"x": 1, "y": 2}
	for k, v := range elems {
		fmt.Printf("  %s=%d\n", k, v)
	}

	// 8) label + break/continue 控制多层循环(C 也支持,Go 写法一致)
outer:
	for a := 0; a < 3; a++ {
		for b := 0; b < 3; b++ {
			if a*b == 2 {
				break outer
			}
		}
	}

	// 9) Go 1.22 起循环变量每轮都是新变量,
	//    与旧版 Go/C++ 里"闭包捕获同一个变量"的行为不同
	fns := make([]func() int, 3)
	for i := range 3 {
		fns[i] = func() int { return i } // 捕获的是本轮的 i
	}
	for _, f := range fns {
		fmt.Print(f(), " ") // 输出 0 1 2
	}
	fmt.Println()
}
