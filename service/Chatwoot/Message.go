package Chatwoot

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"wuzapi/model"
)

/*
Process data before request to Chatwoot API
*/
func praSendMessage(idConversation int, oneSenderWebhook model.OneSenderWebhook) model.SendMessageResult {
	// Will add watermark to message if it's from Chatwoot Admin, prevent duplicate message info Chatwoot Conversation
	if strings.Contains(oneSenderWebhook.MessageText, "🤖 _*Chatwoot*_") {
		return model.SendMessageResult{}
	}

	var sendMessageResult model.SendMessageResult
	dataSendMessage := model.SendMessage{
		Content:     oneSenderWebhook.MessageText,
		MessageType: "incoming",
		Private:     false,
	}

	if oneSenderWebhook.IsFromMe {
		dataSendMessage.MessageType = "outgoing"
		dataSendMessage.Private = true
	}

	if oneSenderWebhook.IsGroup && !oneSenderWebhook.IsFromMe {
		dataSendMessage.Content = oneSenderWebhook.MessageText + "\n\n\n🤖 **" + oneSenderWebhook.SenderPushName + "**"
	}

	if oneSenderWebhook.MessageType != "text" && oneSenderWebhook.MessageType != "unknown" {
		filebase := "./files/user_1/" + oneSenderWebhook.MessageID
		fileName := ""
		switch oneSenderWebhook.MessageType {
		case "video":
			fileName = filebase + ".f4v"
			break
		case "document":
			fileName = "./files/user_1/" + oneSenderWebhook.AttachmentID
			break
		case "audio":
			fileName = filebase + ".oga"
			break
		case "image":
			fileName = filebase + ".jpeg"
			break
		}
		jsonD, _ := json.Marshal(dataSendMessage)
		fmt.Println("requestChatwootAttachment", string(jsonD))
		resultSendMessage := requestChatwootAttachment(idConversation, dataSendMessage, fileName)
		err3 := json.Unmarshal([]byte(resultSendMessage), &sendMessageResult)
		if err3 != nil {
			fmt.Println(err3)
			return model.SendMessageResult{}
		}
	} else {
		jsonD, _ := json.Marshal(dataSendMessage)
		fmt.Println("requestChatwootAttachment", string(jsonD))
		resultSendMessage := sendMessage(idConversation, dataSendMessage)
		err3 := json.Unmarshal([]byte(resultSendMessage), &sendMessageResult)
		if err3 != nil {
			fmt.Println(err3)
			return model.SendMessageResult{}
		}
	}
	return sendMessageResult
}

/*
Send message into Chatwoot Conversation
*/
func sendMessage(idConversation int, contact model.SendMessage) string {
	jsonData, err := json.Marshal(contact)
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}
	path := "conversations/" + strconv.Itoa(idConversation) + "/messages"
	method := "POST"
	payload := strings.NewReader(string(jsonData))
	return requestChatwoot(method, path, payload)
}
