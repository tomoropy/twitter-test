package main

import (
	"log"
	"net/url"
	"os"

	"github.com/ChimeraCoder/anaconda"
)

func main() {
	apiKey := os.Getenv("API_KEY")
	apiSecretKey := os.Getenv("API_SECRET_KEY")
	accessToken := os.Getenv("ACCESS_TOKEN")
	accessTokenSecret := os.Getenv("ACCESS_TOKEN_SECRET")

	anaconda.SetConsumerKey(apiKey)
	anaconda.SetConsumerSecret(apiSecretKey)
	api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)

	tweetText := "Hello, world! This is my first tweet using anaconda library."
	tweet, err := postTweet(api, tweetText)
	if err != nil {
		log.Fatalf("Error posting tweet: %v", err)
	}

	log.Printf("Successfully posted tweet: %s (ID: %s)", tweet.Text, tweet.IdStr)
}

func postTweet(api *anaconda.TwitterApi, text string) (anaconda.Tweet, error) {
	v := url.Values{}
	v.Set("status", text)
	return api.PostTweet(text, v)
}
