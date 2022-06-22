/*
 *
 * rule.go
 * finger
 *
 * Created by lintao on 2022/6/14 4:03 PM
 * Copyright © 2020-2022 LINTAO. All rights reserved.
 *
 */

package config

import "encoding/json"

var webFingerRule *Rule

type Rule struct {
	Fingerprint []Fingerprint `json:"fingerprint" `
}

type Fingerprint struct {
	Cms      string   `json:"cms" `
	Method   string   `json:"method" `
	Location string   `json:"location" `
	Keyword  []string `json:"keyword" `
}

func WebFingerRule() (*Rule, error) {
	if webFingerRule != nil {
		return webFingerRule, nil
	}

	if err := json.Unmarshal([]byte(fingerRule), &webFingerRule); err != nil {
		return nil, err
	}

	return webFingerRule, nil
}
