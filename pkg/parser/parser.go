package parser

import (
	"fmt"
	"io"
	"net/http"
	"parser/config"
	"parser/pkg/files"
)

func Parse(link string, index int, selection int) {
	token, err := config.GetApiToken()
	if err != nil {
		fmt.Println(err)
	}

	version, err := config.GetApiVersion()
	if err != nil {
		fmt.Println(err)
	}

	var url string
	switch selection {
	case 100:
		{
			url = fmt.Sprintf("https://api.vk.com/method/wall.get?access_token=%s&v=%s&domain=%s&extended=1", token, version, link)
		}
	case 200:
		{
			url = fmt.Sprintf("https://api.vk.com/method/wall.getById?access_token=%s&v=%s&posts=%s", token, version, link)
		}

	}

	//get user
	//https://api.vk.com/method/users.get?access_token=токен&v=5.199&user_ids=262818868&fields=aboutactivities,about,blacklisted,blacklisted_by_me,books,bdate,can_be_invited_group,can_post,can_see_all_posts,can_see_audio,can_send_friend_request,can_write_private_message,career,common_count,connections,contacts,city,crop_photo,domain,education,exports,followers_count,friend_status,has_photo,has_mobile,home_town,photo_100,photo_200,photo_200_orig,photo_400_orig,photo_50,sex,site,schools,screen_name,status,verified,games,interests,is_favorite,is_friend,is_hidden_from_feed,last_seen,maiden_name,military,movies,music,nickname,occupation,online,personal,photo_id,photo_max,photo_max_orig,quotes,relation,relatives,timezone,tv,universities,is_verified

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

	files.ToJSON2(string(body), index)

}
