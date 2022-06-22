/*
 *
 * cdn.go
 * cdn
 *
 * Created by lintao on 2022/6/14 4:03 PM
 * Copyright © 2020-2022 LINTAO. All rights reserved.
 *
 */

package cdn

import (
	"fmt"
	"github.com/NSObjects/scout/internal/dns"
	"github.com/NSObjects/scout/pkg"
	"github.com/NSObjects/scout/pkg/plugin"
	"regexp"
)

func NewCDNCheck() CheckCDN {
	return cdnCheck{plugins: plugin.DefaultCDN}
}

type CheckCDN interface {
	HasCDN(target string) bool
}

type cdnCheck struct {
	plugins []pkg.PluginCDN
}

func (c cdnCheck) HasCDN(target string) bool {
	var ip []string
	var re = regexp.MustCompile(`^(?:25[0-5]|2[0-4]\d|[0-1]?\d{1,2})(?:\.(?:25[0-5]|2[0-4]\d|[0-1]?\d{1,2})){3}$`)
	if re.MatchString(target) {
		ip = append(ip, target)
	} else {
		q := dns.NewDNSQuery()
		ips, err := q.QueryA(target)
		if err != nil {
			fmt.Println(err)
			return false
		}
		ip = append(ip, ips...)
	}

	for _, v := range c.plugins {
		for _, i := range ip {
			lookup, err := v.Lookup(i)
			if err != nil {
				fmt.Println(err)
				continue
			}
			if lookup {
				return lookup
			}
		}

	}
	return false
}
