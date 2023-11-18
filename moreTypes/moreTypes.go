package moreTypes

import "fmt"

func Runner() {
	PointerFirstEx()
}
func PointerFirstEx() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}
type Coordinates struct {
    X FloatAndIntegerFluidType
    Y FloatAndIntegerFluidType
    Z FloatAndIntegerFluidType
}
type FloatAndIntegerFluidType interface {
    ToFloat64() float64
}  
func (i int) ToFloat64() float64 {
    return float64(i)
}
func (f float64) ToFloat64() float64 {
    return f
}

