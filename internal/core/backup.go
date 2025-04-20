package core

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
	
	"github.com/naptalie/sims4-mod-manager/pkg/fsutil"
)

// BackupMod backs up a specific mod
func BackupMod(modName, modsPath, storagePath string) error {
	// Find the mod in the mods folder
	modPath := filepath.Join(modsPath, modName)
	
	if _, err := os.Stat(modPath); os.IsNotExist(err) {
		return fmt.Errorf("mod %s not found in your Sims 4 mods folder", modName)
	}
	
	// Create timestamp for version
	timestamp := time.Now().Format("20060102_150405")
	
	// Create backup directory
	backupDir := filepath.Join(storagePath, modName, timestamp)
	if err := os.MkdirAll(backupDir, 0755); err != nil {
		return fmt.Errorf("error creating backup directory: %v", err)
	}
	
	// Copy the mod to the backup directory
	if err := fsutil.CopyDir(modPath, backupDir); err != nil {
		return fmt.Errorf("error backing up mod: %v", err)
	}
	
	return nil
}

// BackupAllMods backs up all mods
func BackupAllMods(modsPath, storagePath string) (int, []string, error) {
	mods, err := ScanModsFolder(modsPath)
	if err != nil {
		return 0, nil, fmt.Errorf("error scanning mods folder: %v", err)
	}
	
	if len(mods) == 0 {
		return 0, nil, fmt.Errorf("no mods found in your Sims 4 mods folder")
	}
	
	// Create timestamp for version
	timestamp := time.Now().Format("20060102_150405")
	
	successCount := 0
	var failedMods []string
	
	for _, mod := range mods {
		// Skip directories that are too deep in the hierarchy
		if filepath.Dir(mod.Path) != "." {
			continue
		}
		
		// Create backup directory
		backupDir := filepath.Join(storagePath, mod.Name, timestamp)
		if err := os.MkdirAll(backupDir, 0755); err != nil {
			failedMods = append(failedMods, mod.Name)
			continue
		}
		
		// Copy the mod to the backup directory
		modPath := filepath.Join(modsPath, mod.Path)
		destPath := filepath.Join(backupDir, mod.Name)
		
		if mod.Size == 0 && fsutil.FileExists(modPath) {
			// Directory mod
			if err := fsutil.CopyDir(modPath, destPath); err != nil {
				failedMods = append(failedMods, mod.Name)
				continue
			}
		} else {
			// File mod
			if err := fsutil.CopyFile(modPath, destPath); err != nil {
				failedMods = append(failedMods, mod.Name)
				continue
			}
		}
		
		successCount++
	}
	
	return successCount, failedMods, nil
}