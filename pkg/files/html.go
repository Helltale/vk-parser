package files

import (
	"fmt"
	"os"
	"parser/config"
	"path/filepath"
	"strconv"
)

func ToHTML(body string, index int) {

	resPath, err := config.GetResultDirectory()
	if err != nil {
		fmt.Printf("error: %s\n", err)
		os.Exit(2)
	}

	// .txt res
	file, err := os.Create(filepath.Join(resPath, "response"+strconv.Itoa(index)+".html"))
	if err != nil {
		fmt.Println("error: can not create .html file: ", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(body)
	if err != nil {
		fmt.Println("error: can not write in "+strconv.Itoa(index)+".html file: ", err)
		return
	}

	fmt.Println("done: create file for responce \"response" + strconv.Itoa(index) + ".html\"")
}
