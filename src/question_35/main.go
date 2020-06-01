package main

func main(){

}
func searchInsert(nums []int, target int) int {
  pos := binarySearch(nums, target)  
  if pos != -1{
    return pos
  }
  slow:=0  
  for i:=1; i< len(nums)-1; i++{
      if target >= nums[slow] && target  <= nums[i]{
          return i+1
      }else if target <= nums[slow]{
          return slow
      }
      slow ++
  }
  return len(nums)
}
func binarySearch(nums []int, target int) int {
  left := 0
  right := len(nums) -1

  for left <= right{
    mid := (left+right) / 2
    if nums[mid] == target{
      return mid
    }else if nums[mid] < target{
      left = mid + 1
    }else if nums[mid] > target{
      right = mid - 1
    }
  }
  return -1
}