package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

type MakeRequestCfg struct {
	Method  string
	Url     string
	Body    string
	Headers map[string]string
}

type MakeRequestRes struct {
	Response *http.Response
	Body string
}

func MakeRequest(cfg *MakeRequestCfg) (*MakeRequestRes, error) {
	client := &http.Client{}

	bodyInpBuff := bytes.NewBufferString(cfg.Body)

	req, err := http.NewRequest(cfg.Method, cfg.Url, bodyInpBuff)
	if err != nil {
		return nil, err
	}

	for key, val := range cfg.Headers {
		req.Header.Set(key, val)
	}
	
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	respBodyData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return &MakeRequestRes{
		Response: resp, 
		Body: string(respBodyData), 
	}, nil
}

func LogRes(mRes *MakeRequestRes) error {
	fmt.Println(Yellow + "=== META" + Reset)
	fmt.Printf(" Status: %s\n Proto: %s\n", mRes.Response.Status, mRes.Response.Proto)

	fmt.Println(Magenta + "=== HEADER" + Reset)
	for key, vals := range mRes.Response.Header {
		fmt.Printf(" %s: %s\n", key, vals)
	}

	fmt.Println(Blue + "=== BODY" + Reset)

	fmt.Println(mRes.Body)

	return nil
}