package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/naptalie/sims4-mod-manager/internal/config"
	"github.com/naptalie/sims4-mod-manager/internal/core"
	"github.com/naptalie/sims4-mod-manager/internal/ui/styles"
)

var restoreCmd = &cobra.Command{
	Use:   "restore [mod name] [version]",
	Short: "Restore a mod to a specific version",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		restoreMod(args[0], args[1])
	},
}

// Restore a mod to a specific version
func restoreMod(modName string, version string) {
	fmt.Printf(styles.TitleStyle.Render("🔄 Restoring mod: %s to version %s\n"), modName, version)
	
	err := core.RestoreMod(
		modName, 
		version, 
		config.AppConfig.SimsModsPath, 
		config.AppConfig.ModStoragePath,
	)
	
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	
	fmt.Println(styles.HighlightStyle.Render("✅ Mod restored successfully!"))
}