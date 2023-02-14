package Chatwoot

import (
	"wuzapi/model"
)

func IncomingMessageApi(c model.OneSenderWebhook) model.SendMessageResult {
	rrContact := prosesContact(c)
	rrConversation := prosesConversation(rrContact)
	rrSend := praSendMessage(rrConversation, c)
	return rrSend
}
