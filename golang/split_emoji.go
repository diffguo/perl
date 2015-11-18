package main

import (
	"fmt"
)

// 返回一个能正确显示的子字符串（emoji表情, utf8格式）, need_len 为返回的汉字数
func split_emoji_str(str_with_emoji string, need_len int) string {
	str := []rune(str_with_emoji)
	return string(str[:need_len])
}

func main() {
	str := "开始😁👍😭"
	out := split_emoji_str(str, 1)
	fmt.Println(out)
	fmt.Println(len(out))
}
