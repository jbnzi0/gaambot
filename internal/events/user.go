package events

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

type User struct {
	email    string
	password string
}

type AuthResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

func ConnectBotUsers() []string {
	users := ReadUsersFile()
	var tokens []string

	for i := range users {
		user := users[i]
		accessToken := SignIn(user.email, user.password)
		tokens = append(tokens, accessToken)
	}

	return tokens
}

func SignIn(email string, password string) string {
	apiURL := "https://chess-services.herokuapp.com/v1/auth/login"
	var authResp AuthResponse

	data := url.Values{
		"email":    {email},
		"password": {password},
	}

	res, err := http.PostForm(apiURL, data)

	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	error := json.Unmarshal(body, &authResp)

	if error != nil {
		log.Fatal(err)
	}

	return authResp.AccessToken

}

func ReadUsersFile() []User {
	return []User{
		{email: os.Getenv("DEV_BOT_EMAIL"), password: os.Getenv("DEV_BOT_PWD")},
		{email: os.Getenv("INFO_BOT_EMAIL"), password: os.Getenv("INFO_BOT_PWD")}}
}
