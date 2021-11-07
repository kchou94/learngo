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

type Poster interface {
	Post(url string, form map[string]string) string
}

type RetrieverPoster interface {
	Retriever
	Poster
}

const url = "http://www.baidu.com"

func download(r Retriever, url string) string {
	return r.Get(url)
}

func inspect(r Retriever) {
	fmt.Println("Inspect", r)

	fmt.Printf("%T %v\n", r, r)

	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}

	fmt.Println()
}

func Post(poster Poster) {
	poster.Post(url, map[string]string{
		"name":   "ccmouse",
		"course": "golang",
	})
}

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{
		"contents": "another faked news",
	})
	return s.Get(url)
}

func main() {
	var r Retriever
	retriever := mock.Retriever{"this is a fake news"}
	r = &retriever
	inspect(r)

	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}

	inspect(r)

	// Type assertion
	if realRetriever, ok := r.(*mock.Retriever); ok {
		fmt.Println(realRetriever.Contents)
	} else {
		fmt.Println("not a mock retriever")
	}

	// fmt.Println(download(r, "http://www.baidu.com"))

	fmt.Println("Try a session")
	fmt.Println(session(&retriever))
}
