package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sims4-mod-manager",
	Short: "A mod manager for The Sims 4",
	Long:  `Sims 4 Mod Manager is a CLI tool to help manage and version control your Sims 4 mods.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// Add all subcommands
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(versionsCmd)
	rootCmd.AddCommand(backupCmd)
	rootCmd.AddCommand(restoreCmd)
	rootCmd.AddCommand(configCmd)
	rootCmd.AddCommand(uiCmd)
}