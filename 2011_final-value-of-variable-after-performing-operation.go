/**
#2011
https://leetcode.com/problems/final-value-of-variable-after-performing-operations/
*/
func finalValueAfterOperations(operations []string) int {
  result :=0;
  for _,ele:=range operations {
    if(ele =="X++" || ele == "++X"){
      result++
    }else{
      result--
    }
  }
  
  return result
}
