package main

import (
	"encoding/base64"
	"io"
	"io/ioutil"
	"net/http"
	"os"
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
	content := t.Content + "\n\n🤖 _*Chatwoot*_ admin: _*" + t.Sender.Name + "*_"
	if t.Conversation.Meta.Sender.CustomAttributes.IsGroup {
		to = t.Conversation.Meta.Sender.CustomAttributes.GroupID
	}

	if len(t.Conversation.Messages[0].Attachments) > 0 {
		attach := t.Conversation.Messages[0].Attachments
		for i, s := range attach {
			switch s.FileType {
			case "image":
				contentAttach := content + "\n📎 _*Attachment " + strconv.Itoa(i+1) + "*_"
				if i > 0 {
					contentAttach = "\n\n🤖 _*Chatwoot*_ admin: _*" + t.Sender.Name + "*_" + "\n📎 _*Attachment " + strconv.Itoa(i+1) + "*_"
				}

				strSS := strings.Split(s.DataURL, "/")
				fileName := strSS[len(strSS)-1]
				err := DownloadFile(fileName, s.DataURL)
				if err != nil {
					return ""
				}
				base64File := fileBase64(fileName)
				SendImageProses(to, contentAttach, base64File, 1, nil, nil)
				break
			}
		}
	} else {
		SendMessageProcess(to, content, 1, nil, nil)
	}

	return "chatwoot service is running"
}

func fileBase64(filepath string) string {
	bytes, err := ioutil.ReadFile("files/user_1/" + filepath)
	if err != nil {
		log.Fatal()
	}
	var base64Encoding string
	mimeType := http.DetectContentType(bytes)
	switch mimeType {
	case "image/jpeg":
		base64Encoding += "data:image/jpeg;base64,"
	case "image/png":
		base64Encoding += "data:image/png;base64,"
	}
	base64Encoding += toBase64(bytes)
	return base64Encoding
}

func toBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create("files/user_1/" + filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
