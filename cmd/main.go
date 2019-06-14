package main

import (
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/lyledean1/tweetbot/config"
	"github.com/lyledean1/tweetbot/twit"
	"github.com/urfave/cli"
	"log"
	"os"
)


var twitterClient *twit.TwitterClient

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
				search, err := twitterClient.Search(&twitter.SearchTweetParams{
					Query: c.String("q"),
					Count: 100})
				if err != nil {
					panic(err)
				}
				for i, _ := range search.Statuses {
					fmt.Println(fmt.Sprintf("Tweet `%s` from user `%s`",search.Statuses[i].Text, search.Statuses[i].User))
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
				tweet := c.String("t")
				//twitterClient.UpdateStatus("Saturdays are for chilling", nil)
				err := twitterClient.UpdateStatus(tweet, nil)
				if err != nil {
					panic("Cannot post tweet")
				}
				fmt.Println(fmt.Sprintf("Tweet %s posted successfully", tweet))
			},
		},
	}
}

func main() {
	client, err := config.NewTwitterClient(config.GetTwitterConfig())
	if err != nil {
		panic(err)
	}
	twitterClient = twit.NewTwitterClient(client)
	info()
	commands()
	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
