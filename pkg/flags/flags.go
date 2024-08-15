package flags

import (
	"flag"
	"os"
	"strings"
)

func FlagHandler() (int, []string) {
	wall := flag.String("wall", "", "generate response{N}.json of wall")
	post := flag.String("post", "", "generate response{N}.json of post")

	download := flag.Bool("download", false, "start to download media")
	fileFlag := flag.String("file", "", "file with urls for downloading")
	// link := flag.String("link", "", "link for download media")

	flag.Parse()

	if *wall == "" && *post == "" && !*download {
		flag.Usage()
		os.Exit(1)
	}

	if *wall != "" {
		wallSlice := strings.Split(*wall, " ")
		return 100, wallSlice
	}
	if *post != "" {
		postSlice := strings.Split(*post, " ")
		return 200, postSlice
	}
	if *download {
		if *fileFlag != "" { //problem?
			fileFlagSlice := strings.Split(*fileFlag, " ")
			return 311, fileFlagSlice //-download -file c:/dir/file.txt (ссылки из конкретного файла)
		} else {
			return 310, nil //-download -file (ссылки из файла по умолчанию)
		}
	} else {
		downloadSlice := strings.Split(*post, " ")
		return 321, downloadSlice //-download example.com/file.jpg (напрямую ссылки)
	}
}
