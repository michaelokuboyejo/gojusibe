package gojusibe

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
	"github.com/araddon/dateparse"
)

type SmsResponse struct {
	Status string `json:"status"`
	MessageId string `json:"message_id"`
	SmsCreditsUsed int `json:"sms_credits_used"`
}

type SmsCreditsResponse struct {
	SmsCredits string `json:"sms_credits"`
}

type DeliveryStatusResponse struct {
	MessageId string `json:"message_id"`
	Status string `json:"status"`
	DateSent string `json:"date_sent"`
	DateDelivered string `json:"date_delivered"`
}

func (statusResponse *DeliveryStatusResponse) DateSentAsTime() (time.Time, error) {
	return dateparse.ParseLocal(statusResponse.DateSent)
}

func (statusResponse *DeliveryStatusResponse) DateDeliveredAsTime() (time.Time, error) {
	return dateparse.ParseLocal(statusResponse.DateDelivered)
}


func (jusibe *Jusibe) SendSms(from, to, message string) (smsResponse * SmsResponse, err error) {

	formValues := initFormValues(from, to, message)
	smsRequestpath := "/send_sms"
	res, err := jusibe.post(formValues, smsRequestpath)

	if err != nil {
		return smsResponse, err
	}

	defer res.Body.Close()

	responseBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return smsResponse, err
	}

	if res.StatusCode != http.StatusOK {
		err = json.Unmarshal(responseBody, smsResponse)
		return smsResponse, err
	}

	smsResponse = new(SmsResponse)
	err = json.Unmarshal(responseBody, smsResponse)
	return smsResponse, err
}


func (jusibe *Jusibe) CheckSmsCredits() (smsCreditsResponse *SmsCreditsResponse, err error) {

	smsCreditsBalancePath := "/sms_credits"

	res, err := jusibe.get(smsCreditsBalancePath)
	if err != nil {
		return smsCreditsResponse, err
	}
	defer res.Body.Close()

	responseBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return smsCreditsResponse, err
	}

	if res.StatusCode != http.StatusOK {
		err = json.Unmarshal(responseBody, smsCreditsResponse)
		return smsCreditsResponse, err
	}

	smsCreditsResponse = new(SmsCreditsResponse)
	err = json.Unmarshal(responseBody, smsCreditsResponse)
	return smsCreditsResponse, err

}

func (jusibe *Jusibe) CheckDeliveryStatus(messageId string) (deliveryStatusResponse *DeliveryStatusResponse, err error) {

	deliveryStatusPath := ""

	res, err := jusibe.get(deliveryStatusPath)
	if err != nil {
		return deliveryStatusResponse, err
	}
	defer res.Body.Close()

	responseBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return deliveryStatusResponse, err
	}

	if res.StatusCode != http.StatusOK {
		err = json.Unmarshal(responseBody, deliveryStatusResponse)
		return deliveryStatusResponse, err
	}

	deliveryStatusResponse = new(DeliveryStatusResponse)
	err = json.Unmarshal(responseBody, deliveryStatusResponse)
	return deliveryStatusResponse, err
}

func initFormValues(from, to, message string) url.Values {

	formValues := url.Values{}

	formValues.Set("to", to)
	formValues.Set("from", from)
	formValues.Set("message", message)

	return formValues
}