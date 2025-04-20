package config

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

// Config struct to hold application configuration
type Config struct {
	SimsModsPath   string
	ModStoragePath string
}

var AppConfig Config

// InitConfig reads in config file and ENV variables if set
func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	
	// Try to find user's Documents folder
	home, err := os.UserHomeDir()
	if err == nil {
		docsPath := filepath.Join(home, "Documents")
		// Default Sims 4 mods path for most installations
		defaultModsPath := filepath.Join(docsPath, "Electronic Arts", "The Sims 4", "Mods")
		defaultStoragePath := filepath.Join(home, ".sims4_mod_manager")
		
		viper.SetDefault("SimsModsPath", defaultModsPath)
		viper.SetDefault("ModStoragePath", defaultStoragePath)
	}
	
	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME/.sims4_mod_manager")
	
	// If a config file is found, read it in
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; create default config
			AppConfig.SimsModsPath = viper.GetString("SimsModsPath")
			AppConfig.ModStoragePath = viper.GetString("ModStoragePath")
			
			// Create mod storage directory if it doesn't exist
			if _, err := os.Stat(AppConfig.ModStoragePath); os.IsNotExist(err) {
				os.MkdirAll(AppConfig.ModStoragePath, 0755)
			}
			
			// Save default config
			viper.Set("SimsModsPath", AppConfig.SimsModsPath)
			viper.Set("ModStoragePath", AppConfig.ModStoragePath)
			viper.SafeWriteConfig()
		}
	}
	
	// Read config into our struct
	AppConfig.SimsModsPath = viper.GetString("SimsModsPath")
	AppConfig.ModStoragePath = viper.GetString("ModStoragePath")
}

// UpdateConfig saves the current configuration
func UpdateConfig() error {
	viper.Set("SimsModsPath", AppConfig.SimsModsPath)
	viper.Set("ModStoragePath", AppConfig.ModStoragePath)
	
	// Create mod storage directory if it doesn't exist
	if _, err := os.Stat(AppConfig.ModStoragePath); os.IsNotExist(err) {
		os.MkdirAll(AppConfig.ModStoragePath, 0755)
	}
	
	if err := viper.WriteConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found, create it
			return viper.SafeWriteConfig()
		}
		return err
	}
	
	return nil
}