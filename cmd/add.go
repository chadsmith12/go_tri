/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/chadsmith12/go_tri/todo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new todo item",
	Long: `Add will create a new todo item and add it onto the list`,
	Run: addRun,
}

var priority int

func init() {
	rootCmd.AddCommand(addCmd)
	
	addCmd.Flags().IntVarP(&priority, "priority", "p", 2, "Priority: 1, 2, 3")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

func addRun(cmd *cobra.Command, args []string) {
	items, readErr := todo.ReadItems(viper.GetString("datafile"))
	if readErr != nil {
		log.Fatal(fmt.Errorf("%v", readErr))
	}
	for _, x := range args {
		item := todo.Item { Text: x }
		item.SetPriority(priority)
		items = append(items, item)
	}

	saveErr := todo.SaveItems(viper.GetString("datafile"), items)
	if saveErr != nil {
		log.Fatal(fmt.Errorf("%v", saveErr))
	}
}
