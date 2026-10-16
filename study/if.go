package main

import "fmt"

// Go 的 if:条件不加括号、可带初始化语句、没有三元运算符、switch 自动 break。

func ifDemo() {
	// 1) 条件表达式不写括号(C/C++ 写了也能编译,但 gofmt/编译器要求不写)
	if true {
		fmt.Println("if 条件不加圆括号")
	}

	// 2) if 可以带初始化语句,变量作用域仅限本 if/else
	//    C/C++ 等价写法要把 n 声明在外面,泄漏到 if 之后
	if n := len("hello"); n > 3 {
		fmt.Println("init 中的 n =", n)
	} else {
		fmt.Println("n <= 3")
	}
	// fmt.Println(n) // 编译错误:n 只在 if/else 内可见

	// 3) 多返回值 + ok 模式,取代 C 的出参/errno、C++ 的 optional/at()
	m := map[string]int{"key": 100}
	if v, ok := m["key"]; ok {
		fmt.Println("取到", v)
	} else {
		fmt.Println("键不存在")
	}

	// 4) 类型断言也写成 if 条件
	var i interface{} = 42
	if v, ok := i.(int); ok {
		fmt.Println("动态类型是 int,值 =", v)
	}

	// 5) 没有 ?: 三元运算符,用函数或字面量代替
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	fmt.Println("max(3,7) =", max(3, 7))

	// 6) switch 不会自动 fallthrough,case 可写多个值
	day := 3
	switch day {
	case 1, 2, 3:
		fmt.Println("周一~周三")
	case 4:
		fmt.Println("周四(不会掉进下一个 case)")
	default:
		fmt.Println("其他")
	}

	// 7) 无表达式 switch == 替代 else-if 链 / goto,自带 break
	switch {
	case day > 5:
		fmt.Println("周末")
	case day > 2:
		fmt.Println("工作日")
	}

	// 8) switch 可带初始化语句(同 if)
	switch v := day * 2; {
	case v > 10:
		fmt.Println("v =", v, "> 10")
	default:
		fmt.Println("v =", v, "<= 10")
	}
}
