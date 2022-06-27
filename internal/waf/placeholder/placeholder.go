/*
 *
 * placeholder.go
 * placeholder
 *
 * Created by lintao on 2022/6/23 5:17 PM
 * Copyright © 2020-2022 LINTAO. All rights reserved.
 *
 */

package placeholder

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"
)

const Seed = 5

type Placeholder interface {
	GetName() string
	CreateRequest(url, data string) (*http.Request, error)
}

var Placeholders map[string]Placeholder

func init() {
	Placeholders = make(map[string]Placeholder)
	Placeholders[DefaultURLParam.GetName()] = DefaultURLParam

}

func Apply(host, placeholder, data string) (*http.Request, error) {
	pl, ok := Placeholders[placeholder]
	if !ok {
		return http.NewRequest("GET", host, nil)
	}
	req, err := pl.CreateRequest(host, data)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func RandomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
