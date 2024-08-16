package archiving

import (
	"time"
)

type Wall struct {
	Name string
	Post Post
}

type Post struct {
	Id   string
	Date string
	Text string
	Img  Img
}

type Img struct {
	Url          string
	WebViewToken string
}

func WallToArchiving(name string, body string) {

}

func UtsToDt(uts int) string {
	unixTimestamp := 1646827986
	return time.Unix(int64(unixTimestamp), 0).Format("2006-01-02 15:04:05")
}
