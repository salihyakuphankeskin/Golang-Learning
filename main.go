package main

import	(
	"fmt"
)
	

func main() {
	Part_I()
	fmt.Scanln()
}

func Part_I()  {	
	myInt := arrayCreater(1,2,3,4,5,6,7,8,9,10)
	fmt.Println(myInt)

	myInt2 :=arrayCreaterFirstSix(1,2,3,4,5,6,7,8,9,10)
	fmt.Println(myInt2)
	
}
func arrayCreater(intArray ...int)[]int{
	return intArray
}

func arrayCreaterFirstSix(intArray ...int)[6]int{
	myArray := [6]int{}
	copy(myArray[:], intArray[:6])
	return myArray
}