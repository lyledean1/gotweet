### Go tweet!

Set up enviroment variables from twitter-env-example.yaml and save this as twitter-env.yaml

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