package main

import "fmt"

func main() {
	q := lengthOfLongestSubstring("vbxpvwkkteaiob")
	fmt.Println("max=", q)
}
func lengthOfLongestSubstring(s string) int {
	jq := 0
	maxLength := 0
	set := make(map[byte]int)
	for i, ch := range []byte(s) {
		_, ok := set[ch]
		if ok {
			for _, ok := set[ch]; ok; {
				_, ok := set[ch]
				if ok {
					delete(set, s[jq])
					jq++
				} else {
					break
				}
			}
			set[ch] = i

			if len(set) > maxLength {
				maxLength = len(set)
			}
		} else {
			//没有找到则可以加进去
			set[ch] = i
			if len(set) > maxLength {
				maxLength = len(set)
			}
		}
	}
	return maxLength
}
