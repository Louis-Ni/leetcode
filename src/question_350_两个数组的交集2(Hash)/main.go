package main

func main() {

}

func intersect(nums1 []int, nums2 []int) []int {
	m := map[int]int{}
	k := 0
	// 使用hash影射的方式
	// 将nums1出现过的值的次数 以 hash方式存储起来。
	// 例如 4:2 5:1 6:1 这样的一个数组
	for _, v := range nums1 {
		m[v] += 1
	}

	// 遍历nums2 并且在map中找到隐射的值有多少， 例如 4 出现在hash中的次数是2次，那么就把4加入到 result中, 并且 把map[key] 的数值--
	for _, v := range nums2 {
		if m[v] > 0 {
			m[v]--
			nums2[k] = v
			k++
		}
	}
	return nums2[0:k]
}
