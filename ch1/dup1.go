//dup1输出标准输入中出现次数大于1的行。前面是次数，后面是出现的字符
package main

import (
	"bufio"
	"os"
	"fmt"
)

func main(){
	counts:=make(map[string]int)
	//map是用来存储一个键/值对集合
	//其中键为：能够进行相等（==）比较的任意类型，常见为字符串。上例为字符串
	//其中值为：任意类型。上例为int
	//可以类比JS中数组（也可以说对象，毕竟JS中数组也只是特殊的对象）

	//make可以用来新建map，当然，其还有其他用途。

	input:=bufio.NewScanner(os.Stdin)
	//bufio是一个用来简便高效处理输入输出的包
	// 其中扫描器(Scanner）的类型可以用来读取输入。并以行或者单词为单位断开
	//该处新建了一个bufio.Scanner类型的input变量

	for input.Scan(){
		//input.Scan()读取下一行，并将结尾的换行符去掉
		//input.Scan()读取到新行时，返回true，否则，返回false

		//?怎样结束输入，连续回车无法解决。
		counts[input.Text()]++
		//input.Text()来获取读到的内容
	}

	for line,n:=range counts{
		//每一次迭代，range都会产生一对值
		if n>1{
			fmt.Printf("%d\t%s\n",n,line)
		}
	}
}
