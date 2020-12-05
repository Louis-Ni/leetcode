package main

func main()  {
	
}

func longestCommonPrefix(strs []string) string {
	length := len(strs)
	// 输入了一个空的情况下
	if (length == 0){
		return ""
	}
	
	strLength := len(strs[0])

	for i:=0; i<strLength; i++{
		char := strs[0][i]
		for j := 1; j<length; j++{
			// 当数组下标j的第i个字符 不等于 char时
			// 当 i 超出数组下标j的长度时 返回'
			// 因为公共的长度 肯定是取最小的长度
			// 如果最小长度都不能满足则不可能有比这个更长的了
			if i == len(strs[j]) || strs[j][i] != char{
				return strs[0][0:i]
			} 
		}
	}
	return strs[0]
}