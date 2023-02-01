package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"time"
)

const (
	UNSPLASH_API_KEY = "52cf_fhzMkca1xEohuF6CFsxc0i06f3tksGk4EL6HY8"
	unsplashApiURL   = "https://api.unsplash.com/search/photos"
	FIREBASE_API_KEY = "AIzaSyBVu89jQsohHOyoGtCb2tFgCYBg2qfmAO8"
)

type UnsplashResponse struct {
	Total      int `json:"total"`
	TotalPages int `json:"total_pages"`
	Results    []struct {
		Id   string `json:"id"`
		Urls struct {
			Raw     string `json:"raw"`
			Full    string `json:"full"`
			Regular string `json:"regular"`
			Small   string `json:"small"`
			Thumb   string `json:"thumb"`
		} `json:"urls"`
	} `json:"results"`
}

func FetchImage(eventType string) string {
	// eventType = url.QueryEscape(eventType)
	fmt.Println("eventType:", url.QueryEscape(eventType))
	var res UnsplashResponse
	url := unsplashApiURL + "?query=" + eventType

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Authorization", "Client-ID "+UNSPLASH_API_KEY)

	response, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(body, &res)

	if res.Total == 0 || res.TotalPages == 0 {
		return "https://www.google.com/url?sa=i&url=https%3A%2F%2Fwww.npr.org%2F2022%2F10%2F26%2F1131622796%2Ftheme-holiday-party-planning-tips&psig=AOvVaw3Ix5Ks0-ULMQ5YFyRobEvR&ust=1675300731665000&source=images&cd=vfe&ved=0CBAQjRxqFwoTCIim3tmT8_wCFQAAAAAdAAAAABAE"
	}

	fmt.Println(res)
	rand.Seed(time.Now().UnixNano())

	min := 0
	max := len(res.Results)

	return res.Results[rand.Intn(max-min+1)+min].Urls.Raw
}

func uploadToCloudStorage(eventId, url string) {
	res, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	//TODO: Download image from url and save it to cloud storage in event bucket
}
