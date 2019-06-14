# Go tweet!

## Installation 
This uses the developer API in twitter so you will have to set up an account with the correct permissions.

https://developer.twitter.com/

Set up enviroment variables from twitter-env-example.yaml and save this as twitter-env.yaml in your home directory.

I.e ```$HOME/.gotweet.yaml```

```
auth:
  consumerKey: KEY
  consumerSecret: SECRET
  accessToken: TOKEN
  accessTokenSecret: TOKENSECRET
```

To get the latest version, which should add this to the go/bin folder

```go get github.com/lyledean1/gotweet```

And finally make sure that the path is linked to the go bin path.

```export PATH=$PATH:$GOPATH/bin ```

### Build the binary from source

To build the binary run 
```go build -mod vendor -o ./bin/gotweet ./cmd/main.go```

To add the binary to the Go bin path, 

```
 cp ./bin/gotweet ~/go/bin/   
```

### Commands 

1) Search Over Tweets

```gotweet search -query "golang"```

2) Tweet to your account (the account linked in the twitter yaml file)

```gotweet tweet -text "Hello World"```

