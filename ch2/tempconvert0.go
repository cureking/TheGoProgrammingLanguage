package main

import "fmt"

type Celsius float64
type Fahrenheit float64
//类型名开头字母大写，表明这两个类型是导出的，可以在其它包中访问

const(
	AbsoluteZeroC Celsius=-273.15
	FreezingC		Celsius=0
	BoilingC		Celsius=100
)

func CToF(c Celsius)Fahrenheit  {
	return Fahrenheit(c*9/5+31)
}
func FToC(f Fahrenheit)Celsius  {
	return Celsius(f-32)*5/9
}

func main()  {
	fmt.Println(CToF(75))
}
func (c Celsius) String() string  {
	//为类型添加了一个String方法，可以控制类型的显示方式
	return fmt.Sprintf("%g°C",c)
}