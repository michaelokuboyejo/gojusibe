package go_jusibe

import (
	"net/http"
	"net/url"
	"strings"
)

const jusibeBaseUrl = "https://jusibe.com/smsapi"

type Jusibe struct {
	PublicKey string
	AccessToken string
	HttpClient *http.Client
}

func JusibeClient(publicKey, accessToken string, HTTPClient *http.Client) *Jusibe {

	if HTTPClient == nil {
		HTTPClient = http.DefaultClient
	}
	return &Jusibe{publicKey, accessToken, HTTPClient}
}

func (jusibe *Jusibe) post(formValues url.Values, path string) (*http.Response, error) {
	requestUrl := jusibeBaseUrl + path
	jusibeRequest, err := http.NewRequest("POST", requestUrl, strings.NewReader(formValues.Encode()))
	if err != nil {
		return nil, err
	}
	jusibeRequest.SetBasicAuth(jusibe.PublicKey, jusibe.AccessToken)
	jusibeRequest.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := jusibe.HttpClient
	if client == nil {
		client = http.DefaultClient
	}
	return client.Do(jusibeRequest)
}

func (jusibe *Jusibe) get(path string) (*http.Response, error) {
	jusibeRequestUrl := jusibeBaseUrl + path
	jusibeRequest, err := http.NewRequest("GET", jusibeRequestUrl, nil)

	if err != nil{
		return nil, err
	}

	jusibeRequest.SetBasicAuth(jusibe.PublicKey, jusibe.AccessToken)

	client := jusibe.HttpClient
	if client == nil {
		client = http.DefaultClient
	}
	return client.Do(jusibeRequest)
}