/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
	"github.com/spf13/cobra"
	"github.com/takekou0130/meta-curl/domain"
	"github.com/takekou0130/meta-curl/view"

	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "meta-curl",
	Short: "this is my application",
	Long: `
	meta-curl is a command for fetch page meta infomation.
	you can use this command for SEO.
	`,
	Args: cobra.ExactArgs(1),
	Run:  fetch,
}

func fetch(cmd *cobra.Command, args []string) {
	url := args[0]
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("User-Agent", "")

	client := new(http.Client)

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	metaInfo := doc2metaInfo(url, doc)

	var t view.View
	t = view.NewTableRenderer()
	err = t.Render(*metaInfo)
	if err != nil {
		log.Fatal(err)
	}
}

func doc2metaInfo(url string, doc *goquery.Document) *domain.MetaInfo {
	title := doc.Find("title").Text()
	// TODO for文でとってくる
	desc, _ := doc.Find("meta[name='description']").Attr("content")
	key, _ := doc.Find("meta[name='keywords']").Attr("content")
	cano, _ := doc.Find("link[rel='canonical']").Attr("href")
	alt, _ := doc.Find("link[rel='alternate']").Attr("href")

	var m *domain.MetaInfo
	m.Url = url
	m.Title = append(m.Title, title)
	m.Description = append(m.Description, desc)
	m.Keywords = append(m.Keywords, key)
	m.Canonical = append(m.Canonical, cano)
	m.Alternate = append(m.Alternate, alt)
	return m
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.meta-curl.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".meta-curl" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".meta-curl")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
