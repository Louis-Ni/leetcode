package main

func main() {
	nums := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 50}}
	searchMatrix(nums, 13)
}

func searchMatrix(matrix [][]int, target int) bool {
	left := 0
	right := len(matrix)
	if right == 0 || len(matrix[0]) == 0 {
		return false
	}
	//查找元素所在的行
	for left < right {
		mid := (left + right) / 2
		if matrix[mid][0] == target {
			return true
		}
		if matrix[mid][0] <= target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	tmp := 0
	if right > 0 {
		tmp = right - 1 //right>0的情況下，因為right處於目標行數的下一行
	} else {
		tmp = right //right = 0的情況
	}
	left = 0
	right = len(matrix[tmp]) //因為我們已經找到元素所在的行了,因此可以直接獲取這一行的元素，再進行一次二分查找

	for left < right {
		mid := (left + right) / 2
		if matrix[tmp][mid] == target {
			return true
		}
		if matrix[tmp][mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return false

	//此處還有解決方案2：
	//  將整個二位數組轉換成一位數組，再進行一次二分查找
	//此處還有解決方案3：
	//  可使用雙指針法，指向頭部和尾部, 進行指針++ 和指針-- 進行整個二維數組遍歷
}
