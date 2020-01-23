package main

import "fmt"

func key (i, j int) string {
  return fmt.Sprintf("%d,%d", i, j)
}

func subarraySumH(nums []int, i, j, k int, fullSum int, seen map[string]bool) int {
  ky := key(i, j)

  if _, saw := seen[ky]; saw {
    return 0
  }

  seen[ky] = true
  var increment int 

  if fullSum == k { 
    increment++
  }

  if  i == j {
    return increment
  }  

  return increment + subarraySumH(nums, i, j-1, k, fullSum-nums[j], seen) + subarraySumH(nums, i+1, j, k, fullSum-nums[i], seen)

}

func subarraySum(nums []int, k int) int {
    seen := make(map[string]bool, len(nums))
    var fullSum int

    for _, num := range nums {
      fullSum+=num
    }

    count := subarraySumH(nums, 0, len(nums)-1, k, fullSum, seen)

    return count
}

func main() {
  fmt.Printf("subArray count %d\n", subarraySum([]int{-19, -82, -70, -46, -29, 7, 15, -72, -7, 100, 303}, 100))
}