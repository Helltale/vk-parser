package parser

import (
	"fmt"
	"io"
	"net/http"
	"parser/config"
	"parser/pkg/files"
	"time"
)

func ParserNew(domain string, index int) {
	token, err := config.GetApiToken()
	if err != nil {
		fmt.Println(err)
	}

	version, err := config.GetApiVersion()
	if err != nil {
		fmt.Println(err)
	}

	var url string
	fnparts := make([]string, 2)
	fnparts[0] = domain
	switch {
	//user
	case index == 1:
		fnparts[1] = "user"
		url = fmt.Sprintf("https://api.vk.com/method/users.get?access_token=%s&v=%s&user_ids=%s&fields=aboutactivities,about,blacklisted,blacklisted_by_me,books,bdate,can_be_invited_group,can_post,can_see_all_posts,can_see_audio,can_send_friend_request,can_write_private_message,career,common_count,connections,contacts,city,crop_photo,domain,education,exports,followers_count,friend_status,has_photo,has_mobile,home_town,photo_100,photo_200,photo_200_orig,photo_400_orig,photo_50,sex,site,schools,screen_name,status,verified,games,interests,is_favorite,is_friend,is_hidden_from_feed,last_seen,maiden_name,military,movies,music,nickname,occupation,online,personal,photo_id,photo_max,photo_max_orig,quotes,relation,relatives,timezone,tv,universities,is_verified", token, version, domain)
	//video
	case index == 2:
		fnparts[1] = "video"
		url = fmt.Sprintf("https://api.vk.com/method/video.get?access_token=%s&v=%s&owner_id=%s&extended=1", token, version, fnparts[0])
	//photo-wall
	case index == 3:
		fnparts[1] = "photo-wall"
		url = fmt.Sprintf("https://api.vk.com/method/photos.get?access_token=%s&v=%s&owner_id=%s&album_id=wall&extended=1", token, version, fnparts[0])
	//photo-profile
	case index == 4:
		fnparts[1] = "photo-profile"
		url = fmt.Sprintf("https://api.vk.com/method/photos.get?access_token=%s&v=%s&owner_id=%s&album_id=profile&extended=1", token, version, fnparts[0])
	//photo-saved
	case index == 5:
		fnparts[1] = "photo-saved"
		url = fmt.Sprintf("https://api.vk.com/method/photos.get?access_token=%s&v=%s&owner_id=%s&album_id=saved&extended=1", token, version, fnparts[0])
	//wall
	case index == 6:
		fnparts[1] = "wall"
		url = fmt.Sprintf("https://api.vk.com/method/wall.get?access_token=%s&v=%s&domain=%s&extended=1", token, version, fnparts[0])
	}

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

	err = files.ToJSON2(string(body), []string{fnparts[0], fnparts[1]})
	if err != nil {
		fmt.Printf("error: with creating file %s \n", err)
	}

	time.Sleep(1 * time.Second) // для vk api
}
