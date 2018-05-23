package main

import "fmt"

func main()  {
	var p=f()
	fmt.Println(p,p,f(),f())
	fmt.Println(p,p,f(),f())
}

func f()*int{
	v:=1
	return &v
}

//连续运行多次，会发现有意思的地方。