package main

import (
	"fmt"
	"learning/TheGoProgrammingLanguage/ch2/tempconv"
	//包的来源根路径是从$GOPAHT设定的工作目录开始的。切记。
)

func main()  {
	fmt.Printf("Brrrr! %v\n",tempconv.CustomC)
	fmt.Println(tempconv.CToF(tempconv.CustomC))
	fmt.Println(tempconv.FToC(tempconv.CToF(tempconv.CustomC)))
	fmt.Println(tempconv.CToK(tempconv.CustomC))
}
