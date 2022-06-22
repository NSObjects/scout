/*
 *
 * cdn_test.go
 * cdn
 *
 * Created by lintao on 2022/6/21 3:30 PM
 * Copyright © 2020-2022 LINTAO. All rights reserved.
 *
 */

package cdn

import (
	"github.com/NSObjects/scout/pkg"
	"testing"
)

func Test_cdnCheck_HasCDN(t *testing.T) {
	type fields struct {
		plugins []pkg.PluginCDN
	}
	type args struct {
		target string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "CDN检测 cloudflare 有CDN",
			args: args{target: "www.cloudflare.com"},
			want: true,
		},

		{
			name: "CDN检测 moresec 没有CDN",
			args: args{target: "www.moresec.cn"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewCDNCheck()
			if got := c.HasCDN(tt.args.target); got != tt.want {
				t.Errorf("HasCDN() = %v, want %v", got, tt.want)
			}

		})
	}
}
