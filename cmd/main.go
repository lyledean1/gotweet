package main

import (
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/fatih/color"
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
					if i % 2 == 0 {
						color.Yellow("%s, User: %s", search.Statuses[i].Text, search.Statuses[i].User.Name)
					} else {
						color.Green("%s, User: %s", search.Statuses[i].Text, search.Statuses[i].User.Name)
					}
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
				text := c.String("t")
				tweet, err := twitterClient.UpdateStatus(text, nil)
				if err != nil {
					panic(fmt.Sprintf("Cannot post tweet %s", tweet))
				}
				fmt.Println(fmt.Sprintf("Tweet %s posted successfully with id %d", tweet.Text, tweet.ID))
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
