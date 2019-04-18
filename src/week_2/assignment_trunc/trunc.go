package assignment_trunc

import "fmt"

func Trunc() {
	var f float32
	fmt.Printf("Enter a floating-point number: ")
	fmt.Scan(&f)
	fmt.Printf("The entered number after truncating is %d\n", int(f))
}
