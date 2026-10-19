package main

import "fmt"

// Go 的指针:有 & 和 *,但没有指针运算,没有 ->,没有 malloc/free。

type widget struct{ id int }

func bump(w *widget) { // 指针传参,就地修改(C 的 *w,C++ 的 w->)
	w.id++
}

func zero(p *int) { *p = 0 }

func pointerDemo() {
	// 1) & 取地址,* 解引用;没有 ->,统一用 .
	n := 7
	p := &n
	fmt.Println("值 =", n, "地址 =", p, "解引用 =", *p)
	*p = 8 // 通过指针修改
	fmt.Println("n =", n)

	// 2) 禁止指针运算:不能 p+1、p++(C 里 *(p+i) 遍历数组)
	// arr := [3]int{10, 20, 30}
	// q := &arr[0]
	// q++ // 编译错误

	// 3) 指针作为函数参数(C 的传引用效果)
	w := widget{1}
	bump(&w)
	fmt.Println("widget =", w)

	// 4) new:分配内存并返回指针,零值初始化
	pi := new(int) // *int,值为 0
	fmt.Println("new:", *pi, *pi == 0)
	zero(pi)
	fmt.Println("zero 后:", *pi)

	// 5) new vs make:new 造原始内存(返回指针),make 只用于
	//    slice/map/channel,返回本身就可用的值
	// ns := new([]int)  // *[]int,是 nil 切片的指针,不能直接 append
	s := make([]int, 0, 2) // 可用的切片
	s = append(s, 1)
	m := make(map[string]int, 2) // 预分配容量
	m["a"] = 1
	ch := make(chan int, 1) // 带缓冲的通道
	ch <- 42
	fmt.Println("make slice/map/chan:", s, m, <-ch)

	// 6) nil:未初始化的指针/切片/map/接口/函数值都是 nil
	var up *widget
	var us []int
	var um map[string]int
	var ui interface{}
	var uf func() int
	fmt.Println(up == nil, us == nil, um == nil, ui == nil, uf == nil)
	// *up // panic: 空指针解引用(同 C 的 segfault)

	// 7) 指针的指针
	pp := &p
	fmt.Println(**pp)

	// 8) 切片/map 本身就是引用语义,函数里改会影响外部
	//    (底层数据共享,不需要显式传指针)
	arr := []int{1, 2, 3}
	incAll(arr)
	fmt.Println("切片引用语义:", arr)

	// 9) 结构体一般直接传值(小结构),要修改才传指针
	//    Go 编译器会自动对 w.bump() 形式的方法调用取址
	w.move(5)
	fmt.Println("自动取址:", w)
}

func (w *widget) move(by int) { w.id += by }

func incAll(s []int) {
	for i := range s {
		s[i]++
	}
}
