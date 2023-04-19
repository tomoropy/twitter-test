package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func main() {
	apiKey := os.Getenv("API_KEY")
	apiSecretKey := os.Getenv("API_SECRET_KEY")
	accessToken := os.Getenv("ACCESS_TOKEN")
	accessTokenSecret := os.Getenv("ACCESS_TOKEN_SECRET")

	config := oauth1.NewConfig(apiKey, apiSecretKey)
	token := oauth1.NewToken(accessToken, accessTokenSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	client := twitter.NewClient(httpClient)

	data, err := ioutil.ReadFile("merge.txt")
	if err != nil {
		log.Fatal(err)
	}

	tweet, _, err := client.Statuses.Update(string(data), nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Posted tweet: %s\n", tweet.Text)
}
