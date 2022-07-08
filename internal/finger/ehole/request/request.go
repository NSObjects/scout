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
	"github.com/NSObjects/scout/internal/finger/ehole/encoding"
	"net/http"
)

type Response struct {
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

func Request(address string, proxy string) (*Response, error) {
	client.RemoveProxy()
	if proxy != "" {
		client.SetProxy(proxy)
	}

	client.SetCookie(&http.Cookie{
		Name:  "rememberMe",
		Value: "Me",
	})
	client.R().SetHeaders(map[string]string{
		"Accept":     "*/*;q=0.8",
		"Connection": "close",
		"User-Agent": RandUserAgent()})

	resp, err := client.R().Get(address)

	if err != nil {
		return nil, err
	}

	httpBody := encoding.Utf8(resp.String(), resp.Header().Get("Content-Type"))

	server := "None"

	if server = resp.Header().Get("Server"); server == "" {
		server = resp.Header().Get("X-Powered-By")
	}

	jsURL := jsJump(httpBody, address)

	favHash, err := favicon(httpBody, address)
	if err != nil {
		return nil, err
	}

	return &Response{
		Url:        address,
		Body:       httpBody,
		Header:     resp.Header(),
		Server:     server,
		StatusCode: resp.StatusCode(),
		Length:     len(httpBody),
		Title:      title(httpBody),
		JSUrl:      jsURL,
		FavHash:    favHash,
	}, nil
}
