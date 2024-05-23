/*
Copyright © 2024 Aditya Chowdhary
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "mg-util",
	Short: "A CLI app to provide common utilities for migration files",
	Long: `A CLI app to provide common utilities for migration files, such as move and swap. This will change the positions of both up and down files, even though you only specify one.
	- move -i 5 2 -> Move a migration file at position 5 to position 2, increasing the position of all the files in between
	- swap 000003_migrate.up.sql 000010_migrate.up.sql -> Swap the migration files at position 3 and 10.
	`,
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
	rootCmd.PersistentFlags().BoolP("integer", "i", false, "Use integers instead of file names")
	rootCmd.PersistentFlags().StringP("directory", "d", "", "Directory of migration files")
	cobra.MarkFlagRequired(rootCmd.Flags(), "directory")
}
