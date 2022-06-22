/*
 *
 * ip.go
 * ip
 *
 * Created by lintao on 2022/6/17 10:50 AM
 * Copyright © 2020-2022 LINTAO. All rights reserved.
 *
 */

package ip

import (
	"github.com/NSObjects/scout/pkg"
	"github.com/yl2chen/cidranger"
	"net"
)

type ipRange struct {
	ranger cidranger.Ranger
}

func (r ipRange) Lookup(ip string) (bool, error) {
	return r.ranger.Contains(net.ParseIP(ip))
}

func NewIPRange() pkg.PluginCDN {
	ranger := cidranger.NewPCTrieRanger()
	for _, v := range cdns {
		_, network, err := net.ParseCIDR(v)
		if err != nil {
			panic(err)
		}

		if err = ranger.Insert(cidranger.NewBasicRangerEntry(*network)); err != nil {
			panic(err)
		}
	}
	return ipRange{ranger: ranger}
}
