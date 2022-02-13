package cli

import (
	"bytes"
	api "echoserver/echoserver/api"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func CreateMessage(text string) {
	var message api.Message
	message = api.Message{
		Text: text,
	}

	messageJson, err := json.Marshal(message)
	responsebody := bytes.NewBuffer(messageJson)
	resp, err := http.Post(echoEndpoint, APPJSON, responsebody)

	if err != nil {
		fmt.Println("An error occurred", err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	var result api.Message
	err = json.Unmarshal([]byte(body), &result)
	if err != nil {
		fmt.Println("Error unmarshaling data from request.")
	}

	fmt.Println("Sent message", result.Text)
}
