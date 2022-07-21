package main
import "fmt"

func main(){
	// Basic data types
	var num int = 4
	const pI float32 = 3.14;
	const s string = "hello world"
	var isBoolean bool = true

	fmt.Println("num is ",num)
	fmt.Println("pI is ",pI)
	fmt.Println("s is",s)
	fmt.Println("isBoolean is",isBoolean)


	var nums [3]int;
	fmt.Println(nums[1])

	fabonacci := [5]int{0,1,1,2,3}

	fmt.Println(fabonacci)
	fmt.Println(len(fabonacci))

	type car struct{
		engine string
		hp int
	}

	mustang := car{engine:"v8"}
	fmt.Println(mustang)
	fmt.Println(mustang.engine)

	mapsTest()

}