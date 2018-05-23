package main

import "fmt"

//用于计算x,y的最大公约数
//可以自行证明，用除法是为了更快，选用除数也是为了更快
func gcd(x,y int)int{
	for y!=0{
		x,y=y,x%y
	}
	return  x
}

func main()  {
	fmt.Println(gcd(18,12))
}
