/*
 *
 * request.go
 * waf
 *
 * Created by lintao on 2022/6/23 4:47 PM
 * Copyright © 2020-2022 LINTAO. All rights reserved.
 *
 */

package scanner

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/NSObjects/scout/internal/finger/ehole/request"
	"github.com/NSObjects/scout/internal/waf/gotestwaf/config"
	"github.com/NSObjects/scout/internal/waf/gotestwaf/encoder"
	"github.com/NSObjects/scout/internal/waf/gotestwaf/placeholder"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
	"time"
)

type Resp struct {
	Header     http.Header
	StatusCode int
	Body       []byte
}

type HTTPClient struct {
	client        *http.Client
	cookies       []*http.Cookie
	headers       map[string]string
	followCookies bool
}

func NewHTTPClient(cfg *config.Config) (*HTTPClient, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: !cfg.TLSVerify},
		IdleConnTimeout: time.Duration(cfg.IdleConnTimeout) * time.Second,
		MaxIdleConns:    cfg.MaxIdleConns,
	}

	if cfg.Proxy != "" {
		proxyURL, _ := url.Parse(cfg.Proxy)
		tr.Proxy = http.ProxyURL(proxyURL)
	}

	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}

	cl := &http.Client{
		Transport: tr,
		CheckRedirect: func() func(req *http.Request, via []*http.Request) error {
			redirects := 0
			return func(req *http.Request, via []*http.Request) error {
				if redirects > cfg.MaxRedirects {
					return errors.New("max redirect number exceeded")
				}
				redirects++
				return nil
			}
		}(),
		Jar: jar,
	}

	configuredHeaders := cfg.HTTPHeaders
	customHeader := strings.SplitN(cfg.AddHeader, ":", 2)
	if len(customHeader) > 1 {
		header := strings.TrimSpace(customHeader[0])
		value := strings.TrimSpace(customHeader[1])
		configuredHeaders[header] = value
	}

	return &HTTPClient{
		client:        cl,
		cookies:       cfg.Cookies,
		headers:       configuredHeaders,
		followCookies: cfg.FollowCookies,
	}, nil
}

func (c *HTTPClient) Send(
	ctx context.Context,
	targetURL, placeholderName, encoderName, payload string,
	testHeaderValue string,
) (Resp, error) {
	encodedPayload, err := encoder.Apply(encoderName, payload)
	if err != nil {
		return Resp{}, errors.Wrap(err, "encoding payload")
	}

	req, err := placeholder.Apply(targetURL, placeholderName, encodedPayload)
	if err != nil {
		return Resp{}, errors.Wrap(err, "apply placeholder")
	}
	req = req.WithContext(ctx)

	for header, value := range c.headers {
		req.Header.Set(header, value)
	}
	if _, ok := c.headers["Host"]; ok {
		req.Host = c.headers["Host"]
	}

	req.Header.Set("User-Agent", request.RandUserAgent())

	if testHeaderValue != "" {
		req.Header.Set("X-GoTestWAF-Test", testHeaderValue)
	}

	if len(c.cookies) > 0 && c.followCookies {
		c.client.Jar.SetCookies(req.URL, c.cookies)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return Resp{}, errors.Wrap(err, "sending http request")
	}
	defer func() {
		if err = resp.Body.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Resp{}, errors.Wrap(err, "reading response body")
	}

	if len(resp.Cookies()) > 0 {
		c.cookies = append(c.cookies, resp.Cookies()...)
	}

	return Resp{
		Body:       body,
		StatusCode: resp.StatusCode,
		Header:     resp.Header,
	}, nil
}
