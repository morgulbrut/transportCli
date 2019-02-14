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
	"log"
	"os"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/morgulbrut/transportCli/webreq/parseJSON"

	"github.com/morgulbrut/transportCli/webreq"
	"github.com/spf13/cobra"
)

// stationCmd represents the station command
var stationCmd = &cobra.Command{
	Use:   "station",
	Short: "Returns the next connections leaving from a specific location.",
	Long: `Returns the next connections leaving from a specific location.
	
	Example: transportCli station Bern

	`,
	Run: func(cmd *cobra.Command, args []string) {
		loc := "?station=Bern"
		if len(args) > 0 {
			loc = "?station=" + strings.Join(args, "%20")
		}

		lim, _ := cmd.Flags().GetString("limit")
		if lim != "" {
			lim = "&limit=" + lim
		}
		printOut(webreq.WebreqStation(loc + lim))
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
}

func printOut(resp parseJSON.ResponseStation) {
	const padding = 3
	w := tabwriter.NewWriter(os.Stdout, 0, 0, padding, ' ', tabwriter.Debug)

	for _, ele := range resp.Stationboard {
		fmt.Println(ele.Stop.Departure)
		tfs := "2006-01-02T15:04:05-0700"
		t, err := time.Parse(tfs, ele.Stop.Departure)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(t)
		output := fmt.Sprintf("%v:%v\t %s \t %s", t.Hour(), t.Minute(), ele.To, ele.Name)

		fmt.Fprintln(w, output)
	}
	w.Flush()
}
