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

type Option func(c *Config)

func SetCookies(cookies []*http.Cookie) Option {
	return func(o *Config) { o.Cookies = cookies }
}
func SetURL(url string) Option {
	return func(o *Config) { o.URL = url }
}
func SetWebSocketURL(url string) Option {
	return func(o *Config) { o.WebSocketURL = url }
}
func SetGRPCPort(grpcPort uint16) Option {
	return func(o *Config) { o.GRPCPort = grpcPort }
}
func SetHTTPHeaders(header map[string]string) Option {
	return func(o *Config) { o.HTTPHeaders = header }
}
func SetTLSVerify(tlsVerify bool) Option {
	return func(o *Config) { o.TLSVerify = tlsVerify }
}
func SetProxy(proxy string) Option {
	return func(o *Config) { o.Proxy = proxy }
}
func SetMaxIdleConns(maxIdleConns int) Option {
	return func(o *Config) { o.MaxIdleConns = maxIdleConns }
}
func SetMaxRedirects(maxRedirects int) Option {
	return func(o *Config) { o.MaxRedirects = maxRedirects }
}
func SetIdleConnTimeout(idleConnTimeout int) Option {
	return func(o *Config) { o.IdleConnTimeout = idleConnTimeout }
}
func SetFollowCookies(followCookies bool) Option {
	return func(o *Config) { o.FollowCookies = followCookies }
}
func SetBlockStatusCode(blockStatusCode int) Option {
	return func(o *Config) { o.BlockStatusCode = blockStatusCode }
}
func SetPassStatusCode(passStatusCode int) Option {
	return func(o *Config) { o.PassStatusCode = passStatusCode }
}
func SetBlockRegex(blockRegex string) Option {
	return func(o *Config) { o.BlockRegex = blockRegex }
}
func SetPassRegex(passRegex string) Option {
	return func(o *Config) { o.PassRegex = passRegex }
}
func SetWorkers(workers int) Option {
	return func(c *Config) {
		c.Workers = workers
	}
}
func SetNonBlockedAsPassed(nonBlockedAsPassed bool) Option {
	return func(c *Config) {
		c.NonBlockedAsPassed = nonBlockedAsPassed
	}
}
func SetRandomDelay(randomDelay int) Option {
	return func(c *Config) {
		c.RandomDelay = randomDelay
	}
}
func SetSendDelay(sendDelay int) Option {
	return func(c *Config) {
		c.SendDelay = sendDelay
	}
}
func SetReportPath(reportPath string) Option {
	return func(c *Config) {
		c.ReportPath = reportPath
	}
}
func SetReportName(reportName string) Option {
	return func(c *Config) {
		c.ReportName = reportName
	}
}
func SetTestCase(testCase string) Option {
	return func(c *Config) {
		c.TestCase = testCase
	}
}
func SetTestCasesPath(testCasesPath string) Option {
	return func(c *Config) {
		c.TestCasesPath = testCasesPath
	}
}
func SetTestSet(testSet string) Option {
	return func(c *Config) {
		c.TestSet = testSet
	}
}
func SetWAFName(wafName string) Option {
	return func(c *Config) {
		c.WAFName = wafName
	}
}
func SetIgnoreUnresolved(ignoreUnresolved bool) Option {
	return func(c *Config) {
		c.IgnoreUnresolved = ignoreUnresolved
	}
}
func SetBlockConnReset(blockConnReset bool) Option {
	return func(c *Config) {
		c.BlockConnReset = blockConnReset
	}
}
func SetSkipWAFBlockCheck(skipWAFBlockCheck bool) Option {
	return func(c *Config) {
		c.SkipWAFBlockCheck = skipWAFBlockCheck
	}
}
func SetAddHeader(sddHeader string) Option {
	return func(c *Config) {
		c.AddHeader = sddHeader
	}
}
func SetRenderToHTML(renderToHTML bool) Option {
	return func(c *Config) {
		c.RenderToHTML = renderToHTML
	}
}
