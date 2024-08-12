package main

import (
	"context"
	"fmt"
	"os"
	"parser/config"
	"parser/pkg/files"
	"parser/pkg/flags"
	"sync"

	"github.com/chromedp/chromedp"
)

func main() {
	imgclass, err := config.GetCassNameImg()
	if err != nil {
		fmt.Printf("error: %s\n", err)
		os.Exit(2)
	}

	fmt.Println(imgclass)

	urls, err := flags.Geturls()
	if err != nil {
		fmt.Printf("error: %s", err)
	}

	var wg sync.WaitGroup
	wg.Add(len(urls))
	fmt.Printf("enter links: %d\n", len(urls))

	for i, url := range urls {
		go func(link string, index int) {
			defer wg.Done()
			parser(link, index, imgclass)
		}(url, i)
	}

	wg.Wait()
}

func parser(link string, index int, imgclass string) {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var body string
	err := chromedp.Run(ctx, chromedp.Navigate(link), chromedp.OuterHTML("html", &body))
	if err != nil {
		fmt.Println("error: can not get url: ", err)
		return
	}

	files.ToHTML(body, index)

	files.ToTXT(body, imgclass, index)
}
