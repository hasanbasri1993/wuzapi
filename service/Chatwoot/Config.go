package Chatwoot

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"regexp"
	"strings"
)

func getConfig() (origin string, account string, inbox string, error error) {
	var text = viper.GetString("chatwoot.baseUrl")
	var regex, err = regexp.Compile(`/(app|(api/v1))/accounts/\d*(/(inbox|inboxes)/\d*)?`)
	if err != nil {
		fmt.Println(err.Error())
		return "", "", "", errors.New("wrong format url")
	}
	var isMatch = regex.MatchString(text)
	if isMatch != true {
		fmt.Println("wrong format url")
		return "", "", "", errors.New("wrong format url")
	}
	var str = regex.Split(text, -1)
	var res1 = regex.FindAllString(text, 2)
	var strSplit = strings.Split(res1[0], "/")
	url := str[0]
	accountStr := strSplit[3]
	inboxStr := strSplit[5]
	return url, accountStr, inboxStr, nil
}

func SetInbox() {
	_, _, inbox, err := getConfig()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//   const updatePayload = {
	//            "channel": {
	//                "additional_attributes": {
	//                    "sessionId": this.client.getSessionId(),
	//                    "hostAccountNumber": `${this.accountNumber}`,
	//                    "instanceId": `${this.client.getInstanceId()}`
	//                }
	//            }
	//        }
	//        if (this.forceUpdateCwWebhook) updatePayload.channel['webhook_url'] = this.expectedSelfWebhookUrl
	//        const updateInboxPromise = this.cwReq('patch', `inboxes/${this.inboxId}`, updatePayload)
	if viper.GetString("chatwoot.forceUpdateCwWebhook") == "true" {
		expectedSelfWebhookUrl := viper.GetString("server.host") + "/chatwoot"
		payload := `{"channel": {"webhook_url": "` + expectedSelfWebhookUrl + `"}}`
		requestChatwoot("PATCH", "inboxes/"+inbox, strings.NewReader(payload))
	}
}
