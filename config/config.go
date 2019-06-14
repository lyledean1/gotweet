package config

import (
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"gopkg.in/yaml.v2"
	"os/user"
	"io/ioutil"
	"log"
)

type TwitterConfig struct {
	Auth Auth `yaml:"auth"`
}

type Auth struct {
	ConsumerKey     string `yaml:"consumerKey"`
	ConsumerSecret string `yaml:"consumerSecret"`
	AccessToken string `yaml:"accessToken"`
	AccessTokenSecret string `yaml:"accessTokenSecret"`
}

func GetTwitterConfig() TwitterConfig {
	var config TwitterConfig
	usr, err := user.Current()
	if err != nil {
		log.Fatal( err )
	}
	file, err := ioutil.ReadFile(fmt.Sprintf("%s/twitter-env.yaml", usr.HomeDir))
	if err != nil {
		panic(err)
	}
	yaml.Unmarshal(file, &config)
	return config
}

func NewTwitterClient(twitterConfig TwitterConfig) (*twitter.Client, error) {
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