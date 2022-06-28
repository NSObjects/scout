/*
 *
 * gotestwaf.go
 * gotestwaf
 *
 * Created by lintao on 2022/6/27 5:03 PM
 * Copyright © 2020-2022 LINTAO. All rights reserved.
 *
 */

package gotestwaf

import (
	"context"
	"github.com/NSObjects/scout/internal/waf/gotestwaf/config"
	"github.com/NSObjects/scout/internal/waf/gotestwaf/db"
	"github.com/NSObjects/scout/internal/waf/gotestwaf/scanner"
	"github.com/NSObjects/scout/pkg"
	"github.com/pkg/errors"
	"log"
	"os"
)

type wafCheck struct {
	opts *config.Config
}

func NewWafCheck(opts ...config.Option) pkg.PluginWaf {
	opt := &config.Config{
		TestCasesPath:   "/Users/lintao/Documents/work/github.com/scout/internal/waf/gotestwaf/testcases",
		MaxIdleConns:    2,
		MaxRedirects:    50,
		IdleConnTimeout: 2,
		BlockStatusCode: 403,
		PassStatusCode:  200,
		Workers:         5,
		SendDelay:       400,
		RandomDelay:     400,
		HTTPHeaders: map[string]string{
			"Connection":                "close",
			"Pragma":                    "no-cache",
			"Cache-Control":             "no-cache",
			"Upgrade-Insecure-Requests": "1",
			"User-Agent":                "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.102 Safari/537.36",
			"Accept":                    `text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9`,
			"Sec-Fetch-Site":            "none",
			"Sec-Fetch-Mode":            "navigate",
			"Sec-Fetch-User":            "?1",
			"Sec-Fetch-Dest":            "document",
			"Accept-Encoding":           "gzip, deflate",
			"Accept-Language":           "en-US,en;q=0.9",
		},
	}

	for _, v := range opts {
		v(opt)
	}

	return wafCheck{opt}
}

func (w wafCheck) WafCheck(url string) (bool, error) {
	logger := log.New(os.Stdout, "GOTESTWAF : ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
	w.opts.URL = url
	client, err := scanner.NewHTTPClient(w.opts)
	if err != nil {
		panic(err)
	}
	cases, err := db.LoadTestCases(w.opts)
	if err != nil {
		panic(err)
	}

	db := db.NewDB(cases)

	s := scanner.New(db, logger, w.opts, client, false)
	check, code, err := s.BenignPreCheck(url)
	if check {
		return false, errors.Errorf("url %s is block", url)
	}

	preCheck, statusCode, err := s.PreCheck(url)
	if preCheck {
		return true, nil
	}

	if code != statusCode {
		return true, nil
	}

	return s.RunWaf(context.Background())
}
