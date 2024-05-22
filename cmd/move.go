/*
Copyright © 2024 Aditya Chowdhary
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// moveCmd represents the move command
var moveCmd = &cobra.Command{
	Use:   "move",
	Short: "Move a migration file from one sequence number to another",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run:  moveCommand,
	Args: cobra.ExactArgs(2),
}

func init() {
	rootCmd.AddCommand(moveCmd)
}

func moveCommand(cmd *cobra.Command, args []string) {
	in, err := cmd.Flags().GetBool("integer")
	if err != nil {
		fmt.Fprintln(cmd.ErrOrStderr(), "Error while parsing flag")
		return
	}

	if !checkMoveArguments(in, args) {
		fmt.Fprintln(cmd.ErrOrStderr(), "Incorrect Arguments")
		return
	}

	fmt.Fprintln(cmd.ErrOrStderr(), "Success")
}

func checkMoveArguments(in bool, args []string) bool {
	if !checkInt(args[1]) {
		return false
	}

	if in {
		if !checkInt(args[0]) {
			return false
		}
	}

	// TODO: Check if the dir and files exist

	return true
}

func checkInt(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}
