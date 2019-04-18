package commons

import "fmt"

func GetPtrStr(ptr *int, ptrName string) string {
	var ptrValStr string
	if ptr == nil {
		ptrValStr = "nil"
	} else {
		ptrValStr = fmt.Sprintf("%d", *ptr)
	}

	var ptrStr = fmt.Sprintf("%s = %p,\t*%s = %s", ptrName, ptr, ptrName, ptrValStr)
	return ptrStr
}

func ShowSep(msg string) {
	fmt.Println("-------------")
	if msg != "" {
		fmt.Println(msg)
	}
}
