package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"https://www.bukalapak.com/",
		"https://www.shopee.co.id/",
		"https://www.google.com/",
		"https://www.tokopedia.com/",
		"https://www.facebook.com/",
		"https://www.stackoverflow.com/",
		"https://www.golang.org/",
		"https://www.amazon.com/",
		"https://www.olx.co.id/",
		"https://id.yahoo.com/",
		"https://www.bumble.com/",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for {
		go checkLink(<-c, c)
	}

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}
