/*
 *
 * geo lite.go
 * cdn
 *
 * Created by lintao on 2022/6/14 4:04 PM
 * Copyright © 2020-2022 LINTAO. All rights reserved.
 *
 */

package geolite

import (
	"errors"
	"fmt"
	"github.com/NSObjects/scout/pkg"
	"github.com/cavaliergopher/grab/v3"
	"github.com/jasonlvhit/gocron"
	"github.com/oschwald/maxminddb-golang"
	"github.com/thoas/go-funk"
	"time"

	"net"
	"os"
)

type Record struct {
	Country struct {
		ISOCode string `maxminddb:"iso_code"`
	} `maxminddb:"country"`
	ASN int    `maxminddb:"autonomous_system_number"`
	ASO string `maxminddb:"autonomous_system_organization"`
}

type geoLite struct {
	path string
}

var path = "./GeoLite2-ASN.mmdb"
var geoURL = "https://raw.githubusercontent.com/P3TERX/GeoLite.mmdb/download/GeoLite2-ASN.mmdb"

func NewGeoLite() pkg.PluginCDN {
	g := &geoLite{}
	if _, err := os.Stat(path); err != nil {
		go g.GetLiteDownloadTask()
	} else {
		g.path = path
	}
	return g
}

func (g geoLite) GetLiteDownloadTask() {
	if err := githubDownload(geoURL, path); err != nil {
		fmt.Println(err)
	} else {
		g.path = path
	}

	if err := gocron.Every(1).Day().At("22:00").Do(func() {
		if time.Now().Weekday() != time.Tuesday {
			return
		}
		if err := githubDownload(geoURL, path); err != nil {
			fmt.Println(err)
		} else {
			g.path = path
		}
	}); err != nil {
		panic(err)
	}
	<-gocron.Start()
}

func githubDownload(url string, dst string) error {
	get, err := grab.Get(dst, fmt.Sprintf("https://git.lt7.top/%s", url))
	if err != nil {
		return err
	}
	fmt.Println("Download saved to", get.Filename)
	return nil
}

func (g *geoLite) Lookup(ip string) (bool, error) {
	if g.path == "" {
		return false, errors.New("GeoLite2-ASN.mmdb file not found")
	}

	db, err := maxminddb.Open(g.path)
	if err != nil {
		return false, err
	}
	defer db.Close()

	parseIP := net.ParseIP(ip)
	var record Record
	err = db.Lookup(parseIP, &record)
	if err != nil {
		return false, err
	}

	return funk.Contains(ASNS, record.ASN), nil
}
