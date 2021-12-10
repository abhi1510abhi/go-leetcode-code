
/**
#1929
https://leetcode.com/problems/concatenation-of-array/
*/
package main

func getConcatenation(nums []int) []int {
	result := make([] int,0)
	for i:= 0;i<len(nums); i++{
	  result = append(result,nums[i])
	}
	
  // using spread operator since in appent 2 list wont append
	result = append(result,nums...)
	return result
  }
