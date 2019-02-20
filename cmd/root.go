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
	"time"

	"github.com/jedib0t/go-pretty/table"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/morgulbrut/transportCli/webreq/parsejson"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "transportCli",
	Short: "A simple CLI to the transport API.",
	Long: `
    _____                                          _____     _______________ ________
    __  /________________ ________________ __________  /_    __  ____/___  / ____  _/
    _  __/__  ___/__  __ \__  ___/___  __ \__  ___/_  __/    _  /     __  /   __  /  
    / /_  _  /    _  / / /_(__  ) __  /_/ /_  /    / /_      / /___   _  /_____/ /   
    \__/  /_/     /_/ /_/ /____/  _  .___/ /_/     \__/      \____/   /_____//___/   
                                  /_/                                                
A simple CLI to the transport API.

	Because why not.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.transportCli.yaml)")

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
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".transportCli" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".transportCli")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func PrintStation(resp parsejson.RespStation) {

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	if runtime.GOOS == "windows" {
		t.SetStyle(table.StyleDouble)
	} else {
		t.SetStyle(table.StyleColoredDark)
	}
	t.AppendHeader(table.Row{"Time", "Destination", "Platform", "Category ", "Number"})

	for _, ele := range resp.Stationboard {
		tfs := "2006-01-02T15:04:05-0700"
		tm, _ := time.Parse(tfs, ele.PassList[0].Departure)
		tms := fmt.Sprintf("%02d:%02d", tm.Hour(), tm.Minute())
		t.AppendRow(table.Row{tms, ele.To, ele.PassList[0].Platform, ele.Category, ele.Number})
	}
	t.Render()
	fmt.Println()
}

func PrintConnection(resp parsejson.RespConnection) {

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	if runtime.GOOS == "windows" {
		t.SetStyle(table.StyleDouble)
	} else {
		t.SetStyle(table.StyleColoredDark)
	}
	t.AppendHeader(table.Row{"Departure", "Time", "Platform", "Arrival", "Time", "Platform", "Duration", "Changes"})

	for _, ele := range resp.Connections {
		tfs := "2006-01-02T15:04:05-0700"
		td, _ := time.Parse(tfs, ele.From.Departure)
		ta, _ := time.Parse(tfs, ele.To.Arrival)
		dur := ta.Sub(td)
		tds := fmt.Sprintf("%02d:%02d", td.Hour(), td.Minute())
		tas := fmt.Sprintf("%02d:%02d", ta.Hour(), ta.Minute())
		durs := fmt.Sprintf("%02d:%02d", int(dur.Hours()), int(dur.Minutes())%60)
		//t.AppendRow(table.Row{tms, ele.To, ele.PassList[0].Platform, ele.Category, ele.Number})
		t.AppendRow(table.Row{ele.From.Station.Name, tds, ele.From.Platform, ele.To.Station.Name, tas, ele.To.Platform, durs, ele.Sections})
	}
	t.Render()
	fmt.Println()
}

func PrintLocation(resp parsejson.RespLocation) {

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
