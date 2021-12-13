
// easy
// https://leetcode.com/problems/jewels-and-stones/

func numJewelsInStones(jewels string, stones string) int {
  result :=0
  for i:=0;i<len(stones);i++{
    // comparing jewels with stones[i] which is in byte to convert wrap it in string
    if strings.Contains(jewels,string(stones[i])) {
      result++
    }
  }
       return result
}
