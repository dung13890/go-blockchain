package main

import (
	"bytes"
	_ "fmt"
	"net/http"
)

type CLI struct{}

func main() {
	sendToSlack("Test Notification!")
}

func sendToSlack(content string) (err error) {
	client := &http.Client{}
	url := "https://hooks.slack.com/services/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"
	var jsonStr = []byte(`{"text":"` + content + `"}`)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Add("Content-Type", "application/json")

	_, responseErr := client.Do(req)

	return responseErr
}
