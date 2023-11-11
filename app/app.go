package app
import "math"

func Adder(a,b int) int {
	return a+b	
}

func HighOrder(a,b int) int {
	c:= math.Pow(float64(a), float64(b))
	return int(c)		
}