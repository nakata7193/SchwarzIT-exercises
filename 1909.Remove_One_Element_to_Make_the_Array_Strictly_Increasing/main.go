package main

import "fmt"

func isStrictlyIncreasing(nums []int) bool {
	isBigger := true
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] >= nums[i+1] {
			if !isBigger {
				return false
			}
			if i > 0 {
				if nums[i+1] <= nums[i-1] {
					nums[i+1] = nums[i]
				}
			}
			isBigger = false
		}
	}
	return true
}

func main() {
	arr := []int{1, 2, 10, 5, 7, 10, 15}
	fmt.Println(isStrictlyIncreasing(arr))
}
