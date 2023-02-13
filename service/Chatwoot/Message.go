package Chatwoot

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"wuzapi/model"
)

func praSendMessage(idConversation int, oneSenderWebhook model.OneSenderWebhook) model.SendMessageResult {
	var sendMessageResult model.SendMessageResult
	dataSendMessage := model.SendMessage{
		Content:     oneSenderWebhook.MessageText,
		MessageType: "incoming",
		Private:     false,
	}

	if oneSenderWebhook.IsFromMe {
		dataSendMessage.MessageType = "outgoing"
	}

	if oneSenderWebhook.IsGroup {
		dataSendMessage.Content =
			oneSenderWebhook.MessageText + "\n\n\n🤖 **" + oneSenderWebhook.SenderPushName + "**"
	}

	if oneSenderWebhook.MessageType != "text" && oneSenderWebhook.MessageType != "unknown" {
		filebase := "./files/user_1/" + oneSenderWebhook.MessageID
		fileName := ""
		switch oneSenderWebhook.MessageType {
		case "video":
			fileName = filebase + ".f4v"
		case "document":
			fileName = "./files/user_1/" + oneSenderWebhook.AttachmentID
		case "audio":
			fileName = filebase + ".oga"
		default:
			fileName = filebase + ".jpeg"
		}
		resultSendMessage := requestChatwootAttachment(idConversation, dataSendMessage, fileName)
		err3 := json.Unmarshal([]byte(resultSendMessage), &sendMessageResult)
		if err3 != nil {
			fmt.Println(err3)
			return model.SendMessageResult{}
		}
	} else {
		resultSendMessage := sendMessage(idConversation, dataSendMessage)
		err3 := json.Unmarshal([]byte(resultSendMessage), &sendMessageResult)
		if err3 != nil {
			fmt.Println(err3)
			return model.SendMessageResult{}
		}
	}
	return sendMessageResult
}

func renameF(org string, new string) {
	e := os.Rename(org, new)
	if e != nil {
		log.Fatal(e)
	}
}
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

func downloadFile(URL, fileName string) error {
	response, err := http.Get(URL)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(response.Body)

	if response.StatusCode != 200 {
		return errors.New("received non 200 response code")
	}
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}
	return nil
}
