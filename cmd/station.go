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
	"strings"
	"text/tabwriter"
	"time"

	"github.com/morgulbrut/transportCli/webreq/parsejson"

	"github.com/morgulbrut/transportCli/webreq"
	"github.com/spf13/cobra"
)

// stationCmd represents the station command
var stationCmd = &cobra.Command{
	Use:   "station",
	Short: "Returns the next connections leaving from a specific station.",
	Long: `Returns the next connections leaving from a specific station.
	
	Example: transportCli station Bern

	`,
	Run: func(cmd *cobra.Command, args []string) {
		var params strings.Builder
		if len(args) > 0 {
			params.WriteString("?station=" + strings.Join(args, "%20"))
		}

		lim, _ := cmd.Flags().GetString("limit")
		if lim != "" {
			params.WriteString("&limit=" + lim)
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

		printOut(webreq.Station(params.String()))
	},
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
}

func printOut(resp parsejson.RespStation) {
	const padding = 3
	w := tabwriter.NewWriter(os.Stdout, 0, 0, padding, ' ', tabwriter.Debug)
	fmt.Printf("\nStationtable for %s\n\n", resp.Station.Name)
	fmt.Fprintln(w, "Time \t Destination \t Platform \t Number")
	fmt.Fprintln(w, " \t \t \t ")

	for _, ele := range resp.Stationboard {
		tfs := "2006-01-02T15:04:05-0700"
		t, _ := time.Parse(tfs, ele.PassList[0].Departure)
		output := fmt.Sprintf("%02d:%02d\t %s \t %s \t %s %s", t.Hour(), t.Minute(), ele.To, ele.PassList[0].Platform, ele.Category, ele.Number)

		fmt.Fprintln(w, output)
	}
	w.Flush()
	fmt.Println()
}
