/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "scout",
	Short: "scout 是网站漏洞扫描工具，帮助您快速检测网站安全隐患",
	Long: `scout 是网站漏洞扫描工具，帮助您快速检测网站安全隐患,
功能: 
Waf识别	   waf -u www.baidu.com,xxx.com -p https://proxy.com -c 20
CDN检测     cdn -u www.baidu.com,xxx.com -p https://proxy.com -c 20
web指纹识别 finger -u www.baidu.com,xxx.com -p https://proxy.com -c 20`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.scout.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
