package main

import "fmt"

func main (){
  nums := []int{4,6,7,8,8,9}
  ans:=searchRange(nums, 8)
  fmt.Println(ans)
}
func searchRange(nums []int, target int) []int {
    arr := []int{-1, -1}
    if len(nums) == 0{
        return arr
    }
    if len(nums) == 1 && nums[0] == target{ //只有一个数, 且相等的情况下
        arr[0] = 0
        arr[1] = 0
        return arr
    }
    left := 0
    right := len(nums) - 1
    p := -1
    
    for left <= right { // <= 是因为要应付当只有2个数的时候
        mid := (left+right)/2
        if nums[mid] == target{
            p = mid
            break;
        }else if nums[mid] < target{
            left = mid + 1
        }else{
            right = mid - 1
        }
    }
    if p == -1{
        return arr
    }else{
        r := p + 1
        for p >= 0 && nums[p] == target{
            p --
        }
        for r <= len(nums) -1 && nums[r] == target{
            r++
        }
        arr[1] = r - 1
        arr[0] = p + 1
    }
    
    return arr
}
