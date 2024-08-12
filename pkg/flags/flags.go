package flags

import (
	"errors"
	"flag"
)

func Geturls() ([]string, error) {
	urlsFlag := flag.String("url", "", "urls for parsing")
	flag.Parse()

	urls := flag.Args()
	if len(urls) == 0 && *urlsFlag == "" {
		return nil, errors.New("err: no urls enter")
	}

	res := make([]string, 0, len(urls)+1)
	res = append(res, urls...)

	return res, nil
}
