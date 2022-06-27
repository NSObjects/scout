/*
 *
 * config.go
 * config
 *
 * Created by lintao on 2022/6/23 5:13 PM
 * Copyright © 2020-2022 LINTAO. All rights reserved.
 *
 */

package config

import "net/http"

type Config struct {
	Cookies            []*http.Cookie
	URL                string            `mapstructure:"url"`
	WebSocketURL       string            `mapstructure:"wsURL"`
	GRPCPort           uint16            `mapstructure:"grpcPort"`
	HTTPHeaders        map[string]string `mapstructure:"headers"`
	TLSVerify          bool              `mapstructure:"tlsVerify"`
	Proxy              string            `mapstructure:"proxy"`
	MaxIdleConns       int               `mapstructure:"maxIdleConns"`
	MaxRedirects       int               `mapstructure:"maxRedirects"`
	IdleConnTimeout    int               `mapstructure:"idleConnTimeout"`
	FollowCookies      bool              `mapstructure:"followCookies"`
	BlockStatusCode    int               `mapstructure:"blockStatusCode"`
	PassStatusCode     int               `mapstructure:"passStatusCode"`
	BlockRegex         string            `mapstructure:"blockRegex"`
	PassRegex          string            `mapstructure:"passRegex"`
	NonBlockedAsPassed bool              `mapstructure:"nonBlockedAsPassed"`
	Workers            int               `mapstructure:"workers"`
	RandomDelay        int               `mapstructure:"randomDelay"`
	SendDelay          int               `mapstructure:"sendDelay"`
	ReportPath         string            `mapstructure:"reportPath"`
	ReportName         string            `mapstructure:"reportName"`
	TestCase           string            `mapstructure:"testCase"`
	TestCasesPath      string            `mapstructure:"testCasesPath"`
	TestSet            string            `mapstructure:"testSet"`
	WAFName            string            `mapstructure:"wafName"`
	IgnoreUnresolved   bool              `mapstructure:"ignoreUnresolved"`
	BlockConnReset     bool              `mapstructure:"blockConnReset"`
	SkipWAFBlockCheck  bool              `mapstructure:"skipWAFBlockCheck"`
	AddHeader          string            `mapstructure:"addHeader"`
	RenderToHTML       bool              `mapstructure:"renderToHTML"`
}
