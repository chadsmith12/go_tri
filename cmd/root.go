/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/adrg/xdg"
	"github.com/chadsmith12/go_tri/path"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go_tri",
	Short: "go_try is a cli todo application written in Golang",
	Long: `go_tri will help you get more done is less time. It's designed to be as simple as possible to help you accomplish your goals.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	cobra.OnInitialize(initConfig)
	filePath, err := path.DataFilePath()
	if (err != nil) {
		fmt.Println("Unable to detect the default directory to use. Please set data file using --datafile")
		os.Exit(1)
	}

	path.DataFile = filePath
	rootCmd.PersistentFlags().StringVar(&path.ConfigFile, "config", "", "Config file to use")
	rootCmd.PersistentFlags().StringVar(&path.DataFile, "datafile", filePath, "Data file to store todos")
	viper.BindPFlag("datafile", rootCmd.PersistentFlags().Lookup("datafile"))
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig() {
	if path.ConfigFile != "" {
		viper.SetConfigFile(path.ConfigFile)	
	} else {
		configPath := filepath.Join(xdg.ConfigHome, "go_tri/")
		os.MkdirAll(configPath, 0700)
		viper.AddConfigPath(configPath)
		viper.SetConfigName(".config")
		viper.SetConfigType("ymal")
	}

	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file: ", viper.ConfigFileUsed())
	}
}
