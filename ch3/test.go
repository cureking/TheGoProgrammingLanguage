package main

import (
	"fmt"
	"learning/TheGoProgrammingLanguage/ch3/basename1"
	"learning/TheGoProgrammingLanguage/ch3/basename2"
	"learning/TheGoProgrammingLanguage/ch3/test"
)

func main() {
	x := 4
	y := 3 * +x
	z := `"'/\jkl`
	if "abcd" > "aadf" {
		fmt.Println(y, z)
		fmt.Println(string(65))
		fmt.Println(test.SayHello())
		fmt.Println(basename.Basename("a/c/d/d.go"))
		fmt.Println(basename2.Basement("a/c/d/cd.go"))
	}

}
