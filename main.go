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
    SlackToken string `json:"slack_token"`
    SlackChannel string `json:"slack_channel"`
    LineChannelAccessToken string `json:"line_channel_access_token"`
    LineChannelSecret string `json:"line_channel_secret"`
    LineUserId string `json:"line_user_id"`
}

func main() {
    var t token
    err := json.Unmarshal(tokenData, &t)
    if err != nil {
        fmt.Printf("error: %#v\n", err)
    }

    slack.Init(t.SlackToken, t.SlackChannel)
    line.Init(t.LineChannelAccessToken, t.LineChannelSecret, t.LineUserId)

    flag.Parse()
    fmt.Println(*message)
    if *enableSlack {
        slack.Post(*message)
    }
    if *enableLine {
        line.Post(*message)
    }
}
