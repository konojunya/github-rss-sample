package main

import (
	"fmt"

	"github.com/mmcdole/gofeed"
)

func main() {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL("https://github.com/konojunya.atom")
	if err != nil {
		panic(err)
	}

	for _, entry := range feed.Items {
		fmt.Println(entry.Title)
	}
}
