/*
 *
 * ip_test.go
 * ip
 *
 * Created by lintao on 2022/6/17 11:10 AM
 * Copyright © 2020-2022 LINTAO. All rights reserved.
 *
 */

package ip

import (
	"reflect"
	"testing"
)

func TestNewIPRange(t *testing.T) {
	tests := []struct {
		name string
		want ipRange
	}{
		{
			name: "创建",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewIPRange(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewIPRange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ipRange_LookUp(t *testing.T) {

	type args struct {
		ip string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "搜索IP",
			args: args{
				ip: "104.16.123.96",
			},
			want:    true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewIPRange()
			got, err := r.Lookup(tt.args.ip)
			if (err != nil) != tt.wantErr {
				t.Errorf("LookUp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("LookUp() got = %v, want %v", got, tt.want)
			}
		})
	}
}
