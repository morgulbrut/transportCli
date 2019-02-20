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

	"github.com/jedib0t/go-pretty/table"
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

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	if runtime.GOOS == "windows" {
		t.SetStyle(table.StyleDouble)
	} else {
		t.SetStyle(table.StyleColoredDark)
	}
	t.AppendHeader(table.Row{"Name", "Coordinates", "Distance"})

	for _, ele := range resp.Stations {
		coords := fmt.Sprintf(" %f %f", ele.Coordinates.X, ele.Coordinates.Y)
		t.AppendRow(table.Row{ele.Name, coords, ele.Distance})
	}
	t.Render()
}
