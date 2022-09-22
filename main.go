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
	fmt.Println("Go-Twitter Bot v0.01")
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

	// Print out the pointer to our client
	// for now so it doesn't throw errors
	// tweet, resp, err := client.Statuses.Update(tweetBody, nil)
	// if err != nil {
	// 	log.Println(err)
	// }
	// log.Printf("%v\n", resp)
	// log.Printf("%+v\n", tweet)

	// Home Timeline
	homeTimelineParams := &twitter.HomeTimelineParams{
		Count:     2,
		TweetMode: "extended",
	}
	tweets, _, _ := client.Timelines.HomeTimeline(homeTimelineParams)
	fmt.Printf("User's HOME TIMELINE:\n%+v\n", tweets)

	// Mention Timeline
	mentionTimelineParams := &twitter.MentionTimelineParams{
		Count:     2,
		TweetMode: "extended",
	}
	tweets, _, _ = client.Timelines.MentionTimeline(mentionTimelineParams)
	fmt.Printf("User's MENTION TIMELINE:\n%+v\n", tweets)

	// Retweets of Me Timeline
	retweetTimelineParams := &twitter.RetweetsOfMeTimelineParams{
		Count:     2,
		TweetMode: "extended",
	}
	tweets, _, _ = client.Timelines.RetweetsOfMeTimeline(retweetTimelineParams)
	fmt.Printf("User's 'RETWEETS OF ME' TIMELINE:\n%+v\n", tweets)

	// Update (POST!) Tweet (uncomment to run)
	// tweet, _, _ := client.Statuses.Update("just
	larryTimeline := &twitter.UserTimelineParams{
		Count:     10,
		TweetMode: "extended",
		UserID:    1168018744804691968,
	}
	tweets, _, _ = client.Timelines.UserTimeline(larryTimeline)
	fmt.Printf("User's 'RETWEETS OF ME' TIMELINE:\n%+v\n", tweets)
	fmt.Printf("%T \n", tweets)

	for _, tweet := range tweets {
		// fmt.Println(index, tweet)
		fmt.Printf("%v", tweet.Text)
	}

}
