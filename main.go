package main

import "fmt"

func productExceptSelf(nums []int) []int {

  result := make([]int, len(nums))
  if len(nums) == 0 {
    return result
  }
  currMult := nums[len(nums)-1]

  for i := 0; i < len(result); i++ {
    result[i] = 1
  }

  for i := len(nums)-2; i >= 0; i-- {
    tmp := nums[i]
    result[i] = currMult
    currMult *= tmp
  }

  currMult = nums[0]

  for i := 1; i < len(nums); i++ {
    result[i] *= currMult
    currMult*=nums[i]
  }

  return result
}

func main() {
  fmt.Printf("productExceptSelf %v\n", productExceptSelf([]int{1,2,3,4}))
}