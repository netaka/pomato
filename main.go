package main

import (
	"flag"
	"fmt"
    "encoding/json"
    _ "embed"
	"github.com/netaka/pomato/line"
	"github.com/netaka/pomato/slack"
)

var (
	message     = flag.String("m", "undefined", "message string")
	enableSlack = flag.Bool("s", false, "enable posting to Slack")
	enableLine  = flag.Bool("l", false, "enable posting to LINE")
)

//go:embed token.json
var tokenData []byte

type token struct {
    slack_token string `json:"slack_token"`
    slack_channel string `json:"slack_channel"`
    line_channel_access_token string `json:"line_channel_access_token"`
    line_channel_secret string `json:"line_channel_secret"`
    line_user_id string `json:"line_user_id"`
}

func main() {
    var t token
    json.Unmarshal([]byte(tokenData), &t)
	slack.Init(t.slack_token, t.slack_channel)
	line.Init(t.line_channel_access_token, t.line_channel_secret, t.line_user_id)

	flag.Parse()
	fmt.Println(*message)
	if *enableSlack {
		slack.Post(*message)
	}
	if *enableLine {
		line.Post(*message)
	}
}
