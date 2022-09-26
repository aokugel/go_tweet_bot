package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

type Credentials struct {
	ConsumerKey       string
	ConsumerSecret    string
	AccessToken       string
	AccessTokenSecret string
}

func getClient(creds *Credentials) (*twitter.Client, error) {
	// Pass in your consumer key (API Key) and your Consumer Secret (API Secret)
	config := oauth1.NewConfig(creds.ConsumerKey, creds.ConsumerSecret)
	// Pass in your Access Token and your Access Token Secret
	token := oauth1.NewToken(creds.AccessToken, creds.AccessTokenSecret)

	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	// Verify Credentials
	verifyParams := &twitter.AccountVerifyParams{
		SkipStatus:   twitter.Bool(true),
		IncludeEmail: twitter.Bool(true),
	}

	// we can retrieve the user and verify if the credentials
	// we have used successfully allow us to log in!
	user, _, err := client.Accounts.VerifyCredentials(verifyParams)
	if err != nil {
		return nil, err
	}

	log.Printf("User's ACCOUNT:\n%+v\n", user)
	return client, nil
}

func main() {
	fmt.Println("I am go twitter bit v1.0 beep boop.")
	creds := Credentials{
		AccessToken:       os.Getenv("ACCESS_TOKEN"),
		AccessTokenSecret: os.Getenv("ACCESS_TOKEN_SECRET"),
		ConsumerKey:       os.Getenv("CONSUMER_KEY"),
		ConsumerSecret:    os.Getenv("CONSUMER_SECRET"),
	}

	client, err := getClient(&creds)
	if err != nil {
		log.Println("Error getting Twitter Client")
		log.Println(err)
	}
	// //Below retweets larrys five most recent tweets
	// larryTimeline := &twitter.UserTimelineParams{
	// 	Count:           5,
	// 	TweetMode:       "extended",
	// 	UserID:          1168018744804691968,
	// 	ExcludeReplies:  twitter.Bool(false),
	// 	IncludeRetweets: twitter.Bool(false),
	// }
	// tweets, _, _ := client.Timelines.UserTimeline(larryTimeline)

	// for index, tweet := range tweets {
	// 	fmt.Printf("%v \n %v \n %T \n", index, tweet.FullText, tweet.FullText)
	// 	client.Statuses.Retweet(tweet.ID, nil)
	// }

	// fmt.Println(tweets[0].ID)

	// client.Statuses.Retweet(tweets[0].ID, nil)

	//This shows 15 tweets suing the givin query
	search, _, err := client.Search.Tweets(&twitter.SearchTweetParams{
		Query:     "Iowa State",
		TweetMode: "extended",
	})

	for index, tweet := range search.Statuses {
		fmt.Printf("Tweet: %v \n User: %v \n Tweet: %v \n", index, tweet.User.Name, tweet.FullText)
	}

}
