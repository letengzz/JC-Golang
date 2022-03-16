package main

import "fmt"

func main() {
	s := []int{1, 3, 5}

	//与数组的方式一样
	//第一种方式:
	for i := 0; i < len(s); i++ {
		fmt.Println(i, s[i])
	}

	//第二种方式:
	for index, value := range s {
		fmt.Println(index, value)
	}
}
