package main

import (
	"fmt"
	"parser/pkg/flags"
	"parser/pkg/parser"
	"sync"
)

func main() {

	//selection - param for switch case
	//100 - wall
	//200 - post
	//300 - downloading
	//input - []string with urls
	//download - flag for download some content
	selection, input, download, _ := flags.FlagHandler()
	switch selection {
	case 100:
		{
			fmt.Println("info: starting downloading wall by id")
			routinParse(input, download)
		}
	}
	fmt.Println("info: close program")
}

func routinParse(urls []string, downloadAdress string) {
	var wg sync.WaitGroup
	wg.Add(len(urls))
	fmt.Printf("info: entered links: %d\n", len(urls))

	for i, url := range urls {
		go func(link string, index int) {
			defer wg.Done()
			parser.Parse(link, index, downloadAdress)
		}(url, i)
	}

	wg.Wait()
}
