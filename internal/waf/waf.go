/*
 *
 * waf.go
 * waf
 *
 * Created by lintao on 2022/6/23 5:41 PM
 * Copyright © 2020-2022 LINTAO. All rights reserved.
 *
 */

package waf

import (
	"context"
	"fmt"
	"github.com/NSObjects/scout/internal/waf/config"
)

func Waf(url string) {
	normal, err := NormalRequest(url)
	if err != nil {
		panic(err)
	}

	xss, err := XssAttack(url)
	if err != nil {
		panic(err)
	}
	if normal.StatusCode != xss.StatusCode {
		fmt.Println("xss attack")
	}
}

func NormalRequest(url string) (Resp, error) {
	client, err := NewHTTPClient(&config.Config{})
	if err != nil {
		return Resp{}, err
	}
	return client.Send(context.Background(), url, "", "", "", "")
}

func XssAttack(url string) (Resp, error) {
	client, err := NewHTTPClient(&config.Config{})
	if err != nil {
		return Resp{}, err
	}
	return client.Send(context.Background(), url, "URLParam", "URL", "<script>alert('union select password from users')</script>", "")
}

func SqlAttack(url string) (Resp, error) {
	client, err := NewHTTPClient(&config.Config{})
	if err != nil {
		return Resp{}, err
	}
	return client.Send(context.Background(), url, "URLParam", "URL", "UNION SELECT ALL FROM information_schema AND ' or SLEEP(5) or '", "")
}
