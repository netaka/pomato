package main

import (
  "flag"
  "fmt"
  "github.com/kazuph/go-binenv"
  "github.com/netaka/pomato/slack"
  "github.com/netaka/pomato/line"
)

var (
  message = flag.String("m", "undefined", "message string")
)

func main() {
  env, err := binenv.Load(Asset)
  if err != nil {
    fmt.Println(err)
  }

  slack.Init(env["SLACK_TOKEN"])
  line.Init(env["LINE_CHANNEL_ACCESS_TOKEN"], env["LINE_CHANNEL_SECRET"], env["LINE_USER_ID"])

  flag.Parse()
  fmt.Println(*message);
  slack.Post(*message);
  line.Post(*message);
}

