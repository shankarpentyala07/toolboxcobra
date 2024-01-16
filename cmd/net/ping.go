/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package net

import (
	"fmt"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

var (
	urlPath string

	client = http.Client{
		Timeout: time.Second*2,
	}
)

func ping(domain string) (int,error){
	url := "http://" + domain
	req, err := http.NewRequest("HEAD",url,nil)
	if err != nil{
		return 0,err
	}

	res, err := client.Do(req)
	if err != nil {
		return 0,err
	}
	res.Body.Close()
	return res.StatusCode,nil

}

// pingCmd represents the ping command
var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "This pings a remote URL and returns the response",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		
		//Logic 
		sc,err := ping(urlPath)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(sc)
	},
}



func init() {
	pingCmd.Flags().StringVarP(&urlPath, "url", "u","","The url to ping")

	// The url flag passing is mandatory
	if err := pingCmd.MarkFlagRequired("url"); err != nil {
		fmt.Println(err)
	}
	//rootCmd.AddCommand(pingCmd)
	NetCmd.AddCommand(pingCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
