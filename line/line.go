package line

import (
 "fmt"
 "github.com/line/line-bot-sdk-go/linebot"
)

var (
  channelAccessToken = "<channel access token>"
  channelSecret = "<channel secret>"
  ID = "<ID>"
)

func Init(token string, secret string, id string) {
  channelAccessToken = token
  channelSecret = secret
  ID = id
}

func Post(text string) {
  bot, err := linebot.New(channelSecret, channelAccessToken)
  if err != nil {
    fmt.Println("error:", err)
  }

  _, err = bot.PushMessage(
    ID,
    linebot.NewTextMessage(text),
  ).Do()
  if err != nil {
    fmt.Println("error:", err)
  }
}

