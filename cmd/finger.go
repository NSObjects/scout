/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/NSObjects/scout/pkg/finger"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
	"strings"
)

var urls string
var proxy string
var c int

// fingerCmd represents the finger command
var fingerCmd = &cobra.Command{
	Use:   "finger",
	Short: "web 指纹识别",
	Long:  `web 指纹识别: scout finger -u www.baidu.com,xxx.com -p https://proxy.com -c 20`,
	Run: func(cmd *cobra.Command, args []string) {
		newFinger := finger.NewFinger(strings.Split(urls, ","), proxy, c)
		scan, err := newFinger.Scan()
		if err != nil {
			fmt.Println(err)
			return
		}

		for _, s := range scan {
			outStr := fmt.Sprintf("[ %s | %s | %s | %d | %d | %s ]", s.Url, s.Cms, s.Server, s.StatusCode, s.Length, s.Title)
			color.RGBStyleFromString("237,64,35").Println(outStr)
		}
	},
}

func init() {
	rootCmd.AddCommand(fingerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	fingerCmd.Flags().StringVarP(&urls, "url", "u", "", "urls to scan, -u www.baidu.com,xxx.com")
	fingerCmd.Flags().StringVarP(&proxy, "proxy", "p", "", "proxy url, -p http://proxy.com")
	fingerCmd.Flags().IntVarP(&c, "count", "c", 5, "concurrency to scan, -c 20")
	rootCmd.MarkFlagRequired("url")
}
