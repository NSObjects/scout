/*
 *
 * placeholder.go
 * placeholder
 *
 * Created by lintao on 2022/6/23 5:17 PM
 * Copyright © 2020-2022 LINTAO. All rights reserved.
 *
 */

package placeholder

import (
	"net/http"
)

const Seed = 5

type Placeholder interface {
	GetName() string
	CreateRequest(url, data string) (*http.Request, error)
}

var Placeholders map[string]Placeholder

func init() {
	Placeholders = make(map[string]Placeholder)
	Placeholders[DefaultURLParam.GetName()] = DefaultURLParam
	Placeholders[DefaultHeader.GetName()] = DefaultHeader
	Placeholders[DefaultHTMLForm.GetName()] = DefaultHTMLForm
	Placeholders[DefaultHTMLMultipartForm.GetName()] = DefaultHTMLMultipartForm
	Placeholders[DefaultJSONBody.GetName()] = DefaultJSONBody
	Placeholders[DefaultJSONRequest.GetName()] = DefaultJSONRequest
	Placeholders[DefaultRequestBody.GetName()] = DefaultRequestBody
	Placeholders[DefaultSOAPBody.GetName()] = DefaultSOAPBody
	Placeholders[DefaultURLParam.GetName()] = DefaultURLParam
	Placeholders[DefaultURLPath.GetName()] = DefaultURLPath
	Placeholders[DefaultXMLBody.GetName()] = DefaultXMLBody

}

func Apply(host, placeholder, data string) (*http.Request, error) {
	pl, ok := Placeholders[placeholder]
	if !ok {
		return http.NewRequest("GET", host, nil)
	}
	req, err := pl.CreateRequest(host, data)
	if err != nil {
		return nil, err
	}

	return req, nil
}
