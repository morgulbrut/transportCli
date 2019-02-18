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
	"text/tabwriter"

	"github.com/morgulbrut/transportCli/webreq/parsejson"
	"github.com/spf13/cobra"
)

// locationCmd represents the location command
var locationCmd = &cobra.Command{
	Use:   "location",
	Short: "Returns nearby stations.",
	Long:  `Returns nearby stations either trough a query or trough coordinates`,
}

func init() {
	rootCmd.AddCommand(locationCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// locationCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// locationCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func PrintLoc(resp parsejson.RespLocation) {
	const padding = 3
	w := tabwriter.NewWriter(os.Stdout, 0, 0, padding, ' ', tabwriter.Debug)
	fmt.Printf("\nNearby Stations\n\n")
	fmt.Fprintln(w, "Name \t Coordinates \tDistance")
	fmt.Fprintln(w, " \t \t ")

	for _, ele := range resp.Stations {
		output := fmt.Sprintf("%s \t %f %f  \t %d m", ele.Name, ele.Coordinates.X, ele.Coordinates.Y, ele.Distance)

		fmt.Fprintln(w, output)
	}
	w.Flush()
	fmt.Println()
}
