package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/vincent-petithory/dataurl"
	"go.mau.fi/whatsmeow"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/types"
	"google.golang.org/protobuf/proto"
	"net/http"
)

func SendImageProses(chatId string, messageId string, body string, media string, userid int, stanzaId *string, participant *string) (int, string, types.MessageID) {
	if messageId == "" {
		messageId = whatsmeow.GenerateMessageID()
	}
	msgid := messageId
	var resp whatsmeow.SendResponse

	recipient, err := validateMessageFields(chatId, stanzaId, participant)
	if err != nil {
		log.Error().Msg(fmt.Sprintf("%s", err))
		response := map[string]interface{}{"Details": fmt.Sprintf("%s", err)}
		responseJson, _ := json.Marshal(response)
		return http.StatusBadRequest, string(responseJson), msgid
	}

	msgid = whatsmeow.GenerateMessageID()

	var uploaded whatsmeow.UploadResponse
	var filedata []byte

	if media[0:10] == "data:image" {
		dataURL, err := dataurl.DecodeString(media)
		if err != nil {
			response := map[string]interface{}{"Details": fmt.Sprintf("%s", err), "Message": "Could not decode base64 encoded data from payload"}
			responseJson, _ := json.Marshal(response)
			return http.StatusBadRequest, string(responseJson), msgid
		} else {
			filedata = dataURL.Data
			uploaded, err = clientPointer[userid].Upload(context.Background(), filedata, whatsmeow.MediaImage)
			if err != nil {
				response := map[string]interface{}{"Details": fmt.Sprintf("Failed to upload file: %v", err)}
				responseJson, _ := json.Marshal(response)
				return http.StatusInternalServerError, string(responseJson), msgid
			}
		}
	} else {
		response := map[string]interface{}{"Details": "image data should start with \"data:image/png;base64,\""}
		responseJson, _ := json.Marshal(response)
		return http.StatusBadRequest, string(responseJson), msgid
	}

	msg := &waProto.Message{ImageMessage: &waProto.ImageMessage{
		Caption:       proto.String(body),
		Url:           proto.String(uploaded.URL),
		DirectPath:    proto.String(uploaded.DirectPath),
		MediaKey:      uploaded.MediaKey,
		Mimetype:      proto.String(http.DetectContentType(filedata)),
		FileEncSha256: uploaded.FileEncSHA256,
		FileSha256:    uploaded.FileSHA256,
		FileLength:    proto.Uint64(uint64(len(filedata))),
	}}

	if stanzaId != nil {
		msg.ExtendedTextMessage.ContextInfo = &waProto.ContextInfo{
			StanzaId:      proto.String(*stanzaId),
			Participant:   proto.String(*participant),
			QuotedMessage: &waProto.Message{Conversation: proto.String("")},
		}
	}

	resp, err = clientPointer[userid].SendMessage(context.Background(), recipient, msg, whatsmeow.SendRequestExtra{ID: msgid})
	if err != nil {
		response := map[string]interface{}{"Details": fmt.Sprintf("Error sending message: %v", err)}
		responseJson, _ := json.Marshal(response)
		return http.StatusInternalServerError, string(responseJson), msgid
	}

	log.Info().Str("timestamp", fmt.Sprintf("%d", resp.Timestamp)).Str("id", msgid).Msg("Message sent")
	response := map[string]interface{}{"Details": "Sent", "Timestamp": resp.Timestamp, "Id": msgid}
	responseJson, err := json.Marshal(response)
	return http.StatusOK, string(responseJson), msgid

}

func SendMessageProcess(chatId string, messageId string, body string, userid int, stanzaId *string, participant *string) (int, string) {
	if messageId == "" {
		messageId = whatsmeow.GenerateMessageID()
	}
	msgid := messageId
	var resp whatsmeow.SendResponse

	msgid = whatsmeow.GenerateMessageID()
	msg := &waProto.Message{
		ExtendedTextMessage: &waProto.ExtendedTextMessage{
			Text: &body,
		},
	}

	recipient, err := validateMessageFields(chatId, stanzaId, participant)
	if err != nil {
		log.Error().Msg(fmt.Sprintf("%s", err))
		response := map[string]interface{}{"Details": fmt.Sprintf("%s", err)}
		responseJson, _ := json.Marshal(response)
		return http.StatusBadRequest, string(responseJson)
	}

	if stanzaId != nil {
		msg.ExtendedTextMessage.ContextInfo = &waProto.ContextInfo{
			StanzaId:      proto.String(*stanzaId),
			Participant:   proto.String(*participant),
			QuotedMessage: &waProto.Message{Conversation: proto.String("")},
		}
	}
	resp, err = clientPointer[userid].SendMessage(context.Background(), recipient, msg, whatsmeow.SendRequestExtra{ID: msgid})
	if err != nil {
		response := map[string]interface{}{"Details": fmt.Sprintf("Error sending message: %v", err)}
		responseJson, _ := json.Marshal(response)
		return http.StatusInternalServerError, string(responseJson)
	}
	log.Info().Str("timestamp", fmt.Sprintf("%d", resp.Timestamp)).Str("id", msgid).Msg("Message sent")
	response := map[string]interface{}{"Details": "Sent", "Timestamp": resp.Timestamp, "Id": msgid}
	responseJson, _ := json.Marshal(response)
	return http.StatusOK, string(responseJson)
}
