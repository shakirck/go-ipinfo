/*
Copyright © 2023 shakirck shakirckdeveloper@gmail.com
*/
package cmd

import (
	"encoding/json"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"io/ioutil"
	"net/http"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "go-ipinfo",
	Short: "Ip Info ",
	Long:  `go-ipinfo is a simple tool to track Ip informations`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {

			ip := args[0]
			prettifyOuput(ip)
		} else if len(args) == 0 {

			color.Red("No Ip Provided , Returning Your information")
			prettifyOuput("")
		} else {
			println("invalid command")

		}

	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

type IpData struct {
	Ip       string `json:"ip"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Loc      string `json:"loc"`
	Timezone string `json:"timezone"`
	Postal   string `json:"postal"`
}

func prettifyOuput(ip string) {

	resByte := getIpInfo(ip)
	data := IpData{}

	err := json.Unmarshal(resByte, &data)
	if err != nil {
		println(err)
		return
	}

	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		panic("issue wiithjson ")
	}

	println(string(jsonData))

}
func getIpInfo(ip string) []byte {

	url := "https://www.ipinfo.io/" + ip

	res, err := http.Get(url)

	if err != nil {
		panic("Invalid  IP or Network Error")

	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		println(err)
	}
	return data

}
