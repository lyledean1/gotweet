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
	"github.com/dghubble/go-twitter/twitter"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search twitter by query",
	Long: `Search twitter by command`,
	Run: func(cmd *cobra.Command, args []string) {
		query, _:= cmd.Flags().GetString("query")
		if query == "" {
			fmt.Println("No query set for search")
			return
		}
		search, err := twitterClient.Search(&twitter.SearchTweetParams{
			Query: query,
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
}

func init() {
	rootCmd.AddCommand(searchCmd)
	searchCmd.Flags().StringP("query", "q", "", "Set query for searching twitter")
}
