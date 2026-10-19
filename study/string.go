package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// Go 的字符串与字节:不可变 string、[]byte、rune、range 按字符迭代。

func stringDemo() {
	// 1) string 是字节序列(UTF-8),不可变(C 的 char* 可变,没有保护)
	s := "hello, 世界"
	fmt.Println("len(s) 字节数 =", len(s)) // 7 + 3*2 = 13 字节,不是字符数

	// 2) 索引得到的是 byte(uint8),不是字符
	fmt.Printf("s[0]=%c s[7]=%x\n", s[0], s[7]) // 0xe4 是"世"的首字节
	// s[0] = 'H' // 编译错误:string 不可变
	// s[1]       // 类型是 uint8,不是 byte/rune 的字符比较

	// 3) range 字符串按 rune(码点)迭代,自动解码 UTF-8
	for i, r := range s { // i 是字节偏移,r 是 rune
		fmt.Printf("  %d:%c(%U) ", i, r, r)
	}
	fmt.Println()

	// 4) rune 就是 int32,byte 就是 uint8
	var b byte = 65  // 'A'
	var r rune = '世' // 0x4E16
	fmt.Println(b, r, string(r))

	// 5) 显式转换:string -> []byte(拷贝)、[]byte -> string(拷贝)
	bs := []byte(s) // 可变的字节切片
	bs[0] = 'H'
	s2 := string(bs)
	// []byte("abc") == []byte{'a','b','c'} 只是字面量糖
	fmt.Println(s2, bs)

	// 6) 字符串拼接:+= 或 strings.Builder(C++ 的 ostringstream)
	var sb strings.Builder
	sb.WriteString("a")
	sb.WriteByte('b')
	sb.WriteRune('世')
	fmt.Println("Builder:", sb.String(), "+=:", "x"+"y")

	// 7) len/range 的区别:字节 vs 字符
	cn := "中文"
	fmt.Println("len =", len(cn), "rune 数 =", utf8.RuneCountInString(cn),
		"range 次数 =", strings.Count(cn, "")-1) // Count("")-1 也是一种数字符法

	// 8) 切片是引用,string 拷贝只是头拷贝(底层数组共享只读数据)
	head := s // 不会复制全部字节
	_ = head

	// 9) 子串/查找/替换(无运算符,全在标准库)
	fmt.Println(strings.Contains(s, "世界"),
		strings.Index(s, ","),
		strings.ToUpper("go"),
		strings.Repeat("-", 3),
		strings.TrimSpace(" pad "))

	// 10) []rune(s) 转成可变的码点数组,截断多字节字符时常用
	rs := []rune(cn)
	rs = rs[:1] // 只留"中",不会切断 UTF-8 编码
	fmt.Println(string(rs))
}
