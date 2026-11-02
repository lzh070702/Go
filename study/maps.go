package main

import (
	"fmt"
	"maps"
)

func mapsExample() {

	m := make(map[string]int) //string 是键的类型，int 是值的类型

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	v3 := m["k3"]
	fmt.Println("v3:", v3) //键不存在时为 0

	fmt.Println("len:", len(m)) //键值对的数量

	delete(m, "k2")
	fmt.Println("map:", m) //移除键值对

	clear(m)
	fmt.Println("map:", m) //清空 map

	_, prs := m["k2"]        //通过空白标识符忽略值
	fmt.Println("prs:", prs) //输出 false

	n := map[string]int{"foo": 1, "bar": 2} //复合字面量
	fmt.Println("map:", n)

	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) { //比较两个 map 是否完全相等
		fmt.Println("n == n2")
	}
}
