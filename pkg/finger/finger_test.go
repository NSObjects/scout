/*
 *
 * finger_test.go
 * finger
 *
 * Created by lintao on 2022/6/22 4:44 PM
 * Copyright © 2020-2022 LINTAO. All rights reserved.
 *
 */

package finger

import (
	"github.com/NSObjects/scout/internal/finger"
	"github.com/NSObjects/scout/pkg"
	"reflect"
	"testing"
)

func Test_fingerCheck_Scan(t *testing.T) {
	type fields struct {
		plugins []pkg.PluginFinger
	}
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    finger.ScanResult
		wantErr bool
	}{
		{
			name: "指纹识别",
			args: args{url: "https://syxh.club/"},
			want: finger.ScanResult{
				Url:        "https://syxh.club/",
				Cms:        "WordPress",
				Server:     "nginx",
				StatusCode: 200,
				Length:     52802,
				Title:      "平江摄影协会",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := NewFinger()
			got, err := f.Scan(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("Scan() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Scan() got = %v, want %v", got, tt.want)
			}
		})
	}
}
