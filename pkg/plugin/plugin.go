/*
 *
 * plugin.go
 * plugin
 *
 * Created by lintao on 2022/6/22 11:23 AM
 * Copyright © 2020-2022 LINTAO. All rights reserved.
 *
 */

package plugin

import (
	"github.com/NSObjects/scout/internal/cdn/geolite"
	"github.com/NSObjects/scout/internal/cdn/ip"
	"github.com/NSObjects/scout/internal/finger/ehole"
	"github.com/NSObjects/scout/pkg"
)

var DefaultCDN []pkg.PluginCDN
var DefaultFinger []pkg.PluginFinger

func init() {
	DefaultCDN = append(DefaultCDN, geolite.NewGeoLite())
	DefaultCDN = append(DefaultCDN, ip.NewIPRange())
	DefaultFinger = append(DefaultFinger, ehole.NewFingerScan(""))
}
