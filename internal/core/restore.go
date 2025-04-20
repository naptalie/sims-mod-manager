package core

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
	
	"github.com/naptalie/sims4-mod-manager/pkg/fsutil"
)

// RestoreMod restores a mod to a specific version
func RestoreMod(modName, version, modsPath, storagePath string) error {
	// Check if backup exists
	backupPath := filepath.Join(storagePath, modName, version)
	if _, err := os.Stat(backupPath); os.IsNotExist(err) {
		return fmt.Errorf("backup for mod %s version %s not found", modName, version)
	}
	
	// Target path in mods folder
	targetPath := filepath.Join(modsPath, modName)
	
	// If the mod exists in the mods folder, create a backup before replacing
	if _, err := os.Stat(targetPath); err == nil {
		// Create timestamp for automatic backup
		timestamp := time.Now().Format("20060102_150405") + "_prerestore"
		autoBackupDir := filepath.Join(storagePath, modName, timestamp)
		
		// Create auto-backup directory
		if err := os.MkdirAll(autoBackupDir, 0755); err != nil {
			return fmt.Errorf("could not create automatic backup: %v", err)
		}
		
		// Copy existing mod to auto-backup
		if err := fsutil.CopyDir(targetPath, autoBackupDir); err != nil {
			return fmt.Errorf("could not create automatic backup: %v", err)
		}
		
		// Remove the existing mod
		if err := os.RemoveAll(targetPath); err != nil {
			return fmt.Errorf("could not remove existing mod: %v", err)
		}
	}
	
	// Create the parent directory if it doesn't exist
	if err := os.MkdirAll(filepath.Dir(targetPath), 0755); err != nil {
		return fmt.Errorf("could not create mod directory: %v", err)
	}
	
	// Copy the backup to the mods folder
	if err := fsutil.CopyDir(backupPath, targetPath); err != nil {
		return fmt.Errorf("error restoring mod: %v", err)
	}
	
	return nil
}

// GetAvailableVersions returns a list of available versions for a mod
func GetAvailableVersions(modName, storagePath string) ([]string, error) {
	modPath := filepath.Join(storagePath, modName)
	
	if _, err := os.Stat(modPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("no backups found for mod %s", modName)
	}
	
	entries, err := os.ReadDir(modPath)
	if err != nil {
		return nil, fmt.Errorf("error reading backup folder: %v", err)
	}
	
	var versions []string
	for _, entry := range entries {
		if entry.IsDir() {
			versions = append(versions, entry.Name())
		}
	}
	
	if len(versions) == 0 {
		return nil, fmt.Errorf("no backups found for mod %s", modName)
	}
	
	return versions, nil
}