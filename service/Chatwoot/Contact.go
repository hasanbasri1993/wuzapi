package Chatwoot

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"io"
	"net/http"
	"strings"
	"wuzapi/model"
)

func prosesContact(oneSenderWebhook model.OneSenderWebhook) int {
	var searchContactChatwoot model.SearchContactChatwoot
	var createContactChatwoot model.ContactChatwoot
	var idContact = 0
	var chatIdParse = strings.Split(oneSenderWebhook.Chat, "@")[0]
	if oneSenderWebhook.IsGroup {
		chatIdParse = chatIdParse[3:]
	}
	_, _, inbox, err := getConfig()
	if err != nil {
		fmt.Println(err)
	}
	resultContact := searchContact(chatIdParse)
	err = json.Unmarshal([]byte(resultContact), &searchContactChatwoot)
	if err != nil {
	}
	if searchContactChatwoot.Meta.Count == 0 {
		fmt.Println("Contact not found")
		dataContact := model.CreateContact{
			InboxID:     inbox,
			Name:        oneSenderWebhook.SenderPushName,
			Email:       chatIdParse[3:] + "@gmail.com",
			PhoneNumber: oneSenderWebhook.SenderPhone,
		}
		if oneSenderWebhook.IsGroup {
			var groupInfoSt GroupInfoResponse
			info := infoGroup(oneSenderWebhook.Chat)
			if err != nil {
				fmt.Println(err)
			}
			err = json.Unmarshal([]byte(info), &groupInfoSt)
			dataContact.PhoneNumber = chatIdParse[3:]
			dataContact.Name = groupInfoSt.Data.Name
			dataContact.ContactAttributes = model.CustomAttributes{
				GroupID: oneSenderWebhook.Chat,
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

func infoGroup(id string) string {
	url := viper.GetString("server.host") + "/group/info"
	method := "POST"
	payload := strings.NewReader(`{ "GroupJID": "` + id + `" }`)
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}
	req.Header.Add("token", "hjve6uly")
	req.Header.Add("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}

	fmt.Println(string(body))
	return string(body)
}

type GroupInfoResponse struct {
	Code    int64 `json:"code"`
	Data    Data  `json:"data"`
	Success bool  `json:"success"`
}

type Data struct {
	AnnounceVersionID             string        `json:"AnnounceVersionID"`
	DefaultMembershipApprovalMode string        `json:"DefaultMembershipApprovalMode"`
	DisappearingTimer             int64         `json:"DisappearingTimer"`
	GroupCreated                  string        `json:"GroupCreated"`
	IsAnnounce                    bool          `json:"IsAnnounce"`
	IsDefaultSubGroup             bool          `json:"IsDefaultSubGroup"`
	IsEphemeral                   bool          `json:"IsEphemeral"`
	IsLocked                      bool          `json:"IsLocked"`
	IsParent                      bool          `json:"IsParent"`
	Jid                           string        `json:"JID"`
	LinkedParentJID               string        `json:"LinkedParentJID"`
	MemberAddMode                 string        `json:"MemberAddMode"`
	Name                          string        `json:"Name"`
	NameSetAt                     string        `json:"NameSetAt"`
	NameSetBy                     string        `json:"NameSetBy"`
	OwnerJID                      string        `json:"OwnerJID"`
	ParticipantVersionID          string        `json:"ParticipantVersionID"`
	Participants                  []Participant `json:"Participants"`
	Topic                         string        `json:"Topic"`
	TopicDeleted                  bool          `json:"TopicDeleted"`
	TopicID                       string        `json:"TopicID"`
	TopicSetAt                    string        `json:"TopicSetAt"`
	TopicSetBy                    string        `json:"TopicSetBy"`
}

type Participant struct {
	AddRequest   interface{} `json:"AddRequest"`
	Error        int64       `json:"Error"`
	IsAdmin      bool        `json:"IsAdmin"`
	IsSuperAdmin bool        `json:"IsSuperAdmin"`
	Jid          string      `json:"JID"`
}
