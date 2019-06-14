package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
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
	file, err := ioutil.ReadFile("./twitter-env.yaml")
	if err != nil {
		panic(err)
	}
	yaml.Unmarshal(file, &config)
	return config
}