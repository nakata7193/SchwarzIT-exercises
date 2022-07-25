package main

import (
	"fmt"
)

func thirdMax(nums []int) int {
	var max,second,third int
	for i := 0; i < len(nums); i++ {
		if nums[i] == max || nums[i] == second || nums[i] == third {
			continue
		}
		if nums[i] > max {
			third = second
			second = max
			max = nums[i]
		} else if nums[i] > second {
			third = second
			second = nums[i]
		} else if nums[i] > third {
			third = nums[i]
		}>
	}
	if second <= 0 || third <= 0 {
		return max
	}
	return third
}

func main() {
	nums := []int{1,2,3,4,5}
	fmt.Println(thirdMax(nums))
}