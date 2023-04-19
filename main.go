package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dghubble/oauth1"
	"github.com/dghubble/oauth1/twitter"
)

const tweetEndpoint = "https://api.twitter.com/2/tweets"

func main() {
	apiKey := os.Getenv("API_KEY")
	apiSecretKey := os.Getenv("API_SECRET_KEY")
	accessToken := os.Getenv("ACCESS_TOKEN")
	accessTokenSecret := os.Getenv("ACCESS_TOKEN_SECRET")

	config := oauth1.Config{
		ConsumerKey:    apiKey,
		ConsumerSecret: apiSecretKey,
		Endpoint:       twitter.AuthorizeEndpoint,
	}

	token := oauth1.NewToken(accessToken, accessTokenSecret)
	httpClient := config.Client(context.Background(), token)

	tweetText := "テスト投稿"

	res, err := postTweet(httpClient, tweetText)
	if res.StatusCode != http.StatusOK {
		fmt.Println(res)
		log.Fatal(err)
	}
}

func postTweet(httpClient *http.Client, text string) (*http.Response, error) {
	data := map[string]string{
		"text": text,
	}
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", tweetEndpoint, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	return httpClient.Do(req)
}
