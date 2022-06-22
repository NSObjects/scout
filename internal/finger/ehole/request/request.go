/*
 *
 * request.go
 * finger
 *
 * Created by lintao on 2022/6/14 4:03 PM
 * Copyright © 2020-2022 LINTAO. All rights reserved.
 *
 */

package request

import (
	"crypto/tls"
	"github.com/NSObjects/scout/internal/finger/ehole/encoding"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Resps struct {
	Url        string
	Body       string
	Header     map[string][]string
	Server     string
	StatusCode int
	Length     int
	Title      string
	JSUrl      []string
	FavHash    string
}

func Request(address string, proxy string) (*Resps, error) {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	if proxy != "" {
		transport.Proxy = func(_ *http.Request) (*url.URL, error) {
			return url.Parse(proxy)
		}
	}

	client := &http.Client{
		Timeout:   10 * time.Second,
		Transport: transport,
	}

	req, err := http.NewRequest("GET", address, nil)
	if err != nil {
		return nil, err
	}

	cookie := &http.Cookie{
		Name:  "rememberMe",
		Value: "Me",
	}

	req.AddCookie(cookie)
	req.Header.Set("Accept", "*/*;q=0.8")
	req.Header.Set("Connection", "close")
	req.Header.Set("User-Agent", RandUserAgent())
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	contentType := strings.ToLower(resp.Header.Get("Content-Type"))

	httpBody := encoding.ToUtf8(string(result), contentType)

	server := "None"
	if server = resp.Header.Get("Server"); server == "" {
		server = resp.Header.Get("X-Powered-By")
	}

	jsURL := JsJump(httpBody, address)

	favHash, err := favicon(httpBody, address)
	if err != nil {
		return nil, err
	}

	return &Resps{
		Url:        address,
		Body:       httpBody,
		Header:     resp.Header,
		Server:     server,
		StatusCode: resp.StatusCode,
		Length:     len(httpBody),
		Title:      title(httpBody),
		JSUrl:      jsURL,
		FavHash:    favHash,
	}, nil
}
