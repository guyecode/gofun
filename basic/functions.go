package main

import "fmt"

// 可变参数的函数
func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}


func main(){
	fmt.Println(sum(1,2,3))
	fmt.Println(sum(1))
}
