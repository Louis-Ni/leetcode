package main

import "fmt"
import "strings"

func main() {
	str := 	"0P"
	ans := isPalindrome(str)
	fmt.Println(ans)
}
func isPalindrome(s string) bool {
	if len(s) == 1{
		return true
	}
	start := 0
	end := len(s) - 1
	ans := true
	for start < end{
		fmt.Println(strings.ToLower(string(s[start])), strings.ToLower(string(s[end])))
		if  s[start] < 48 || s[start] > 57 && s[start] < 65 || s[start]>90 && s[start]<97 || s[start]>122{
			fmt.Println("start", s[start])
			start ++
			continue
		}

		if s[end] < 48 || s[end] > 57 && s[end]< 65 || s[end]>90 && s[end]<97 || s[end]>122 {
			fmt.Println("end", strings.ToLower(string(s[end])))
			end --
			continue
		}

		if strings.ToLower(string(s[start])) == strings.ToLower(string(s[end])) {
			start ++
			end --
		}else{
			fmt.Println("jinru")
			ans = false
			break
		}
	}

	return ans
}