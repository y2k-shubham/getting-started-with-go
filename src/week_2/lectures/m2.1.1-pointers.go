package lectures

import "fmt"
import com "../commons"

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
	com.ShowSep("before init")
	fmt.Printf("%s\n", com.GetPtrStr(ptr, "ptr"))
	com.ShowSep("after init")
	fmt.Printf("%s\n", com.GetPtrStr(ptr, "ptr"))
}

func showAll(when string, x int, y int, ip *int) {
	com.ShowSep(when)
	fmt.Printf("x = %d\ny = %d\n%s\n", x, y, com.GetPtrStr(ip, "ip"))
	com.ShowSep("")
}
