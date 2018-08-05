package main
import(
	"fmt"
	"math/cmplx"
	"math"
)

//basic types
var (
	ToBe bool =false
	MaxInt uint64 =1<<64-1
	z complex128=cmplx.Sqrt(-5+12i)
)

//Contsants
const(
	Big=1<<100
	Small=Big>>99
)
func needInt(x int )int{
	return x*10+1
}
func needFloat(x float64)float64{
	return x*0.1
}

func main(){
	//basic types
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n",MaxInt,MaxInt)
	fmt.Printf("Type: %T Value: %v\n",z,z)
	//Zero Values 
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
	
	//type conversion
	var x,y int =3,4
	var d float64=math.Sqrt(float64(x*x + y*y))
	var z uint=uint(d)
	fmt.Println(x,y,z)

	//type inference
	u:=42
	fmt.Printf("u is type of %T\n",u)

	//Constants
	const Pi=3.14
	fmt.Println("Happy",Pi,"Day")
	//Constants
	fmt.Println("First : ",needInt(Small))
	fmt.Println("Second : ",needFloat(Small))
	fmt.Println("Third : ",needFloat(Big))
	//Pointers
	 k,l :=5,7
	 m:=&k
	fmt.Println("value of K is : ",*m,"and Poninter of L",&l)
}