package Chatwoot

import (
	"bytes"
	"fmt"
	"github.com/gabriel-vasile/mimetype"
	"github.com/spf13/viper"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
	"path/filepath"
	"strconv"
	"wuzapi/model"
)

func requestChatwoot(method string, path string, payload io.Reader) string {
	url, account, _, err := getConfig()
	if err != nil {
		fmt.Println(err)
	}
	url = url + "/api/v1/accounts/" + account + "/" + path
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("api_access_token", viper.GetString("chatwoot.accountToken"))

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		_ = err
		fmt.Println(err)
	}
	return string(body)
}

func requestChatwootAttachment(idConversation int, contact model.SendMessage, fileName string) string {
	url, account, _, err := getConfig()
	if err != nil {
		fmt.Println(err)
	}
	url = url + "/api/v1/accounts/" + account + "/"
	path := url + "conversations/" + strconv.Itoa(idConversation) + "/messages"
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	file, _ := os.Open(fileName)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	mtype, err := mimetype.DetectFile(fileName)
	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition",
		fmt.Sprintf(`form-data; name="%s"; filename="%s"`,
			"attachments[]", filepath.Base(fileName)))
	h.Set("Content-Type", mtype.String())
	part3, _ := writer.CreatePart(h)

	_, _ = io.Copy(part3, file)
	_ = writer.WriteField("message_type", "incoming")
	_ = writer.WriteField("content", contact.Content)

	err = writer.Close()
	if err != nil {
		fmt.Println(err)
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, path, payload)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("api_access_token", viper.GetString("chatwoot.accountToken"))
	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(res.Body)
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	return string(body)
}
