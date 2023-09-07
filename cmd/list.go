/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"sort"
	"text/tabwriter"

	"github.com/chadsmith12/go_tri/path"
	"github.com/chadsmith12/go_tri/todo"
	"github.com/spf13/cobra"
)
// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all your todos",
	Long: `List all of your todos. You can pass in a '--done' flag to show all the items that are also done.`,
	Run: listRun,
}

var (
	doneOpt bool;
	allOpt bool
)

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")
	listCmd.Flags().BoolVar(&doneOpt, "done", false, "Show 'Done' Todos")
	listCmd.Flags().BoolVar(&allOpt, "all", false, "Show all Todos")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func listRun(cmd *cobra.Command, args []string) {
	items, err := todo.ReadItems(path.DataFile)
	if (err != nil){
		log.Fatal(fmt.Errorf("%v", err))
	}

	writer := tabwriter.NewWriter(os.Stdout, 3, 0, 1, ' ', 0)

	sort.Sort(todo.ByPriority(items))
	for _, item := range items {
		priorityString := item.PrettyP()
		label := item.Label()
		if allOpt || item.Done == doneOpt {
			doneString := item.PrettyDone()
			itemString := label + "\t" + doneString + "\t" + priorityString + "\t" + item.Text + "\t\n"
			fmt.Fprintf(writer, itemString)
		}	
	}
	writer.Flush()
}
