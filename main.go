package main

import (
  "flag"
  "fmt"
  "./slack"
)

var (
  message = flag.String("m", "undefined", "message string")
)

func main() {
  flag.Parse()
  fmt.Println(*message);
  slack.Post(*message);
}

