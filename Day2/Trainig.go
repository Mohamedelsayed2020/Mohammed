package main
import (
	"fmt"
)
var i,j=1,2
func swap(x,y string)(string,string){
return y,x
}
func main(){
	a,b:=swap("Mohammed","Hello")
	fmt.Println(a,b)
	fmt.Println(spilt(17))
	var c,paython,java =true,false,"no!"
	 var r,t int =1,2
	 o,p,q:=true,false,"nono!!"
	fmt.Println(i,j,c,paython,java)
	fmt.Println(r,t,o,p,q)
}
func spilt(sum int)(x,y int ){

	x=sum*4/9
	y=sum-x
	return
}