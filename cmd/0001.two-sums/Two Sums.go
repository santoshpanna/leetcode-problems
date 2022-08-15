package main

import "fmt"

func twoSums(nums []int, target int) []int {
	var res []int

	var hashMap = make(map[int] int)

	for i:=0; i<len(nums); i++ {
		var diff int = target - nums[i]

		if _, ok := hashMap[diff]; ok {
			return []int{hashMap[diff], i}
		} else {
			hashMap[nums[i]] = i
		}
	}

	return res
}

func main() {
	var nums []int = []int{2, 7, 11, 15}
	var target int = 9

	fmt.Println(twoSums(nums, target))
}