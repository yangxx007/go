package main
import "fmt"
func main() {
	var a=1
	var b=2
	a,b=b,a//直接就可以实现两个变量的交换
	fmt.Printf("%d",a)
	fmt.Printf("%d\n",b)
}


