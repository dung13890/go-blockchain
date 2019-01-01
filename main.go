package main

import (
	"bufio"
	"bytes"
	"fmt"
	"net/http"
	"os"
)

type CLI struct{}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("What do you want messages to Slack? \n")
	text, _ := reader.ReadString('\n')

	sendToSlack(text)
}

func sendToSlack(content string) (err error) {
	fmt.Println(content)
	client := &http.Client{}
	url := "https://hooks.slack.com/services/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"
	var jsonStr = []byte(`{"text":"` + content + `"}`)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Add("Content-Type", "application/json")

	_, responseErr := client.Do(req)

	return responseErr
}
