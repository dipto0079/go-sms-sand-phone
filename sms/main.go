package sms

import (
	"encoding/json"

	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

func SendSms(accountSid, authToken, to, from, message string) (resp *twilioApi.ApiV2010Message, err error, response string) {
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

	params := &twilioApi.CreateMessageParams{}
	params.SetTo(to)
	params.SetFrom(from)
	params.SetBody(message)

	resp, err = client.Api.CreateMessage(params)
	if err != nil {
		return resp, err, "message"
	} else {
		response, _ := json.Marshal(*resp)
		return resp, err, string(response)
	}
}
