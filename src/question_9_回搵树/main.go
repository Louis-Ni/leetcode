package main

func main()  {
	
}

func isPalindrome(x int) bool {
	if (x < 0){
		return false
	}
	init := x
	ans := 0

	for x != 0{
		ans = ans * 10 + (x % 10)
		x /= 10
	}

	return init == ans
}