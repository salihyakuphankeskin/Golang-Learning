package mathPlus

import (
	"fmt"
	"math"
)

func Runner()  {
	c:= Adder(1,2)
	d:= Substracter(3,2)
	e:= HighOrder(2,3)
	fmt.Println(c," and ", d, " and ", e)

	
}
func Adder(a,b int) int {
	return a+b	
}

func HighOrder(a,b int) int {
	c:= math.Pow(float64(a), float64(b))
	return int(c)		
}
func Substracter(a,b int) int {
	return a-b	
}