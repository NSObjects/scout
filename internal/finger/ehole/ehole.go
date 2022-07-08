/*
 *
 * ehole.go
 * finger
 *
 * Created by lintao on 2022/6/22 2:58 PM
 * Copyright © 2020-2022 LINTAO. All rights reserved.
 *
 */

package ehole

import (
	"encoding/json"
	"github.com/NSObjects/scout/internal/finger"
	"github.com/NSObjects/scout/internal/finger/ehole/config"
	"github.com/NSObjects/scout/internal/finger/ehole/request"
	"github.com/NSObjects/scout/internal/workpool"
	"github.com/NSObjects/scout/pkg"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

type FinScan struct {
	Proxy string
	Rule  *config.Rule
}

func NewFingerScan(proxy string) pkg.PluginFinger {
	s := &FinScan{
		Proxy: proxy,
	}
	r, err := config.WebFingerRule()
	if err != nil {
		panic(err)
	}
	s.Rule = r
	return s
}

func (s *FinScan) Scan(addr string) (finger.ScanResult, error) {
	parse, err := url.Parse(addr)
	if err != nil {
		return finger.ScanResult{}, err
	}

	if parse.Scheme != "http" && parse.Scheme != "https" {
		addr = "https://" + addr
	}

	data, err := request.Request(addr, s.Proxy)

	if err != nil || data.StatusCode != http.StatusOK {
		if parse.Scheme == "https" {
			addr = strings.Replace(addr, "https", "http", 1)
		} else {
			addr = strings.Replace(addr, "http", "https", 1)
		}

		d, err := request.Request(addr, s.Proxy)
		if err != nil && data == nil {
			return finger.ScanResult{}, err
		}
		if d != nil {
			data = d
		}
	}
	pool := workpool.NewWorkerPool(20)

	var cms []string
	for _, fin := range s.Rule.Fingerprint {
		fi := fin
		pool.AddTask(func() {
			if c := s.cmsFinger(fi, data); c != "" {
				cms = append(cms, c)
			}
		})
	}
	pool.Wait()

	cmss := strings.Join(removeDuplicates(cms), ",")

	out := finger.ScanResult{Url: data.Url, Cms: cmss, Server: data.Server, StatusCode: data.StatusCode, Length: data.Length, Title: data.Title}
	return out, nil
}

func removeDuplicates(arr []string) []string {
	wordsString := map[string]bool{}
	for i := range arr {
		wordsString[arr[i]] = true
	}
	var desiredOutput []string
	for j := range wordsString {
		desiredOutput = append(desiredOutput, j)
	}
	return desiredOutput
}

func (s *FinScan) cmsFinger(fin config.Fingerprint, data *request.Response) string {
	headers := MapToJson(data.Header)
	switch fin.Location {
	case "body":
		if fin.Method == "keyword" && isKeyword(data.Body, fin.Keyword) {
			return fin.Cms
		}
		if fin.Method == "faviconhash" && data.FavHash == fin.Keyword[0] {
			return fin.Cms
		}
		if fin.Method == "regular" && isRegular(data.Body, fin.Keyword) {
			return fin.Cms
		}
	case "header":
		if fin.Method == "keyword" && isKeyword(headers, fin.Keyword) {
			return fin.Cms
		}
		if fin.Method == "regular" && isRegular(headers, fin.Keyword) {
			return fin.Cms
		}

	case "title":
		if fin.Method == "keyword" && isKeyword(data.Title, fin.Keyword) {
			return fin.Cms
		}
		if fin.Method == "regular" && isRegular(data.Title, fin.Keyword) {
			return fin.Cms
		}
	}
	return ""
}

func isKeyword(str string, keyword []string) bool {
	x := true
	for _, k := range keyword {
		if !strings.Contains(str, k) {
			x = false
		}
	}
	return x
}

func isRegular(str string, keyword []string) bool {
	x := true
	for _, k := range keyword {
		re := regexp.MustCompile(k)
		if !re.Match([]byte(str)) {
			x = false
		}
	}
	return x
}

func MapToJson(param map[string][]string) string {
	dataType, _ := json.Marshal(param)
	dataString := string(dataType)
	return dataString
}
