/*
 *
 * finger.go
 * finger
 *
 * Created by lintao on 2022/6/22 4:22 PM
 * Copyright © 2020-2022 LINTAO. All rights reserved.
 *
 */

package finger

import (
	"fmt"
	"github.com/NSObjects/scout/internal/finger"
	"github.com/NSObjects/scout/pkg"
	"github.com/NSObjects/scout/pkg/plugin"
)

type Finger interface {
	Scan(url string) (finger.ScanResult, error)
}

type fingerCheck struct {
	plugins []pkg.PluginFinger
}

func NewFinger() Finger {
	return fingerCheck{plugins: plugin.DefaultFinger}
}

func (f fingerCheck) Scan(url string) (finger.ScanResult, error) {
	for _, v := range f.plugins {
		scan, err := v.Scan(url)
		if err != nil {
			fmt.Println(err)
			continue
		}
		return scan, nil

	}
	return finger.ScanResult{}, nil
}
