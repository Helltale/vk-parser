package donwloader

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func download(imageURL string) {

	resp, err := http.Get(imageURL)
	if err != nil {
		fmt.Println("error: can not get new image by url", err)
		return
	}
	defer resp.Body.Close()

	file, err := os.Create("image.jpg")
	if err != nil {
		fmt.Println("error: can not create new image", err)
		return
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		fmt.Println("error: can not save new image", err)
		return
	}

	fmt.Println("done: save image")
}
