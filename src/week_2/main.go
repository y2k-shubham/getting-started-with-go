package main

import "./lectures"
import "./commons"

func main() {
	commons.ShowSep("Pointers")
	lectures.PointersDemo()
	lectures.PointersNewDemo()

	commons.ShowSep("Garbage-Collection")
	lectures.FooBar()
}
