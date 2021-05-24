package main

import "fmt"

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://golang.org",
		"http://stackoverflow.com",
		"http://amazon.com",
	}

	for _, link := range links {
		fmt.Println(link)
	}
}
