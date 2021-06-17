package main
import "unicode"


func main(){

	Solve("tere")
}

func Solve(s string) []int {
	upper:=0
  lower:=0
  number:=0
  special:=0
  result:=[]int{}
  
  for _,v:=range s{
    
    if unicode.IsUpper(v){
      upper++
    }else if unicode.IsLower(v){
      lower++
    }else if unicode.IsNumber(v){
      number++
    }else{
      special++
    }
    
  }
  result=append(result,upper)
  result=append(result,lower)
  result=append(result,number)
  result=append(result,special)
  
  return result
}
