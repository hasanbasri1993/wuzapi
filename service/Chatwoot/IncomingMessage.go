package Chatwoot

import (
	"encoding/json"
	"fmt"
	"strings"
	"wuzapi/model"
)

func IncomingMessage(c string) model.SendMessageResult {
	var result map[string]any
	var err = json.Unmarshal([]byte(c), &result)
	if err != nil {
		fmt.Println(err.Error())
		return model.SendMessageResult{}
	}
	preProsesHook := preProses(c)
	rrContact := prosesContact(preProsesHook)
	rrConversation := prosesConversation(rrContact)
	rrSend := praSendMessage(rrConversation, preProsesHook)
	return rrSend
}

func preProses(c string) model.OneSenderWebhook {
	var oneSenderWebhook model.OneSenderWebhook
	var result map[string]any
	var err = json.Unmarshal([]byte(c), &result)
	if err != nil {
		fmt.Println(err.Error())
	}
	message := result["event"].(map[string]any)["Message"]
	info := result["event"].(map[string]any)["Info"].(map[string]any)
	oneSenderWebhook.Chat = info["Chat"].(string)
	oneSenderWebhook.Sender = info["Sender"].(string)
	oneSenderWebhook.SenderPushName = info["PushName"].(string)
	oneSenderWebhook.SenderPhone = strings.Split(info["Chat"].(string), "@")[0]
	oneSenderWebhook.IsGroup = info["IsGroup"].(bool)
	oneSenderWebhook.IsFromMe = info["IsFromMe"].(bool)
	oneSenderWebhook.MessageID = info["ID"].(string)
	oneSenderWebhook.MessageTimestamp = info["Timestamp"].(string)
	if message.(map[string]any)["conversation"] != nil {
		conversation := message.(map[string]any)["conversation"]
		oneSenderWebhook.MessageText = conversation.(string)
		oneSenderWebhook.MessageType = "text"
	}

	if message.(map[string]any)["extendedTextMessage"] != nil {
		extendedTextMessage := message.(map[string]any)["extendedTextMessage"].(map[string]any)["text"]
		oneSenderWebhook.MessageText = extendedTextMessage.(string)
		oneSenderWebhook.MessageType = "text"
	}

	if message.(map[string]any)["imageMessage"] != nil {
		imageMessage := message.(map[string]any)["imageMessage"]
		oneSenderWebhook.AttachmentID = info["ID"].(string)
		oneSenderWebhook.AttachmentType = "image"
		oneSenderWebhook.MessageText = imageMessage.(map[string]any)["caption"].(string)
		oneSenderWebhook.MessageType = "image"
	}

	if message.(map[string]any)["videoMessage"] != nil {
		videoMessage := message.(map[string]any)["videoMessage"]
		oneSenderWebhook.AttachmentID = info["ID"].(string)
		oneSenderWebhook.AttachmentType = "video"
		oneSenderWebhook.MessageText = videoMessage.(map[string]any)["caption"].(string)
		oneSenderWebhook.MessageType = "video"
	}

	if message.(map[string]any)["audioMessage"] != nil {
		oneSenderWebhook.AttachmentID = info["ID"].(string)
		oneSenderWebhook.AttachmentType = "audio"
		oneSenderWebhook.MessageType = "audio"
	}

	if message.(map[string]any)["documentMessage"] != nil {
		documentMessage := message.(map[string]any)["documentMessage"]
		oneSenderWebhook.AttachmentID = documentMessage.(map[string]any)["fileName"].(string)
		oneSenderWebhook.MessageText = documentMessage.(map[string]any)["caption"].(string)
		oneSenderWebhook.AttachmentType = "document"
		oneSenderWebhook.MessageType = "document"
	}
	return oneSenderWebhook
}
