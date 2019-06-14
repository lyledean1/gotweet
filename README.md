### Go tweet!

This uses the developer API in twitter so you will have to set up an account with the correct permissions.

Set up enviroment variables from twitter-env-example.yaml and save this as twitter-env.yaml in your home directory.

I.e ```$HOME/twitter-env.yaml```

```
auth:
  consumerKey: KEY
  consumerSecret: SECRET
  accessToken: TOKEN
  accessTokenSecret: TOKENSECRET
```

To build the binary run 
```go build -mod vendor -o ./bin/gotweet ./cmd/main.go```

### Commands

1) Search Over Tweets

```gotweet search -query "golang"```

2) Tweet to your account (the account linked in the twitter yaml file)

```gotweet tweet -text "Hello World"```