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
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	// Creating a channel that will communicate with type string
	c := make(chan string)

	for _, l := range links {
		go checkLink(l, c)
	}

	// When range keyword is used on a channel it waits for the channel to return a value
	// That value is assigned to l. For loop runs with each value received by the channel
	for l := range c {
		// Function literal syntax + invocation
		go func(link string) {
			time.Sleep(1 * time.Second)
			checkLink(link, c)
		}(l)
	}

	// for i := 0; i < len(links); i++ { do stuff }
	// for counter < limit { do stuff }
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down")
		c <- link
		return
	}

	fmt.Println(link, "is up and running.")
	c <- link
}

/*
// for without qualifiers is an infinite loop
for {
	go checkLink(<-c, c)
}
*/
