package main

import (
	"os"
	"net/http"
	"fmt"
	"io/ioutil"
)

func main()  {
	for _,url:=range os.Args[1:]{
		//通过命令行参数的方式获取网站地址
		resp,err:=http.Get(url)
		//通过网站地址获取信息，并将信息返回（无论是内容，还是错误信息）
		if err!=nil{
			//获取网站信息失败时，采用的方法。
			fmt.Fprintf(os.Stderr,"fetch:%v\n",err)
			//打印返回的错误信息
			os.Exit(1)
			//退出进程，并返回自定义的错误码
		}

		//顺利获取网站信息时，采用的方式
		b,err:=ioutil.ReadAll(resp.Body)
		//读取获取信息的Body部分，并将信息返回（无论是内容，还是错误信息）
		resp.Body.Close()
		//关闭,确保资源不会泄露
		if err!=nil{
			//读取Body失败时，采用的方法
			fmt.Fprintf(os.Stderr,"fetch:reading %s:%v\n",url,err)
			os.Exit(1)
		}
		//打印Body
		fmt.Printf("%s",b)
	}
}
