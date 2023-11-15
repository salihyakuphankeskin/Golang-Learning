package flowControl

import "fmt"

func MyForLoop(counter int) {
	for i := 0; i < counter; i++ {
		fmt.Println("This is the loop ", i)
	}
}