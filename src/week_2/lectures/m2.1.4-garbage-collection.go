package lectures

import "fmt"
import com "../commons"

func foo() *int {
	x := 1
	fmt.Println(com.GetPtrStr(&x, "&x"))
	return &x
}

func FooBar() {
	var y *int
	y = foo()
	fmt.Println(com.GetPtrStr(y, "y"))
}
