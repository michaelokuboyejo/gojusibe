
[![Build Status](https://travis-ci.org/michaelokuboyejo/go-jusibe.svg?branch=master)](https://travis-ci.org/michaelokuboyejo/go-jusibe)

## Overview
This is an unofficial Golang library for [Jusibe](http://www.jusibe.com/). This library supports sending sms, checking available credits and delivery status of SMS's sent using the Jusibe SMS HTTP API.

## License
Go-Jusibe is licensed under a BSD license.

## Installation
To install this package, simply run `go get github.com/michaelokuboyejo/go_jusibe`.

## Example

	package main

	import (
		"github.com/michaelokuboyejo/gojusibe"
	)

	func main() {
		publicKey := "yourPublicKey"
		accessToken := "yourAccessToken"
		jusibe := gojusibe.JusibeClient(publicKey, accessToken)

		//sendSMS
		from := "Spiderman"
		to := "+23423456789"
		message := "Hi There, Mate!"
		jusibe.SendSMS(from, to, message)
		
		//check Available Credits
		//check Delivery status of message
	}
