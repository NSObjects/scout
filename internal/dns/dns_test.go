/*
 *
 * dns_test.go
 * dns
 *
 * Created by lintao on 2022/6/17 11:47 AM
 * Copyright © 2020-2022 LINTAO. All rights reserved.
 *
 */

package dns

import (
	"github.com/miekg/dns"
	"reflect"
	"testing"
)

func Test_dnsQuery_Query(t *testing.T) {
	type args struct {
		domain string
	}
	tests := []struct {
		name    string
		args    args
		want    []dns.RR
		wantErr bool
	}{
		{
			name:    "DNS查询",
			args:    args{domain: "facebook.com"},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := dnsQuery{}
			got, err := q.Query(tt.args.domain, A)
			if (err != nil) != tt.wantErr {
				t.Errorf("Query() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Query() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dnsQuery_Query1(t *testing.T) {
	type args struct {
		host string
		t    uint16
	}
	tests := []struct {
		name    string
		args    args
		want    []dns.RR
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := dnsQuery{}
			got, err := q.Query(tt.args.host, tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("Query() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Query() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dnsQuery_QueryA(t *testing.T) {
	type args struct {
		host string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := dnsQuery{}
			got, err := q.QueryA(tt.args.host)
			if (err != nil) != tt.wantErr {
				t.Errorf("QueryA() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QueryA() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dnsQuery_QueryDNS(t *testing.T) {
	type args struct {
		host string
		t    uint16
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "NDS查询",
			args: args{
				host: "www.cloudflare.com",
				t:    A,
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "NDS查询",
			args: args{
				host: "www.cloudflare.com",
				t:    CNAME,
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := dnsQuery{}
			got, err := q.QueryDNS(tt.args.host, tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("QueryDNS() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QueryDNS() got = %v, want %v", got, tt.want)
			}

		})
	}
}
