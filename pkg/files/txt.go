package files

import (
	"fmt"
	"net/url"
	"os"
	"parser/config"

	"path/filepath"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func ToTXT(body string, imgclass string, index int) {
	//find <img>
	imgSources := findImgSources(body, imgclass)

	resPath, err := config.GetResultDirectory()
	if err != nil {
		fmt.Printf("error: %s\n", err)
		os.Exit(2)
	}

	// .txt res
	file, err := os.Create(filepath.Join(resPath, "image_sources"+strconv.Itoa(index)+".txt"))
	if err != nil {
		fmt.Println("error: can not create "+strconv.Itoa(index)+".txt file with links: ", err)
		return
	}
	defer file.Close()

	for _, source := range imgSources {
		decodedSource, err := url.QueryUnescape(source)
		if err != nil {
			fmt.Println("error: can not decode result: ", err)
			continue
		}

		_, err = file.WriteString(decodedSource + "\n")
		if err != nil {
			fmt.Println("error: can not write in "+strconv.Itoa(index)+".txt file with links: ", err)
			return
		}
	}

	fmt.Println("done: links in \"image_sources" + strconv.Itoa(index) + ".txt\"")
}

func findImgSources(html, className string) []string {
	var sources []string

	// goquery
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		fmt.Println("error: can not parsing html: ", err)
		return sources
	}

	// <img> with class="MediaGrid__imageElement"
	doc.Find("img." + className).Each(func(i int, s *goquery.Selection) {
		src, exists := s.Attr("src")
		if exists {
			sources = append(sources, src)
		}
	})

	return sources
}
