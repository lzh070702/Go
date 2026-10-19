package main

import "fmt"

// Go 的变量与常量:短声明、零值、iota、强制显式类型转换。

func varDemo() {
	// 1) := 短声明(只能在函数内),类型由右侧推断
	x := 42      // int
	name := "go" // string
	ok := true   // bool
	f := 3.14    // float64
	// C/C++ 里 int x = 42; double 会隐式转 int,Go 禁止

	// 2) var 声明(函数内外均可),带零值
	var a, b int         // a=0, b=0
	var s string         // s=""  (不是 NULL)
	var p *int           // p=nil
	var slice []int      // slice=nil
	var m map[string]int // m=nil
	fmt.Println(a, b, s == "", p == nil, slice == nil, m == nil)

	// 3) var 带初始化 + 类型推断
	var c, d = 1, "one" // c int, d string
	var e = f           // e 是 float64

	// 4) _ 空白标识符:丢弃值 / 占位
	var _ int = 99 // 声明后立即丢弃
	// 交换(类型必须一致,Go 不做隐式转换)
	e1, e2 := 1, 2
	e1, e2 = e2, e1
	_ = d
	fmt.Println(x, name, ok, f, c, d, e, e1, e2)

	// 5) 无隐式类型转换,必须显式转换
	var i int = 3
	var fl float64 = float64(i) // 需要写 float64(i)
	var u uint = uint(i)
	// var j int = fl // 编译错误:cannot use fl (variable of type float64)
	// var k int = i + fl // 同样错误
	fv := 3.99
	fmt.Println(fl, u, int(fv)) // 截断为 3,不四舍五入

	// 6) const + iota(编译期常量,不是 C 的 #define 或 const int)
	const (
		Red   = iota // 0
		Green        // 1,同组省略则沿用上一行的表达式
		Blue         // 2
	)
	const (
		KB = 1 << (10 * (iota + 1)) // 1<<10, 1<<20, 1<<30
		MB
		GB
	)
	fmt.Println(Red, Green, Blue, KB, MB, GB)

	// 7) 字符串常量、枚举技巧
	const (
		Hot  = "hot"
		Cold = "cold"
	)
	fmt.Println(Hot, Cold)

	// 8) 变量遮蔽:内层作用域的 := 会新建变量,不修改外层
	y := 10
	if y := 20; y > 15 { // 这个 y 是新的
		fmt.Println("内层 y =", y)
	}
	fmt.Println("外层 y =", y)

	// 9) 循环里 := 同理(Go 1.22 前后语义见 for.go)
	// 10) 多变量赋值 & 多返回值
	var m1, m2 int
	m1, m2 = m2, m1 // 交换无需临时变量
	fmt.Println(m1, m2)
}
