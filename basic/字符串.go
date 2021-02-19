package main

import (
	"fmt"
	"unicode/utf8"
)

func main(){
	s := "你好hello wolrd"
	fmt.Println(len(s))
	fmt.Println(s[0], s[7])
	fmt.Println(s[7:8])
	fmt.Println(s[0:7])

	// 原生字符串，不转义，原样输出
	reg := `\r\n\b\t`
	fmt.Println(reg)

	// 用utf8包处理中文字符

	ch := "Hello, 世界"
	fmt.Println(len(ch))
	fmt.Println(utf8.RuneCountInString(ch))
	for i := 0; i < len(ch); {
		r, size := utf8.DecodeRuneInString(ch[i:])
		fmt.Printf("%d    %c\n", i, r)
		i += size
	}
	// 用range可以简化上面的操作
	for i, r := range ch {
		fmt.Printf("%d    %c\n", i, r)
	}

	fmt.Println(fmt.Sprintf("%d", 123))
}