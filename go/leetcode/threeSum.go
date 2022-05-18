package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSum([]int{0, 0, 0, 0}))
	fmt.Println(threeSum1([]int{0, 0, 0, 0}))
}

// 解法一 最优解，双指针 + 排序
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result, start, end, index, addNum, length := make([][]int, 0), 0, 0, 0, 0, len(nums)
	for index = 1; index < length-1; index++ {
		start, end = 0, length-1
		if index > 1 && nums[index] == nums[index-1] {
			start = index - 1
		}
		for start < index && end > index {
			if start > 0 && nums[start] == nums[start-1] {
				start++
				continue
			}
			if end < length-1 && nums[end] == nums[end+1] {
				end--
				continue
			}
			addNum = nums[start] + nums[end] + nums[index]
			if addNum == 0 {
				result = append(result, []int{nums[start], nums[index], nums[end]})
				start++
				end--
			} else if addNum > 0 {
				end--
			} else {
				start++
			}
		}
	}
	return result
}

// 解法一 最优解，双指针 + 排序
func threeSum1(nums []int) [][]int {
	sort.Ints(nums)
	result, addNum, length := make([][]int, 0), 0, len(nums)
	target := 0
	for index, endIndex := 1, length-1; index < endIndex; index++ {
		start, end := 0, endIndex
		dbIndex := nums[index] << 1
		if nums[index] == nums[index-1] {
			start = index - 1
		}
		for start < index && end > index {
			if (nums[start] == nums[start+1] && start+1 != index) || nums[start]+dbIndex > target {
				start++
				continue
			}
			if (nums[end] == nums[end-1] && end-1 != index) || nums[end]+dbIndex < target {
				end--
				continue
			}
			addNum = nums[start] + nums[end] + nums[index]
			if addNum == target {
				result = append(result, []int{nums[start], nums[index], nums[end]})
				start++
				end--
			} else if addNum > target {
				end--
			} else {
				start++
			}
		}
	}
	return result
}
