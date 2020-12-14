package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	var showIndex = 0 //i
	if len(nums) < 2 {
		return len(nums)
	}
	for fastIndex := 1; fastIndex < len(nums); fastIndex++ {
		if nums[showIndex] != nums[fastIndex] {
			showIndex++
			nums[showIndex] = nums[fastIndex]
			fmt.Println("***********start****************")
			fmt.Println("fastIndex=", fastIndex)
			fmt.Println("showIndex=", showIndex)
			fmt.Println("***********end****************")
		} else {
			fmt.Println("===========start=============")
			fmt.Println("fastIndex=", fastIndex)
			fmt.Println("showIndex=", showIndex)
			fmt.Println("===========end=============")
		}
	}

	return showIndex + 1
}

func main() {
	var nums = []int{1, 2}
	count := removeDuplicates(nums)
	fmt.Println("count=", count)
}
