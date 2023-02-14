package Chatwoot

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"wuzapi/model"
)

// OutcomeChawoot Chatwoot is the main function for the Chatwoot service
func OutcomeChawoot(t model.ChatwootWebhook) string {
	if t.Event != "message_created" {
		return "skip t.Even" + t.Event
	}
	if t.Private == true {
		return "skip t.Private " + strconv.FormatBool(t.Private) // => "true"
	}
	if t.MessageType == "incoming" {
		return "skip t.MessageType " + t.MessageType
	}

	to := t.Conversation.Meta.Sender.PhoneNumber[1:]
	if t.Conversation.Meta.Sender.CustomAttributes.IsGroup {
		to = t.Conversation.Meta.Sender.CustomAttributes.GroupID
	}
	url := "http://localhost:9001/chat/send/text"
	method := "POST"
	payload := strings.NewReader(`{"Phone": "` + to + `","Body": "` + t.Content + `\n\n🤖 _*Chatwoot*_ admin: _*` + t.Sender.Name + `*_"}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("token", "hjve6uly")
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	_, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Outcoming Chatwoot", res.StatusCode)
	return "chatwoot service is running"
}
