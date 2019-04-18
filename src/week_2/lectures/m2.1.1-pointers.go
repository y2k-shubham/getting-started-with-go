package lectures

import "fmt"

func PointersDemo() {
	var x = 1
	var y int
	var ip *int

	showAll("before init", x, y, ip)

	ip = &x
	y = *ip

	showAll("after init", x, y, ip)
}

func PointersNewDemo() {
	ptr := new(int)
	fmt.Println("----------\nbefore init")
	fmt.Printf("*ptr = %s\n", getPtrStr(ptr, "ptr"))
	fmt.Println("----------\nafter init")
	fmt.Printf("*ptr = %s\n", getPtrStr(ptr, "ptr"))
}

func showAll(when string, x int, y int, ip *int) {
	fmt.Println("-----------")
	fmt.Println(when)
	fmt.Printf("x = %d\ny = %d\nip = %p\n*ip = %s\n", x, y, ip, getPtrStr(ip, "ip"))
	fmt.Println("-----------")
}

func getPtrStr(ptr *int, ptrName string) string {
	var ptrStr string
	if ptr == nil {
		ptrStr = "nil"
	} else {
		ptrStr = fmt.Sprintf("%d", *ptr)
	}
	return ptrStr
}
