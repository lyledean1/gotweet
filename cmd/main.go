package main

import (
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/lyledean1/tweetbot/config"
	"github.com/lyledean1/tweetbot/twit"
	"github.com/urfave/cli"
	"log"
	"os"
)

func TwitterClient(twitterConfig config.TwitterConfig) (*twitter.Client, error) {
	config := oauth1.NewConfig(twitterConfig.Auth.ConsumerKey, twitterConfig.Auth.ConsumerSecret)
	token := oauth1.NewToken(twitterConfig.Auth.AccessToken, twitterConfig.Auth.AccessTokenSecret)

	httpClient := config.Client(oauth1.NoContext, token)

	client := twitter.NewClient(httpClient)

	verifyParams := &twitter.AccountVerifyParams{
		SkipStatus:   twitter.Bool(true),
		IncludeEmail: twitter.Bool(true),
	}
	_, _, err := client.Accounts.VerifyCredentials(verifyParams)
	if err != nil {
		return nil, err
	}

	return client, nil
}

var app = cli.NewApp()

func info() {
	app.Name = "Twitter CLI"
	app.Usage = "Search and Query Tweets"
	app.Author = "Lyle Dean"
	app.Version = "0.1.0"
}


func commands() {
	app.Commands = []cli.Command{
		{
			Name:    "search",
			Aliases: []string{"s"},
			Flags: []cli.Flag{cli.StringFlag{
				Name:  "query, q",
				Usage: "Search over query",
			},
			},
			Usage: "search",
			Action: func(c *cli.Context) {
				client, err := TwitterClient(config.GetTwitterConfig())
				if err != nil {
					panic(err)
				}
				twitterClient := twit.NewTwitterClient(client)

				//twitterClient.UpdateStatus("Saturdays are for chilling", nil)
				search, err := twitterClient.Search(&twitter.SearchTweetParams{
					Query: c.String("q"),
					Count: 100})

				if err != nil {
					panic(err)
				}
				fmt.Println(len(search.Statuses))

				for i, _ := range search.Statuses {
					fmt.Println(search.Statuses[i])
				}

			},
		},
		{
			Name:    "tweet",
			Aliases: []string{"tweet"},
			Usage:   "tweet",
			Flags: []cli.Flag{cli.StringFlag{
				Name:  "text, t",
				Usage: "Text for tweet",
			}},
			Action: func(c *cli.Context) {
				client, err := TwitterClient(config.GetTwitterConfig())
				if err != nil {
					panic(err)
				}
				twitterClient := twit.NewTwitterClient(client)
				tweet := c.String("t")
				//twitterClient.UpdateStatus("Saturdays are for chilling", nil)
				err = twitterClient.UpdateStatus(tweet, nil)
				if err != nil {
					panic("Cannot post tweet")
				}
				fmt.Println(fmt.Sprintf("Tweet %s posted successfully", tweet))
			},
		},
	}
}

func main() {
	info()
	commands()
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
