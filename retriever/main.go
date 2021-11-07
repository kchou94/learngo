package main

import (
	"fmt"
	"retriever/mock"
	"retriever/real"
	"time"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever, url string) string {
	return r.Get(url)
}

func main() {
	var r Retriever
	r = mock.Retriever{"this is a fake news"}
	fmt.Printf("%T %v\n", r, r)

	r = real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	fmt.Printf("%T %v\n", r, r)
	// fmt.Println(download(r, "http://www.baidu.com"))
}
