package main

import "fmt"

func sum(nums ...int) { //一个函数最多只能有一个变参函数且必须在最后
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func variadicfunctions() {

	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4}
	sum(nums...)//这里的 ... 是把切片拆解成多个参数
}
