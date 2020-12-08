/*
Copyright © 2019 morgulbrut
This work is free. You can redistribute it and/or modify it under the
terms of the Do What The Fuck You Want To Public License, Version 2,
as published by Sam Hocevar. See the LICENSE file or
 http://www.wtfpl.net/ for more details.
*/

package cmd

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/jedib0t/go-pretty/table"
	"github.com/morgulbrut/transportCli/webreq"
	"github.com/morgulbrut/transportCli/webreq/parsejson"

	"github.com/spf13/cobra"
)

// stationCmd represents the station command
var stationCmd = &cobra.Command{
	Use:   "station",
	Short: "Returns the next connections leaving from a specific station.",
	Long: `
    _____                                          _____     _______________ ________
    __  /________________ ________________ __________  /_    __  ____/___  / ____  _/
    _  __/__  ___/__  __ \__  ___/___  __ \__  ___/_  __/    _  /     __  /   __  /  
    / /_  _  /    _  / / /_(__  ) __  /_/ /_  /    / /_      / /___   _  /_____/ /   
    \__/  /_/     /_/ /_/ /____/  _  .___/ /_/     \__/      \____/   /_____//___/   
                                  /_/                                               
Returns the next connections leaving from a specific station.
	
	Example: 	transportCli station Bad Ragaz
				transportCli station "Bad Ragaz"

	`,
	Run: func(cmd *cobra.Command, args []string) {
		var params strings.Builder
		if len(args) > 0 {
			params.WriteString("?station=" + strings.Join(args, " "))
		}

		lim, _ := cmd.Flags().GetString("limit")
		if lim != "" {
			params.WriteString("&limit=" + lim)
		} else { // default
			params.WriteString("&limit=10")
		}

		b, _ := cmd.Flags().GetBool("bus")
		if b {
			params.WriteString("&transportations[]=bus")
		}
		b, _ = cmd.Flags().GetBool("tram")
		if b {
			params.WriteString("&transportations[]=tram")
		}
		b, _ = cmd.Flags().GetBool("train")
		if b {
			params.WriteString("&transportations[]=train")
		}
		b, _ = cmd.Flags().GetBool("ship")
		if b {
			params.WriteString("&transportations[]=ship")
		}
		b, _ = cmd.Flags().GetBool("cablecar")
		if b {
			params.WriteString("&transportations[]=cablecar")
		}
		// b, _ = cmd.Flags().GetBool("arrival")
		// if b {
		// 	params.WriteString("&type=arrival")
		// }

		PrintStation(webreq.Station(params.String()))
	},
}

func PrintStation(resp parsejson.RespStation) {

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	if runtime.GOOS == "windows" {
		t.SetStyle(table.StyleDouble)
	} else {
		t.SetStyle(table.StyleColoredBlackOnRedWhite)
	}
	t.AppendHeader(table.Row{"Time", "Destination", "Platform", "Category ", "Number"})

	for _, ele := range resp.Stationboard {
		tfs := "2006-01-02T15:04:05-0700"
		tm, _ := time.Parse(tfs, ele.PassList[0].Departure)
		tms := fmt.Sprintf("%02d:%02d", tm.Hour(), tm.Minute())
		t.AppendRow(table.Row{tms, ele.To, ele.PassList[0].Platform, ele.Category, ele.Number})
	}
	t.Render()
}

func init() {
	rootCmd.AddCommand(stationCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// stationCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// stationCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	stationCmd.Flags().StringP("limit", "l", "", "Number of departing connections to return.")
	stationCmd.Flags().Bool("train", false, "Include trains")
	stationCmd.Flags().Bool("tram", false, "Include trams")
	stationCmd.Flags().Bool("bus", false, "Include buses")
	stationCmd.Flags().Bool("ship", false, "Include ships")
	stationCmd.Flags().Bool("cablecar", false, "Include cablecar")
	//stationCmd.Flags().Bool("arrival", false, "Show arrival table")
}
