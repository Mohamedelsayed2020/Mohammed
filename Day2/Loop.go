package main
import "fmt"
func main(){
	sum:=0
	for i:=0;i<10;i++{
		sum+=i;
	}
	fmt.Println("sum is  : ",sum)

	s:=1
	for ; s<1000;{
		s+=s
	}
	fmt.Println("sum = ",s)

	x:=1
	for x<1000{
		x+=x
	}	
		fmt.Println("sum is = ",x)

}