package main

import (
	"fmt"
	"os"
	"parser/config"
	"parser/pkg/parser"
	"sync"
)

func main() {
	imgclass, err := config.GetCassNameImg()
	if err != nil {
		fmt.Printf("error: %s\n", err)
		os.Exit(2)
	}

	fmt.Println(imgclass)

	// urls, _, err := flags.Geturls()
	urls := []string{"nickless47"}
	if err != nil {
		fmt.Printf("error: %s", err)
	}

	var wg sync.WaitGroup
	wg.Add(len(urls))
	fmt.Printf("enter links: %d\n", len(urls))

	for i, url := range urls {
		go func(link string, index int) {
			defer wg.Done()
			parser.Parse(link, index, imgclass)
		}(url, i)
	}

	wg.Wait()
}
