package main

import "fmt"
import "strings"
import "math"
func main() {
	ans := convert(clean("  -42"))
	fmt.Println(ans)
}

func convert(sign int, ans string) int{
	abs := 0
	for _, b := range ans{
		abs = abs * 10 + int(b-'0')
		switch {
		case sign == 1 && abs > math.MaxInt32:
			return math.MaxInt32
		case sign == -1 && abs * sign < math.MinInt32:
			return math.MinInt32
		}
	}
	fmt.Println("this is " + ans)
	return sign * abs
}

func clean(s string)(sign int, ans string)  {
	s = strings.TrimSpace(s)
	fmt.Println("i am s" + s)
	if s == "" {
		return
	}
	switch s[0]{
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		sign ,ans = 1, s
	case '+':
		sign, ans = 1, s[1:]
	case '-':
		sign, ans = -1, s[1:]
	default:
		ans = ""
		return
	}

	for i, b := range ans {
		// 比如 s=123abc，那么就取123，也就是s[:3]
		if b < '0' || '9' < b{
			ans = ans[:i]
			break
		}
	}

	return 
}
