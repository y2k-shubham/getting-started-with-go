package assignment_findian

import "fmt"
import "strings"

func Findian() {
	var str string

	fmt.Printf("Enter a string: ")
	fmt.Scan(&str)

	var strLower = strings.ToLower(str)
	if strings.HasPrefix(strLower, "i") && strings.ContainsRune(strLower, 'a') && strings.HasSuffix(strLower, "n") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Now Found!")
	}
}
