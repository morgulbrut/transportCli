/*
Copyright © 2019 morgulbrut
This work is free. You can redistribute it and/or modify it under the
terms of the Do What The Fuck You Want To Public License, Version 2,
as published by Sam Hocevar. See the LICENSE file or
 http://www.wtfpl.net/ for more details.
*/

package cmd

import (
	"strings"

	"github.com/morgulbrut/transportCli/webreq"
	"github.com/spf13/cobra"
)

// connectionCmd represents the connection command
var connectionCmd = &cobra.Command{
	Use:   "connection",
	Short: "Returns connections between two station",
	Long: `
    _____                                          _____     _______________ ________
    __  /________________ ________________ __________  /_    __  ____/___  / ____  _/
    _  __/__  ___/__  __ \__  ___/___  __ \__  ___/_  __/    _  /     __  /   __  /  
    / /_  _  /    _  / / /_(__  ) __  /_/ /_  /    / /_      / /___   _  /_____/ /   
    \__/  /_/     /_/ /_/ /____/  _  .___/ /_/     \__/      \____/   /_____//___/   
                                  /_/                                          
Returns connections between two station, needs two stations.

Stationnames longer than one word must be written in quotation marks. 

	Example: 	transportCli station "Bad Ragaz" Zürich
	`,
	Run: func(cmd *cobra.Command, args []string) {
		var params strings.Builder
		if len(args) == 2 {
			params.WriteString("?from=" + args[0] + "&to=" + args[1])
		} else if len(args) == 1 {
			params.WriteString("?station=" + args[0])
		} else {
			cmd.Help()
		}

		lim, _ := cmd.Flags().GetString("limit")
		if lim != "" {
			params.WriteString("&limit=" + lim)
		} else { // default
			params.WriteString("&limit=1")
		}

		if len(args) == 2 {
			PrintConnection(webreq.Connection(params.String()))
		} else if len(args) == 1 {
			PrintStation(webreq.Station(params.String()))
		}
	},
}

func init() {
	rootCmd.AddCommand(connectionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// connectionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// connectionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	connectionCmd.Flags().StringP("limit", "l", "", "Number of departing connections to return.")

}
