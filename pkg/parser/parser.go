package parser

import (
	"fmt"
	"io"
	"net/http"
	"parser/config"
	"parser/pkg/files"
)

func Parse(link string, index int, downloadAdress string) {
	token, err := config.GetApiToken()
	if err != nil {
		fmt.Println(err)
	}

	// key, err := config.GetApiKey()
	// if err != nil {
	// 	fmt.Println(err)
	// }

	version, err := config.GetApiVersion()
	if err != nil {
		fmt.Println(err)
	}

	url := fmt.Sprintf("https://api.vk.com/method/wall.get?access_token=%s&v=%s&domain=%s", token, version, link)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		fmt.Println(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	files.ToJSON(string(body), index)

}
