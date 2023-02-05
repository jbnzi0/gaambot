package unsplash

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"time"
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
	var (
		unsplashAPIKey = os.Getenv("UNSPLASH_API_KEY")
		unsplashApiURL = os.Getenv("UNSPLASH_API_URL")
	)

	var res UnsplashResponse
	url := unsplashApiURL + "?query=" + url.QueryEscape(eventType)

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Authorization", "Client-ID "+unsplashAPIKey)

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

	rand.Seed(time.Now().UnixNano())

	min := 0
	max := len(res.Results) - 1

	return res.Results[rand.Intn(max-min+1)+min].Urls.Raw
}

func DownloadImage(url string) []byte {
	response, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		log.Println(response.Status)
	}

	body, err := io.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	return body
}
