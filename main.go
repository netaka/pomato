package main

import (
	"flag"
	"fmt"
	"github.com/kazuph/go-binenv"
	"github.com/netaka/pomato/line"
	"github.com/netaka/pomato/slack"
)

var (
	message     = flag.String("m", "undefined", "message string")
	enableSlack = flag.Bool("s", false, "enable posting to Slack")
	enableLine  = flag.Bool("l", false, "enable posting to LINE")
)

func main() {
	env, err := binenv.Load(Asset)
	if err != nil {
		fmt.Println(err)
	}

	slack.Init(env["SLACK_TOKEN"], env["SLACK_CHANNEL"])
	line.Init(env["LINE_CHANNEL_ACCESS_TOKEN"], env["LINE_CHANNEL_SECRET"], env["LINE_USER_ID"])

	flag.Parse()
	fmt.Println(*message)
	if *enableSlack {
		slack.Post(*message)
	}
	if *enableLine {
		line.Post(*message)
	}
}
