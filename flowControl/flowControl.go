package flowControl

import (
	"fmt"
	"runtime"
	"time"
)

func Runner(){
	MyForLoop(5,"Kevin")
	DeferTest()
	SystemCheck()
	WhereIsSaturday()
	
}
func MyForLoop(counter int, name string) {
	for i := 0; i < counter; i++ {
		fmt.Println(name," hey you look at me! ", i)
	}
}
func MyWhoileLoop() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
func SystemCheck(){
	
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}
func WhereIsSaturday()  {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()

	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}	
}

func DeferTest(){
	defer fmt.Println("This is the end!")
	fmt.Println("This is the beginning")
	fmt.Println("This is the middle!")
	// Defer waits until every function finishes in this scope
	
}