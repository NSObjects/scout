/*
 *
 * interface.go
 * internal
 *
 * Created by lintao on 2022/6/22 11:32 AM
 * Copyright © 2020-2022 LINTAO. All rights reserved.
 *
 */

package pkg

import (
	"github.com/NSObjects/scout/internal/finger"
)

type PluginCDN interface {
	Lookup(ip string) (bool, error)
}

type PluginFinger interface {
	Scan(url string) (finger.ScanResult, error)
}

type PluginWaf interface {
	WafCheck(url string) (bool, error)
}
