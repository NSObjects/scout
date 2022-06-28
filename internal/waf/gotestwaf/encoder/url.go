/*
 *
 * url.go
 * encoder
 *
 * Created by lintao on 2022/6/23 5:16 PM
 * Copyright © 2020-2022 LINTAO. All rights reserved.
 *
 */

package encoder

import "net/url"

type URLEncoder struct {
	name string
}

var DefaultURLEncoder = URLEncoder{name: "URL"}

func (enc URLEncoder) GetName() string {
	return enc.name
}

func (enc URLEncoder) Encode(data string) (string, error) {
	return url.QueryEscape(data), nil
}
