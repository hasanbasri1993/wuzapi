package Chatwoot

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"wuzapi/model"
)

func prosesConversation(idContact int) int {
	var searchConversationChatwoot model.SearchConversationChatwoot
	var createConversationChatwoot model.ConversationChatwoot
	var idConversation = 0
	_, _, inbox, err := getConfig()
	if err != nil {
		fmt.Println(err)
	}
	resultSearchConversation := searchConversation(idContact)
	errS := json.Unmarshal([]byte(resultSearchConversation), &searchConversationChatwoot)
	if errS != nil {
		fmt.Println(errS)
	}
	dataConversation := model.CreateConversation{
		ContactID: idContact,
		InboxID:   inbox,
		Status:    "open",
	}
	if len(searchConversationChatwoot.Payload) == 0 {
		fmt.Println("Conversation not found")
		resultCreateConversation := createConversation(dataConversation)
		err := json.Unmarshal([]byte(resultCreateConversation), &createConversationChatwoot)
		if err != nil {
			fmt.Println(err)
		}
		idConversation = createConversationChatwoot.ID
	} else {
		intVar, _ := strconv.Atoi(inbox)
		for i := 0; i < len(searchConversationChatwoot.Payload); i++ {
			if searchConversationChatwoot.Payload[i].InboxID == intVar {
				idConversation = searchConversationChatwoot.Payload[i].ID
				break
			}
		}
		if idConversation == 0 {
			fmt.Println("Conversation found but not for this inbox")
			resultCreateConversation := createConversation(dataConversation)
			err := json.Unmarshal([]byte(resultCreateConversation), &createConversationChatwoot)
			if err != nil {
				fmt.Println("", err)
			}
			idConversation = searchConversationChatwoot.Payload[0].ID
		}
	}
	return idConversation
}

func searchConversation(idContact int) string {
	path := "contacts/" + strconv.Itoa(idContact) + "/conversations"
	method := "GET"
	payload := strings.NewReader(``)
	return requestChatwoot(method, path, payload)
}

func createConversation(conversation model.CreateConversation) string {
	jsonData, err := json.Marshal(conversation)
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}
	path := "conversations"
	method := "POST"
	payload := strings.NewReader(string(jsonData))
	return requestChatwoot(method, path, payload)
}
