// +build ignore

package main

import "fmt"

func find(n int, ary []int) bool {
	for _, v := range ary {
		if n == v {
			return true
		}
	}
	return false
}

func position(n int, ary []int) int {
	for i, v := range ary {
		if n == v {
			return i
		}
	}
	return -1
}

func count(n int, ary []int) int {
	c := 0
	for _, v := range ary {
		if n == v {
			c++
		}
	}
	return c
}

func main() {
	a := []int{1, 2, 3, 1, 2, 3, 4, 5}
	fmt.Println(find(4, a))
	fmt.Println(find(6, a))
	fmt.Println(position(5, a))
	fmt.Println(position(7, a))
	fmt.Println(count(3, a))
	fmt.Println(count(8, a))
}
