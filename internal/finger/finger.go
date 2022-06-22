/*
 *
 * finger.go
 * finger
 *
 * Created by lintao on 2022/6/22 4:33 PM
 * Copyright © 2020-2022 LINTAO. All rights reserved.
 *
 */

package finger

type ScanResult struct {
	Url        string `json:"url"`
	Cms        string `json:"cms"`
	Server     string `json:"server"`
	StatusCode int    `json:"status_code" `
	Length     int    `json:"length"`
	Title      string `json:"title"`
}
