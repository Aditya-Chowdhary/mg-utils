/*
Copyright © 2024 Aditya Chowdhary
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// swapCmd represents the swap command
var swapCmd = &cobra.Command{
	Use:   "swap",
	Short: "Swap two migration files",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: swapCommand,
}

func init() {
	rootCmd.AddCommand(swapCmd)
}

func swapCommand(cmd *cobra.Command, args []string) {
	in, err := cmd.Flags().GetBool("integer")
	if err != nil {
		fmt.Fprintln(cmd.ErrOrStderr(), "Error while parsing flag")
		return
	}

	if !checkSwapArguments(in, args) {
		fmt.Fprintln(cmd.ErrOrStderr(), "Incorrect Arguments")
		return
	}

	fmt.Fprintln(cmd.ErrOrStderr(), "Success")
}

func checkSwapArguments(in bool, args []string) bool {
	if in {
		if !checkInt(args[0]) || !checkInt(args[1]) {
			return false
		}
	}

	// TODO: Check if the dir and files exist

	return true
}
