package main

import (
  "flag"
  "fmt"
  "github.com/netaka/pomato/slack"
  "github.com/netaka/pomato/line"
)

var (
  message = flag.String("m", "undefined", "message string")
)

func main() {
  flag.Parse()
  fmt.Println(*message);
  slack.Post(*message);
  line.Post(*message);
}

