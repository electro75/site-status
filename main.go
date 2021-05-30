package main

import (
	"fmt"
	"net/http"
	"time"
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

	// create an infinite loop
	// for loop waits for channel to return a value
	// the value returned by the channel is receivd by 'l'
	for l := range c {
		time.Sleep(time.Second * 3)
		go checkLink(l, c)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down :(")
		c <- link
		return
	}

	fmt.Println(link, "is up")
	c <- link
}
