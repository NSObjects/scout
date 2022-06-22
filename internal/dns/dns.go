/*
 *
 * dns.go
 * dns
 *
 * Created by lintao on 2022/6/17 11:26 AM
 * Copyright © 2020-2022 LINTAO. All rights reserved.
 *
 */

package dns

import (
	"github.com/miekg/dns"
	"time"
)

const (
	A     = dns.TypeA
	NS    = dns.TypeNS
	MX    = dns.TypeMX
	TXT   = dns.TypeTXT
	SOA   = dns.TypeSOA
	CNAME = dns.TypeCNAME
)

type dnsQuery struct {
}

func NewDNSQuery() dnsQuery {
	return dnsQuery{}
}

var dnsService = "223.5.5.5"

var client = &dns.Client{
	Timeout: 5 * time.Second,
}

func (q dnsQuery) Query(host string, t uint16) ([]dns.RR, error) {
	var lastErr error
	for i := 0; i < 3; i++ {
		dm := &dns.Msg{}
		dm.SetQuestion(dns.Fqdn(host), t)
		r, _, err := client.Exchange(dm, dnsService+":53")
		if err != nil {
			lastErr = err
			time.Sleep(1 * time.Second * time.Duration(i+1))
			continue
		}
		if len(r.Answer) > 0 {
			return r.Answer, nil
		}
	}

	return nil, lastErr
}

func (q dnsQuery) QueryA(host string) ([]string, error) {
	return q.QueryDNS(host, A)
}

func (q dnsQuery) QueryDNS(host string, t uint16) ([]string, error) {
	ips, err := q.Query(host, t)
	if err != nil {
		return nil, err
	}
	var result []string
	for _, ip := range ips {
		switch ip.(type) {
		case *dns.A:
			if t == A {
				result = append(result, ip.(*dns.A).A.String())
			}
		case *dns.CNAME:
			if t == CNAME {
				result = append(result, ip.(*dns.CNAME).Target)
			}
		case *dns.NS:
			if t == NS {
				result = append(result, ip.(*dns.NS).Ns)
			}

		case *dns.MX:
			if t == MX {
				result = append(result, ip.(*dns.MX).Mx)
			}
		case *dns.TXT:
			if t == TXT {
				result = append(result, ip.(*dns.TXT).Txt...)
			}

		case *dns.SOA:
			if t == SOA {
				result = append(result, ip.(*dns.SOA).Ns)
				result = append(result, ip.(*dns.SOA).Mbox)
			}
		}
	}

	return result, nil

}
