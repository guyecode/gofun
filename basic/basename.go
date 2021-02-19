package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(basename("/root/a.txt"))
	fmt.Println(basename("a/b/c"))
	fmt.Println(basename("c.d.go"))
	fmt.Println(basename("abc"))
}

func basename(s string) string {
	slash := strings.LastIndex(s, "/")	//如果没有找到/则返回-1
	s = s[slash+1:]
	if dot:= strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}
