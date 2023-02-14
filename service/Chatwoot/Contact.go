package Chatwoot

import (
	"encoding/json"
	"fmt"
	"strings"
	"wuzapi/model"
)

func prosesContact(oneSenderWebhook model.OneSenderWebhook) int {
	var searchContactChatwoot model.SearchContactChatwoot
	var createContactChatwoot model.ContactChatwoot
	var idContact = 0
	_, _, inbox, err := getConfig()
	if err != nil {
		fmt.Println(err)
	}
	resultContact := searchContact(oneSenderWebhook.Chat)
	err = json.Unmarshal([]byte(resultContact), &searchContactChatwoot)
	if err != nil {
	}
	if searchContactChatwoot.Meta.Count == 0 {
		fmt.Println("Contact not found")
		dataContact := model.CreateContact{
			InboxID:     inbox,
			Name:        oneSenderWebhook.SenderPushName,
			Email:       oneSenderWebhook.Chat[3:] + "@gmail.com",
			PhoneNumber: oneSenderWebhook.SenderPhone,
		}
		if oneSenderWebhook.IsGroup {
			dataContact.PhoneNumber = oneSenderWebhook.Chat[3:]
			dataContact.Name = oneSenderWebhook.Chat
			dataContact.ContactAttributes = model.CustomAttributes{
				GroupID: oneSenderWebhook.Chat + "@g.us",
				IsGroup: true,
			}
		}
		dataContact.PhoneNumber = "+" + dataContact.PhoneNumber
		resultCreateContact := createContact(dataContact)
		_ = json.Unmarshal([]byte(resultCreateContact), &createContactChatwoot)
		idContact = createContactChatwoot.Payload.Contact.ID
	} else {
		idContact = searchContactChatwoot.Payload[0].ID
	}
	return idContact
}

func searchContact(number string) string {
	path := "contacts/search?q=" + number
	method := "GET"
	payload := strings.NewReader(``)
	return requestChatwoot(method, path, payload)
}

func createContact(contact model.CreateContact) string {
	jsonData, err := json.Marshal(contact)
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}
	path := "contacts"
	method := "POST"
	payload := strings.NewReader(string(jsonData))
	return requestChatwoot(method, path, payload)
}
