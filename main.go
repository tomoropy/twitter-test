package main

import (
	"context"
	"log"
	"os"

	"github.com/sivchari/gotwtr"
)

func main() {
	bearerToken := os.Getenv("BEARER_TOKEN")

	client := gotwtr.New(bearerToken)

	tweetText := "Hello, world! This is my first tweet using gotwtr library."
	tweet, err := postTweet(client, tweetText)
	if err != nil {
		log.Fatalf("Error posting tweet: %v", err)
	}

	log.Printf("Successfully posted tweet: %s (ID: %s)", tweet.PostTweetData.Text, tweet.PostTweetData.ID)
}

func postTweet(client *gotwtr.Client, text string) (*gotwtr.PostTweetResponse, error) {
	tweet, err := client.PostTweet(context.Background(), &gotwtr.PostTweetOption{Text: text})
	if err != nil {
		return nil, err
	}

	return tweet, nil
}
