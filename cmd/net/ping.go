/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package net

import (
	"log"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

var (
	urlPath []string
	value   []string
	client  = http.Client{
		Timeout: 2 * time.Second,
	}
)

func ping(domain string) (int, error) {
	url := "http://" + domain
	req, err := http.NewRequest("HEAD", url, nil)
	if err != nil {
		return 0, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	resp.Body.Close()
	return resp.StatusCode, nil
}

// pingCmd represents the ping command
var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "This pings is a remote URL and returns the response",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		// Logic
		if resp, err := ping(urlPath[0]); err != nil {
			log.Println(err)
		} else {
			log.Println(resp)
		}
	},
}

func init() {

	pingCmd.Flags().StringArrayVarP(&urlPath, "url", "u", value, "The url to ping")

	if err := pingCmd.MarkFlagRequired("url"); err != nil {
		log.Println(err)
	}

	// Here you will define your flags and configuration settings.
	NetCmd.AddCommand(pingCmd)
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
