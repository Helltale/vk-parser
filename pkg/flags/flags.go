package flags

import (
	"errors"
	"flag"
)

// func Geturls() ([]string, error) {
// 	urlsFlag := flag.String("url", "", "urls for parsing")
// 	flag.Parse()

// 	urls := flag.Args()
// 	if len(urls) == 0 && *urlsFlag == "" {
// 		return nil, errors.New("err: no urls enter")
// 	}

// 	res := make([]string, 0, len(urls)+1)
// 	res = append(res, urls...)

// 	return res, nil
// }

func Geturls() ([]string, bool, error) {
	urlsFlag := flag.String("url", "", "urls for parsing")
	downloadFlag := flag.Bool("download", false, "download the URLs")

	flag.Parse()

	if *urlsFlag == "" {
		return nil, false, errors.New("err: no URLs entered")
	}

	if *downloadFlag && *urlsFlag == "" {
		return nil, false, errors.New("err: -download flag can only be used with -url flag")
	}

	urls := flag.Args()
	res := make([]string, 0, len(urls)+1)
	res = append(res, *urlsFlag)
	res = append(res, urls...)

	return res, *downloadFlag, nil
}
