package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/naptalie/sims4-mod-manager/internal/config"
	"github.com/naptalie/sims4-mod-manager/internal/core"
	"github.com/naptalie/sims4-mod-manager/internal/ui/styles"
)

var backupCmd = &cobra.Command{
	Use:   "backup [mod name]",
	Short: "Backup a specific mod or all mods",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			backupMod(args[0])
		} else {
			backupAllMods()
		}
	},
}

// Backup a specific mod
func backupMod(modName string) {
	fmt.Printf(styles.TitleStyle.Render("🔄 Backing up mod: %s\n"), modName)
	
	err := core.BackupMod(modName, config.AppConfig.SimsModsPath, config.AppConfig.ModStoragePath)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	
	fmt.Println(styles.HighlightStyle.Render("✅ Backup completed successfully!"))
}

// Backup all mods
func backupAllMods() {
	fmt.Println(styles.TitleStyle.Render("🔄 Backing up all mods"))
	
	successCount, failedMods, err := core.BackupAllMods(
		config.AppConfig.SimsModsPath, 
		config.AppConfig.ModStoragePath,
	)
	
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	
	fmt.Printf("Successfully backed up %d mods\n", successCount)
	
	if len(failedMods) > 0 {
		fmt.Println("Failed to backup the following mods:")
		for _, mod := range failedMods {
			fmt.Printf("- %s\n", mod)
		}
	}
	
	fmt.Println(styles.HighlightStyle.Render("✅ Backup operation completed!"))
}