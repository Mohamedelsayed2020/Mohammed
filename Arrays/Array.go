package main
import "fmt"
func main(){
	array := [5]int{10, 20, 30, 40, 50}
	varray := [5]int {1: 10 ,2: 20}
	varray[0]=15
	fmt.Println("array is",array)
	fmt.Println("Varray is",varray)
	var arr[5]string
	arr2:=[5]string{"A","B","C","D","E"}
	arr=arr2
	fmt.Println("Arr is ",arr)
}