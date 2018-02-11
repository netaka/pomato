package main

import (
  "flag"
  "fmt"
)

var (
  message = flag.String("m", "undefined", "message string")
)

func main() {
  flag.Parse()
  fmt.Println(*message);
}

