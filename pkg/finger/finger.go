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
	"github.com/NSObjects/scout/internal/finger"
	"github.com/NSObjects/scout/internal/workpool"
	"github.com/NSObjects/scout/pkg"
	"github.com/NSObjects/scout/pkg/plugin"
)

type Finger interface {
	Scan() ([]finger.ScanResult, error)
}

type fingerCheck struct {
	plugins     []pkg.PluginFinger
	urls        []string
	proxy       string
	concurrency int
}

func NewFinger(urls []string, proxy string, c int) Finger {
	return fingerCheck{plugins: plugin.DefaultFinger, proxy: proxy, urls: urls, concurrency: c}
}

func (f fingerCheck) Scan() ([]finger.ScanResult, error) {
	wp := workpool.NewWorkerPool(f.concurrency)
	var results []finger.ScanResult

	for _, u := range f.urls {
		url := u
		wp.AddTask(func() {
			for _, plugin := range f.plugins {
				result, err := plugin.Scan(url)
				if err != nil {
					continue
				}
				results = append(results, result)
			}

		})
	}
	wp.Wait()

	return results, nil
}
