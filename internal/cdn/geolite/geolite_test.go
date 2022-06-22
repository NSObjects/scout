/*
 *
 * geolite_test.go
 * geolite
 *
 * Created by lintao on 2022/6/17 11:12 AM
 * Copyright © 2020-2022 LINTAO. All rights reserved.
 *
 */

package geolite

import (
	"testing"
)

func Test_geoLite_Lookup(t *testing.T) {
	type fields struct {
		path string
	}

	type args struct {
		ip string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{
			name:    "look up ip geo",
			fields:  fields{path: "./GeoLite2-ASN.mmdb"},
			args:    args{ip: "104.16.123.96"},
			wantErr: false,
			want:    true,
		},
		{
			name:    "look up ip geo",
			fields:  fields{path: "./GeoLite2-ASN.mmdb"},
			args:    args{ip: "203.119.207.243"},
			wantErr: false,
			want:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewGeoLite()
			got, err := g.Lookup(tt.args.ip)
			if (err != nil) != tt.wantErr {
				t.Errorf("Lookup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Lookup() got = %v, want %v", got, tt.want)
			}
		})
	}
}
