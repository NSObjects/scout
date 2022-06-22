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

type PluginCDN interface {
	Lookup(ip string) (bool, error)
}
