package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	c := make(chan string)

	links := []string{
		"http://facebook.com",
		"http://golang.org",
		"http://google.com",
	}

	for _, link := range links {
		go checkLink(link, c)
	}

	//v1
	for l := range c {
		go func(link string) {
			time.Sleep(time.Second)
			checkLink(link, c)
		}(l)
	}

	//v2
	//time.Sleep(time.Second)
	//go checkLink(l,c)

	//fuction literal: tao nhieu goroutine ma k can phai doi
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down")
		c <- link
		return
	}
	fmt.Println(link, "is up")
	c <- link
}

//for l := range c {
//	go func() {
//		time.Sleep(time.Second)
//		checkLink(l,c)
//	}()
//}
