/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// tweetCmd represents the tweet command
var tweetCmd = &cobra.Command{
	Use:   "tweet",
	Short: "Tweet to twitter",
	Long: "Tweet to twitter",
	Run: func(cmd *cobra.Command, args []string) {

		text, _:= cmd.Flags().GetString("name")
		if text == "" {
			panic("No text set for tweet")
		}
		tweet, err := twitterClient.UpdateStatus(text, nil)
		if err != nil {
			panic(fmt.Sprintf("Cannot post tweet %s", tweet))
		}
		fmt.Println(fmt.Sprintf("Tweet %s posted successfully with id %d", tweet.Text, tweet.ID))
	},
}

func init() {
	rootCmd.AddCommand(tweetCmd)
	tweetCmd.Flags().StringP("text", "t", "", "Set the text for the tweet")
}
