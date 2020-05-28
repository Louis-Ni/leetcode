package main

import "fmt"


func main() {
	str := "bb"
	newStr:=longestPalindrome(str)
	fmt.Println(newStr)

}
func longestPalindrome(s string) string {
	if len(s) < 2{
		return s
	}
	 maxLength  := 0
	 start  := 0
	for i := 0; i< len(s); i++{
		isCharacterMatch(i, i+1, s, &maxLength, &start)
		isCharacterMatch(i-1, i+1, s, &maxLength, &start)

	}
	fmt.Println(start, maxLength)
	if start ==0 && maxLength ==0 {
		return string(s[0])
	}
	str := s[start:start+maxLength]
	return str
}
func isCharacterMatch(left int, right int, str string, max *int, st *int) {
	for {
		if left >= 0 && right < len(str) && str[left] == str[right]{
			if right - left + 1 > *max{
				*max = right - left + 1
				*st = left
			}
			left = left - 1
			right = right + 1
		}else{
			break
		}
	}
}