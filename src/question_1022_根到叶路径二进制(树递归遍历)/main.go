package main

func main() {
	
}

type TreeNode struct {
     Val int
     Left *TreeNode
Right *TreeNode
}

func sumRootToLeaf(root *TreeNode) int {
	return getSum(root, 0)
}

func getSum(root *TreeNode, sum int) int {
	if root == nil{
		return 0
	}
	// sum 相当于 每次 二进制 左移1位
	// 例如 1 0 0 = 4
	//     (0*2)+1   (1*2)+0 (2*2)+0
	sum = sum * 2 + root.Val

	if root.Left == nil && root.Right == nil{
		return sum
	}

	return getSum(root.Left, sum) + getSum(root.Right, sum)
}