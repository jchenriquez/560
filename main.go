package main

import "fmt"

func key (i, j int) string {
  return fmt.Sprintf("%d,%d", i, j)
}

func subarraySumH(nums []int, i int, k int) int {

  if i >= len(nums) {
    return 0
  }

  if k - nums[i] == 0 {
    return 1
  }

  return subarraySumH(nums, i+1, k) + subarraySumH(nums, i+1, k-nums[i])
}

func subarraySum(nums []int, k int) int {
  return subarraySumH(nums, 0, k)
}

func main() {
  fmt.Printf("subArray count %d\n", subarraySum([]int{-19, -82, -70, -46, -29, 7, 15, -72, -7, 100, 303}, 100))
}