// #easy
// https://leetcode.com/problems/richest-customer-wealth/

func maximumWealth(accounts [][]int) int {
  result:=0
  sum :=0
  for i:=0;i<len(accounts);i++ {
    for j:=0;j<len(accounts[i]);j++{
      sum = sum+accounts[i][j];
    }
    if result < sum{
      result = sum
    }
    sum=0
  }
  return result
}
