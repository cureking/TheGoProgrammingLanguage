package main

import (
	"time"
	"os"
	"fmt"
	"net/http"
	"io"
	"io/ioutil"
)

func main() {
	start := time.Now()
	//设定整个程序开始的时间
	ch := make(chan string)
	//通过make创建一个字符串通道ch
	//通道是一种允许某一例程（线程）向另一个例程传递指定类型的值的通信机制
	//chan是一个FIFO队列，chan分成两种类型同步和异步
	//详情后观，或者访问https://studygolang.com/articles/3107
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
		//通过go语句，启动一个额外的goroutine
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
		//从通道ch接收
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		//将错误信息返回给通道ch
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	//通过io.Copy读取返回的内容，并将内容写入到ioutil.Discard输出流进行丢弃。
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}
//信息返回的顺序并非我们给出的参数网站顺序，而是按照返回时间的顺序排列的
//如果不翻墙的，https://golang.org节点无法返回信息。可以换一个节点。或者翻墙。
