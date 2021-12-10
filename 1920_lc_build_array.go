// https://leetcode.com/problems/build-array-from-permutation/

package main

func buildArray(nums []int) []int {

	result := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		result = append(result, nums[nums[i]])
	}
	return result
}
