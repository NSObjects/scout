/*
 *
 * parse.go
 * finger
 *
 * Created by lintao on 2022/6/14 4:03 PM
 * Copyright © 2020-2022 LINTAO. All rights reserved.
 *
 */

package request

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/go-resty/resty/v2"
	"github.com/twmb/murmur3"
	"log"
	"math/rand"
	"net/url"
	"regexp"
	"strings"
	"time"
)

var client = resty.New()

func init() {
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})

	client.SetTimeout(10 * time.Second)
}

func title(httpBody string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(httpBody))
	if err != nil {
		return "Not found"
	}
	return strings.Trim(strings.Replace(doc.Find("Title").Text(), "\n", "", -1), " ")
}

func favicon(httpBody string, address string) (string, error) {
	faviconPaths := regexps(`href="(.*?favicon....)"`, httpBody)
	var faviconPath string

	u, err := url.Parse(address)

	if err != nil {
		return "", err
	}

	address = u.Scheme + "://" + u.Host

	if len(faviconPaths) > 0 {
		fav := faviconPaths[0][1]
		if fav[:2] == "//" {
			faviconPath = "http:" + fav
		} else {
			if fav[:4] == "http" {
				faviconPath = fav
			} else {
				faviconPath = address + "/" + fav
			}
		}
	} else {
		faviconPath = address + "/favicon.ico"
	}

	return faviconHash(faviconPath)
}

func faviconHash(host string) (string, error) {
	client.R().SetHeader("User-Agent", RandUserAgent())
	resp, err := client.R().Get(host)
	if err != nil {
		return "", err
	}
	return mmh3Hash32(standBase64(resp.Body()))
}

func mmh3Hash32(raw []byte) (string, error) {
	h32 := murmur3.New32()
	_, err := h32.Write(raw)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%d", int32(h32.Sum32())), nil
}

func standBase64(raw []byte) []byte {
	bcs := base64.StdEncoding.EncodeToString(raw)
	var buffer bytes.Buffer
	for i := 0; i < len(bcs); i++ {
		ch := bcs[i]
		buffer.WriteByte(ch)
		if (i+1)%76 == 0 {
			buffer.WriteByte('\n')
		}
	}
	buffer.WriteByte('\n')
	return buffer.Bytes()
}

func regexps(regular string, resp string) [][]string {
	reg := regexp.MustCompile(regular)
	if reg == nil {
		log.Println("regexp err")
		return nil
	}

	return reg.FindAllStringSubmatch(resp, -1)
}

func RandUserAgent() string {
	ua := []string{
		"Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2227.1 Safari/537.36",
		"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2227.0 Safari/537.36",
		"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2227.0 Safari/537.36",
		"Mozilla/5.0 (Windows NT 6.3; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2226.0 Safari/537.36",
		"Mozilla/5.0 (Windows NT 6.1; WOW64; rv:40.0) Gecko/20100101 Firefox/40.1",
		"Mozilla/5.0 (Windows NT 6.3; rv:36.0) Gecko/20100101 Firefox/36.0",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10; rv:33.0) Gecko/20100101 Firefox/33.0",
		"Mozilla/5.0 (X11; Linux i586; rv:31.0) Gecko/20100101 Firefox/31.0",
		"Mozilla/5.0 (Windows NT 6.1; WOW64; rv:31.0) Gecko/20130401 Firefox/31.0",
		"Mozilla/5.0 (Windows NT 5.1; rv:31.0) Gecko/20100101 Firefox/31.0",
		"Mozilla/5.0 (Windows NT 6.1; WOW64; Trident/7.0; AS; rv:11.0) like Gecko",
		"Mozilla/5.0 (compatible, MSIE 11, Windows NT 6.3; Trident/7.0; rv:11.0) like Gecko",
		"Mozilla/5.0 (Windows; Intel Windows) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.67"}

	return ua[rand.Intn(len(ua))]
}

func jsJump(str string, url string) []string {
	regs := []string{`(window|top)\.location\.href = "(.*?)"`, `redirectUrl = "(.*?)"`, `<meta.*?http-equiv=.*?refresh.*?Url=(.*?)>`}
	var results []string
	for _, reg := range regs {
		result := regexps(reg, str)
		if result == nil {
			continue
		}

		for _, m := range result {
			s := len(m)
			if strings.Contains(m[s-1], "http") {
				continue
			}
			str2 := strings.Trim(m[s-1], "/")
			str2 = strings.ReplaceAll(str2, "../", "/")
			if len(str2) == 0 {
				continue
			}
			if str2[:1] == "/" {
				results = append(results, url+str2)
			} else {
				results = append(results, url+"/"+str2)
			}
		}
	}
	return results
}
