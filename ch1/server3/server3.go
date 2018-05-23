package main

import (
	"net/http"
	"log"
	"fmt"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter,r *http.Request)  {
	fmt.Fprintf(w,"%s %s %s\n",r.Method,r.URL,r.Proto)
	for k,v:=range r.Header{
		fmt.Fprintf(w,"Header[%q] = %q\n",k,v)
	}
	fmt.Fprintf(w,"Host = %q\n",r.Host)
	fmt.Fprintf(w,"RemoteAddr = %q\n",r.RemoteAddr)
	//fmt.Fprintf(w,"response = %q\n",r)
	if err:=r.ParseForm();err!=nil{
		//这里在if条件前添加了一个简单的语句，并用;隔开它与条件
		//这样的写法不仅简便，而且缩小了err的作用域，杜绝 了其对全局环境变量的污染。
		log.Print(err)
	}
	for k,v:=range r.Form{
		fmt.Fprintf(w,"Form[%q] = %q\n",k,v)
	}
}
