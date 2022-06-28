/*
 *
 * model.go
 * testcases
 *
 * Created by lintao on 2022/6/27 3:33 PM
 * Copyright © 2020-2022 LINTAO. All rights reserved.
 *
 */

package db

type Info struct {
	Payload            string
	Encoder            string
	Placeholder        string
	Set                string
	Case               string
	ResponseStatusCode int
	Reason             string
	Type               string
}

type Case struct {
	Payloads       []string `yaml:"payload"`
	Encoders       []string `yaml:"encoder"`
	Placeholders   []string `yaml:"placeholder"`
	Type           string   `yaml:"type"`
	Set            string
	Name           string
	IsTruePositive bool
}
