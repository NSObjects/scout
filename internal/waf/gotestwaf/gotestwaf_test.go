/*
 *
 * gotestwaf_test.go
 * gotestwaf
 *
 * Created by lintao on 2022/6/27 5:28 PM
 * Copyright © 2020-2022 LINTAO. All rights reserved.
 *
 */

package gotestwaf

import (
	"testing"
)

func Test_wafCheck_WafCheck(t *testing.T) {

	type args struct {
		url string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name:    "有Waf的域名",
			args:    args{url: "https://congjing.net"},
			want:    true,
			wantErr: false,
		},
		{
			name:    "无Waf的域名",
			args:    args{url: "https://baidu.com"},
			want:    false,
			wantErr: false,
		},
		{
			name:    "有Waf的域名",
			args:    args{url: "https://www.ctyun.cn"},
			want:    true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := NewWafCheck()
			got, err := w.WafCheck(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("WafCheck() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("WafCheck() got = %v, want %v", got, tt.want)
			}
		})
	}
}
