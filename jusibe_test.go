package go_jusibe

import (
	"testing"
)

var params map[string]string

func init() {
	params = make(map[string]string)
	params["SID"] = "AC0f30491286ab4abb4a108abefbd05d8a"
	params["TOKEN"] = "1dcf52d7a1f3853ed78f0ee20d056dd0"
	params["FROM"] = "+15005550006"
	params["TO"] = "+19135551234"
}

func TestSMS(t *testing.T) {
	msg := "Welcome to gotwilio"
	twilio := NewTwilioClient(params["SID"], params["TOKEN"])
	_, exc, err := twilio.SendSMS(params["FROM"], params["TO"], msg, "", "")
	if err != nil {
		t.Fatal(err)
	}

	if exc != nil {
		t.Fatal(exc)
	}
}

func TestSMS(t *testing.T) {
	msg := "Welcome to gotwilio"
	twilio := NewTwilioClient(params["SID"], params["TOKEN"])
	_, exc, err := twilio.SendSMS(params["FROM"], params["TO"], msg, "", "")
	if err != nil {
		t.Fatal(err)
	}

	if exc != nil {
		t.Fatal(exc)
	}
}

func TestSMS(t *testing.T) {
	msg := "Welcome to gotwilio"
	twilio := NewTwilioClient(params["SID"], params["TOKEN"])
	_, exc, err := twilio.SendSMS(params["FROM"], params["TO"], msg, "", "")
	if err != nil {
		t.Fatal(err)
	}

	if exc != nil {
		t.Fatal(exc)
	}
}