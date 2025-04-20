package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/naptalie/sims4-mod-manager/internal/config"
	"github.com/naptalie/sims4-mod-manager/internal/ui/styles"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure the Sims 4 Mod Manager",
	Run: func(cmd *cobra.Command, args []string) {
		configureApp()
	},
}

// Configure the application
func configureApp() {
	// Display current configuration
	fmt.Println(styles.TitleStyle.Render("⚙️ Current Configuration:"))
	fmt.Printf("Sims 4 Mods Path: %s\n", config.AppConfig.SimsModsPath)
	fmt.Printf("Mod Storage Path: %s\n", config.AppConfig.ModStoragePath)
	fmt.Println()
	
	// Ask for new configuration
	var newModsPath string
	var newStoragePath string
	
	fmt.Print("Enter new Sims 4 Mods Path (leave empty to keep current): ")
	fmt.Scanln(&newModsPath)
	
	fmt.Print("Enter new Mod Storage Path (leave empty to keep current): ")
	fmt.Scanln(&newStoragePath)
	
	// Update configuration if new values provided
	if newModsPath != "" {
		config.AppConfig.SimsModsPath = newModsPath
	}
	
	if newStoragePath != "" {
		config.AppConfig.ModStoragePath = newStoragePath
	}
	
	// Save configuration
	if err := config.UpdateConfig(); err != nil {
		fmt.Printf("Error saving configuration: %v\n", err)
		return
	}
	
	fmt.Println(styles.HighlightStyle.Render("✅ Configuration updated successfully!"))
}