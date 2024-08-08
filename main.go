package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	url = flag.String("u", "", "Url request")
	method = flag.String("m", "GET", "Method request") 
	body = flag.String("b", "", "Body request")
	fileCfgPath = flag.String("f", "", "Config file")
)

type Headers map[string]string

func (hs *Headers) String() string {
	return fmt.Sprint(*hs)
}

func (hs *Headers) Set(val string) error {
    parts := strings.Split(val, ":")
	if len(parts) == 2 {
		key, val := parts[0], parts[1]
		(*hs)[key] = val 
	} else {
		return errors.New("header wrong format")
	}

	return nil
}

type Config struct {
	BASE_API string
	HEADERS map[string]string
}

func main() {
	log.SetFlags(0)

	var headers = make(Headers)
	flag.Var(&headers, "h", "Header")
	flag.Parse()

	if *url == "" {
		log.Fatal("Url require, type --help for more info")
	}

	if *fileCfgPath != "" {
		cfgContent, err := os.ReadFile(*fileCfgPath)
		if err != nil {
			log.Fatal(err)
		}

		var cfg Config
		err = json.Unmarshal(cfgContent, &cfg)
		if err != nil {
			log.Fatal(err)
		}

	
		*url = cfg.BASE_API + *url
		for key, val := range cfg.HEADERS {
			headers[key] = val
		}
	}

	res, err := MakeRequest(&MakeRequestCfg{
		Method: *method,
		Url: *url,
		Headers: headers,
		Body: *body,
	})
	if err != nil {
		log.Fatal(err)
	}

	err = LogRes(res)
	if err != nil {
		log.Fatal(err)
	}
}
