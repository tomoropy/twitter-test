package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dghubble/oauth1"
	"github.com/dghubble/oauth1/twitter"
)

const tweetEndpoint = "https://api.twitter.com/1.1/statuses/update.json"

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

	tweetText := "Hello, world! This is my first tweet using a custom Go Twitter client."

	if _, err := postTweet(httpClient, tweetText); err != nil {
		log.Fatalf("Error posting tweet: %v", err)
	}

	fmt.Printf("Successfully posted tweet: %s\n", tweetText)
}

func postTweet(httpClient *http.Client, text string) (*http.Response, error) {
	form := make(map[string][]string)
	form["status"] = []string{text}
	return httpClient.PostForm(tweetEndpoint, form)
}
