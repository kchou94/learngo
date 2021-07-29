package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	lastOccurred := make(map[rune]int)
	start := 0
	maxlength := 0
	for i, ch := range []rune(s) {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxlength {
			maxlength = i - start + 1
		}
		lastOccurred[ch] = i
	}

	return maxlength
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println(lengthOfLongestSubstring(""))
	fmt.Println(lengthOfLongestSubstring("这是一个测这试的桑出去个的"))
	fmt.Println(lengthOfLongestSubstring("为什噩梦为晒为喵额为什么塞饭传闻塞饭吃为什喵额"))
	fmt.Println(lengthOfLongestSubstring("一二三三二一"))
}
