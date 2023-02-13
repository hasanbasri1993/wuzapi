package model

type OneSenderWebhook struct {
	Chat             string `json:"chat"`
	Sender           string `json:"sender"`
	SenderPushName   string `json:"sender_push_name"`
	SenderPhone      string `json:"sender_phone"`
	IsGroup          bool   `json:"is_group"`
	IsFromMe         bool   `json:"is_from_me"`
	MessageID        string `json:"message_id"`
	MessageType      string `json:"message_type"`
	MessageText      string `json:"message_text"`
	MessageTextID    string `json:"message_text_id"`
	MessageTimestamp string `json:"message_timestamp"`
	AttachmentURL    string `json:"attachment_url"`
	AttachmentID     string `json:"attachment_id"`
	AttachmentType   string `json:"attachment_type"`
}

//     "Info": {
//      "Chat": "6282213542319@s.whatsapp.net",
//      "Sender": "6282213542319.0:81@s.whatsapp.net",
//      "IsFromMe": false,
//      "IsGroup": false,
//      "BroadcastListOwner": "",
//      "ID": "3EB0ADAD43F6B126CB46",
//      "Type": "",
//      "PushName": "Hasan Basri",
//      "Timestamp": "2023-02-12T14:12:08+07:00",
//      "Category": "",
//      "Multicast": false,
//      "MediaType": "",
//      "VerifiedName": null,
//      "DeviceSentMeta": null
//    },

type OnesenderText struct {
	RecipientType string `json:"recipient_type"`
	To            string `json:"to"`
	Type          string `json:"type"`
	Text          Text   `json:"text"`
}

type Text struct {
	Body string `json:"body,omitempty"`
}

type OnesenderDocument struct {
	RecipientType string   `json:"recipient_type"`
	To            string   `json:"to"`
	Type          string   `json:"type"`
	Document      Document `json:"document"`
}

type Document struct {
	Link string `json:"link"`
}

type OnesenderImage struct {
	RecipientType string `json:"recipient_type"`
	To            string `json:"to"`
	Type          string `json:"type"`
	Image         Image  `json:"image"`
}

type Image struct {
	Link    string `json:"link"`
	Caption string `json:"caption"`
}

type OnesenderRespond struct {
	Code     int64              `json:"code"`
	Messages []OnesenderMessage `json:"messages"`
}

type OnesenderMessage struct {
	ID            string `json:"id"`
	Type          string `json:"type"`
	To            string `json:"to"`
	RecipientType string `json:"recipient_type"`
}
