package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://golang.org",
		"http://amazon.com",
		"http://stackoverflow.com",
	}

	c := make(chan string) // create a new channel of type string

	for _, link := range links {
		go checkLink(link, c) // create a new routine for a http call
	}

	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down :(")
		c <- "might be down"
		return
	}

	fmt.Println(link, "is up")
	c <- "status: ok"
}
