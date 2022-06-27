/*
 *
 * waf_test.go
 * waf
 *
 * Created by lintao on 2022/6/23 5:54 PM
 * Copyright © 2020-2022 LINTAO. All rights reserved.
 *
 */

package waf

import "testing"

func TestWaf(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "xss测试",
			args: args{url: "https://congjing.net"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Waf(tt.args.url)
		})
	}
}
