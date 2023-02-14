package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mau.fi/whatsmeow"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/types"
	"google.golang.org/protobuf/proto"
	"net/http"
)

func SendMessageProcess(chatId string, body string, userid int, stanzaId *string, participant *string) (int, string, types.MessageID) {

	msgid := ""
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
		return http.StatusBadRequest, string(responseJson), msgid
	}

	if stanzaId != nil {
		msg.ExtendedTextMessage.ContextInfo = &waProto.ContextInfo{
			StanzaId:      proto.String(*stanzaId),
			Participant:   proto.String(*participant),
			QuotedMessage: &waProto.Message{Conversation: proto.String("")},
		}
	}
	resp, err = clientPointer[userid].SendMessage(context.Background(), recipient, msg)
	if err != nil {
		response := map[string]interface{}{"Details": fmt.Sprintf("Error sending message: %v", err)}
		responseJson, _ := json.Marshal(response)
		return http.StatusInternalServerError, string(responseJson), msgid
	}
	log.Info().Str("timestamp", fmt.Sprintf("%d", resp.Timestamp)).Str("id", msgid).Msg("Message sent")
	response := map[string]interface{}{"Details": "Sent", "Timestamp": resp.Timestamp, "Id": msgid}
	responseJson, _ := json.Marshal(response)
	return http.StatusOK, string(responseJson), msgid
}
