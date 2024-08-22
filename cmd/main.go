package main

import (
	"fmt"
	"os"
	"parser/config"
	"parser/pkg/directory"
	"parser/pkg/files"
	"parser/pkg/flags"
	"parser/pkg/parser"
	"sync"
)

func main() {

	code, input := flags.FlagHandler()
	switch {
	//parse
	case code == 100 && input != nil:
		fmt.Println("info: program started")
		directory.CreateFullPath(input)
		multithreadingParse(input)
		err := files.AllInOneJSON(input)
		if err != nil {
			fmt.Println("error: can not make full file: ", err)
		}
		fmt.Println("info: program complete")

	//show config
	case code == 200:
		err := config.PrintYAMLFile("api.yaml")
		if err != nil {
			fmt.Printf("error: can not print config: %s", err)
			os.Exit(1)
		}

	//change config
	case code == 210 && input != nil:
		fmt.Println(input[0], input[1])
		err := config.UpdateYAMLField("api.yaml", input[0], input[1])
		if err != nil {
			fmt.Printf("error: can not update config: %s", err)
			os.Exit(1)
		}
	}

}

func multithreadingParse(domains []string) {
	var wg sync.WaitGroup
	ch := make(chan struct{}, 3)
	for _, domain := range domains {
		//3 request per second
		wg.Add(6)
		go func(domain string) {
			defer wg.Done()
			ch <- struct{}{}
			parser.ParserNew(domain, 1)
			<-ch
		}(domain)

		go func(domain string) {
			defer wg.Done()
			ch <- struct{}{}
			parser.ParserNew(domain, 2)
			<-ch
		}(domain)

		go func(domain string) {
			defer wg.Done()
			ch <- struct{}{}
			parser.ParserNew(domain, 3)
			<-ch
		}(domain)

		go func(domain string) {
			defer wg.Done()
			ch <- struct{}{}
			parser.ParserNew(domain, 4)
			<-ch
		}(domain)

		go func(domain string) {
			defer wg.Done()
			ch <- struct{}{}
			parser.ParserNew(domain, 5)
			<-ch
		}(domain)

		go func(domain string) {
			defer wg.Done()
			ch <- struct{}{}
			parser.ParserNew(domain, 6)
			<-ch
		}(domain)

		wg.Wait()
	}

}
