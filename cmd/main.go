package main

import (
	"fmt"
	"os"
	"parser/config"
	"parser/pkg/flags"
	"parser/pkg/parser"
	"sync"
)

func main() {

	code, input := flags.FlagHandler()
	switch {
	//парсим
	case code == 100 && input != nil:
		fmt.Println("info: program started")
		multithreadingParse(input)

	//вывод конфига
	case code == 200:
		err := config.PrintYAMLFile("api.yaml")
		if err != nil {
			fmt.Printf("error: can not print config: %s", err)
			os.Exit(1)
		}

	//изменение конфига
	case code == 210 && input != nil:
		fmt.Println(input[0], input[1])
		err := config.UpdateYAMLField("api.yaml", input[0], input[1])
		if err != nil {
			fmt.Printf("error: can not update config: %s", err)
			os.Exit(1)
		}
	}

}

func multithreadingParse(urls []string) {
	var wg sync.WaitGroup
	ch := make(chan struct{}, 3)
	for _, url := range urls {

		//3 запроса в секундку, мб лучше не рисковать...
		wg.Add(6)
		go func(url string) {
			defer wg.Done()
			ch <- struct{}{}
			parser.ParserNew(url, 1)
			<-ch
		}(url)

		go func(url string) {
			defer wg.Done()
			ch <- struct{}{}
			parser.ParserNew(url, 2)
			<-ch
		}(url)

		go func(url string) {
			defer wg.Done()
			ch <- struct{}{}
			parser.ParserNew(url, 3)
			<-ch
		}(url)

		go func(url string) {
			defer wg.Done()
			ch <- struct{}{}
			parser.ParserNew(url, 4)
			<-ch
		}(url)

		go func(url string) {
			defer wg.Done()
			ch <- struct{}{}
			parser.ParserNew(url, 5)
			<-ch
		}(url)

		go func(url string) {
			defer wg.Done()
			ch <- struct{}{}
			parser.ParserNew(url, 6)
			<-ch
		}(url)

		wg.Wait()
	}

}
