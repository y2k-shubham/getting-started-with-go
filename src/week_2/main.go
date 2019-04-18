package main

import "./lectures"
import "./commons"
import "./assignment_trunc"
import "./assignment_findian"

func main() {
	commons.ShowSep("Pointers")
	lectures.PointersDemo()
	lectures.PointersNewDemo()

	commons.ShowSep("Garbage-Collection")
	lectures.FooBar()

	commons.ShowSep("Assignment-Trunc")
	assignment_trunc.Trunc()

	commons.ShowSep("Findian")
	assignment_findian.Findian()
}
