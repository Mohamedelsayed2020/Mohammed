package main
import(
	"fmt"
	"math"
	"runtime"
	"time"
)
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}else{
		fmt.Printf("%g >= %g\n", v , lim)
	}
	return lim
}

func sqrt(x float64)string{
	if x<0{
		return sqrt(-x)+"i"
	}
	return fmt.Sprint(math.Sqrt(x))
}
func main(){
	fmt.Println(sqrt(2),sqrt(-4))

	fmt.Println(
		pow(3,2,10),
		pow(3,3,20),
	)
	//Switch
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}
fmt.Println("\n")
	//Switch evaluation order
	fmt.Println("When's Sunday?")
	today:=time.Now().Weekday()
	switch time.Sunday{
	case today+0:
		fmt.Println("Today.")
	case today+1:
		fmt.Println("Tomorrow.")
	default:
		fmt.Println("Too far away.")
	}
fmt.Println("\n")
	//Switch with no condition

	t:=time.Now()
	switch{
	case t.Hour()<12:
		fmt.Println("Good Morning")
	case t.Hour()<17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good Evining")
	}
	fmt.Println("\n")
	//Defer
	defer fmt.Println("World")
	fmt.Println("Hello")
}