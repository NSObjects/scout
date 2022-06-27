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
	"fmt"
	"github.com/NSObjects/scout/internal/finger"
	"github.com/NSObjects/scout/internal/finger/ehole/config"
	"github.com/NSObjects/scout/internal/finger/ehole/request"
	"github.com/NSObjects/scout/pkg"
	"github.com/gookit/color"
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
		return nil
	}
	s.Rule = r
	return s
}

func (s *FinScan) Scan(url string) (finger.ScanResult, error) {
	data, err := request.Request(url, s.Proxy)
	if err != nil {
		url = strings.ReplaceAll(url, "https://", "http://")
		if data, err = request.Request(url, s.Proxy); err != nil {
			return finger.ScanResult{}, err
		}
	}

	var cms []string
	for _, fin := range s.Rule.Fingerprint {
		if c := s.cmsFinger(fin, data); c != "" {
			cms = append(cms, c)
		}
	}

	cmss := strings.Join(removeDuplicates(cms), ",")

	out := finger.ScanResult{Url: data.Url, Cms: cmss, Server: data.Server, StatusCode: data.StatusCode, Length: data.Length, Title: data.Title}
	if len(out.Cms) != 0 {
		outStr := fmt.Sprintf("[ %s | %s | %s | %d | %d | %s ]", out.Url, out.Cms, out.Server, out.StatusCode, out.Length, out.Title)
		color.RGBStyleFromString("237,64,35").Println(outStr)
	} else {
		outStr := fmt.Sprintf("[ %s | %s | %s | %d | %d | %s ]", out.Url, out.Cms, out.Server, out.StatusCode, out.Length, out.Title)
		fmt.Println(outStr)
	}

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
