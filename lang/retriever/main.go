package main

import (
	"fmt"
	"richie.com/richie/learn_golang/lang/retriever/real"
	"richie.com/richie/learn_golang/lang/retriever/richie"
	"time"
)

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

const url = "http://www.immoc.com"

func dowload(r Retriever) string {
	return r.Get(url)
}

func post(poster Poster) {
	poster.Post(
		url,
		map[string]string{
			"name":   "richie",
			"course": "golang",
		})
}

type RetrieverPoster interface {
	Retriever
	Poster
}

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{
		"contents": "another faked immoc.com",
	})
	return s.Get(url)
}

func main() {
	var r Retriever
	retriever := richie.Retriever{
		"this is fake richie.com"} // 拷贝值
	r = &retriever
	inspect(r)

	r = &real.Retriever{ // 指针传递
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	//fmt.Println(dowload(r))

	inspect(r)
	// Type assertion
	if richieRetriever, ok := r.(*richie.Retriever); ok {
		fmt.Println(richieRetriever.Contents)
	} else {
		fmt.Println("not a richie retriever")
	}

	fmt.Println("Try a session")
	fmt.Println(session(&retriever))
}

func inspect(r Retriever) {
	fmt.Println("Inspectiong", r)
	fmt.Printf("> %T %v\n", r, r)
	fmt.Print("> Type switch:")
	switch v := r.(type) {
	case * richie.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
}
