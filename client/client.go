package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const baseURL = "https://www.gran-turismo.com/jp/api/gt7sp/"
const profileEndpoint = "profile/"

type profileJSON struct {
	Profile json.RawMessage `json:"profile"`
}

type profileInternalJSON struct {
	ID string `json:"id"`
}

// GetUserName requests user_id from user_no
func GetUserName(userNo int) (result string, err error) {
	client := http.Client{}

	form := url.Values{}
	form.Add("job", "1")
	form.Add("user_no", fmt.Sprint(userNo))
	body := strings.NewReader(form.Encode())

	req, err := http.NewRequest("POST", baseURL+profileEndpoint, body)
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	res, err := client.Do(req)
	if err != nil {
		return
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}

	var profile profileJSON
	if err = json.Unmarshal(resBody, &profile); err != nil {
		return
	}

	var internal profileInternalJSON
	if err = json.Unmarshal(profile.Profile, &internal); err != nil {
		return
	}

	result = internal.ID
	return
}

type statsJSON struct {
	Stats json.RawMessage `json:"stats"`
}

// UserProfile user profile
type UserProfile struct {
	DriverPoint string `json:"driver_point"`
	MannerPoint string `json:"manner_point"`
}

// GetUserProfile returns driver rating and sportsmanship rating of user
func GetUserProfile(userNo int) (result UserProfile, err error) {
	client := http.Client{}

	form := url.Values{}
	form.Add("job", "3")
	form.Add("user_no", fmt.Sprint(userNo))
	body := strings.NewReader(form.Encode())

	req, err := http.NewRequest("POST", baseURL+profileEndpoint, body)
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	res, err := client.Do(req)
	if err != nil {
		return
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}

	var stats statsJSON
	if err = json.Unmarshal(resBody, &stats); err != nil {
		return
	}

	if err = json.Unmarshal(stats.Stats, &result); err != nil {
		return
	}

	return
}
