package main

import (
	"bytes"
	"io"
	"net/http"
	"testing"
)

func TestGet(t *testing.T) {
	resp, err := http.Get("http://example.com/")
	if err != nil {
		t.Fatal(err)
	}

	respBodyData, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(string(respBodyData))
}

func TestPost(t *testing.T) {

	// this is data we receive from user
	bodyInpStr := `{"email":"cong@gmail.com","password":"123456Aa"}`

	// convert it to buffer
	bodyInpBuff := bytes.NewBufferString(bodyInpStr)
	
	// resp return resp
	resp, err := http.Post(
		"http://localhost:8000/api/auth/login",
		"application/json",
		bodyInpBuff,
	)

	if err != nil {
		t.Fatal(err)
	}

	respBodyData, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(respBodyData))

	// get body data similar get request
}

func TestDo(t *testing.T) {

	client := &http.Client{}

	// this is data we receive from user
	bodyInpStr := `{"email":"cong@gmail.com","password":"123456Aa"}`

	// convert it to buffer
	bodyInpBuff := bytes.NewBufferString(bodyInpStr)

	req, err := http.NewRequest(
		"POST", 
		"http://localhost:8000/api/auth/login",
		bodyInpBuff,
	)
	if err != nil {
		t.Fatal(err)
	}

	// you can config like this 
	// set(key, val)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)

	if err != nil {
		t.Fatal(err)
	}

	respBodyData, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(string(respBodyData))
}

func TestMakeRequest(t *testing.T) {
	res, err := MakeRequest(&MakeRequestCfg{
		Method: "POST",
		Url: "http://localhost:8000/api/auth/login",
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: `{"email":"cong@gmail.com","password":"123456Aa"}`,
	})

	if err != nil {
		t.Fatal(err)
	}

	t.Log(res.Body)
}