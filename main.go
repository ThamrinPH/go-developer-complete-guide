package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"https://www.google.com/",
		"https://www.facebook.com/",
		"https://www.stackoverflow.com/",
		"https://www.golang.org/",
		"https://www.amazon.com/",
		"https://www.tokopedia.com/",
		"https://www.shopee.com/",
	}

	for _, link := range links {
		resp, err := http.Get(link)

		if err != nil {
			fmt.Println("Error :", err)
		}

		fmt.Println(resp)
	}
}
