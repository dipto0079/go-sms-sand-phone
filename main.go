package main

import (
	"fmt"
	sms "go-sms/sms"
)

func main() {
	//You will have to replace the placeholders for the Twilio
	//account SID, auth token, and phone numbers with your own value
	//Replace this with your own account SID and auth token
	//Replace this with your own phone number

	accountSid := ""
	authToken := ""
	setTo := ""
	setFrom := ""
	yourSms := ""

	resp, err, response := sms.SendSms(accountSid, authToken, setTo, setFrom, yourSms)
	fmt.Printf("resp: %+v", resp)
	fmt.Printf("response: %+v", response)

	if err != nil {
		fmt.Println(err.Error())
	}
}
