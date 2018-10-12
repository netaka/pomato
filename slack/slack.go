package slack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Arguments struct {
	Channel string `json:"channel"`
	Text    string `json:"text"`
	AsUser  bool   `json:"as_user"`
}

var (
	methodURL   = "https://slack.com/api/chat.postMessage"
	apiToken    = "xxxx-xxxxxxxxx-xxxx"
	postChannel = "#general"
)

func Init(token string, channel string) {
	apiToken = token
	postChannel = channel
}

func Post(text string) {
	arguments := Arguments{
		Channel: postChannel,
		Text:    text,
		AsUser:  true,
	}

	b, err := json.Marshal(arguments)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(b))

	buf := bytes.NewReader(b)
	client := &http.Client{}
	req, err := http.NewRequest("POST", methodURL, buf)
	if err != nil {
		fmt.Println("error:", err)
	}
	req.Header.Add("Content-type", `application/json`)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", apiToken))
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error:", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
