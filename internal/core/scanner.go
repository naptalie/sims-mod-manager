package core

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	
	"github.com/naptalie/sims4-mod-manager/internal/models"
)

// ScanModsFolder searches the mods folder and returns a list of mods
func ScanModsFolder(modsPath string) ([]models.Mod, error) {
	var mods []models.Mod
	
	// Check if mods folder exists
	if _, err := os.Stat(modsPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("mods folder not found at %s", modsPath)
	}
	
	// Walk through the mods folder
	err := filepath.Walk(modsPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		
		// Skip the root directory
		if path == modsPath {
			return nil
		}
		
		// Skip Resource.cfg and other configuration files
		if info.Name() == "Resource.cfg" {
			return nil
		}
		
		// Only include folders and mod files (.package, .ts4script)
		ext := strings.ToLower(filepath.Ext(path))
		if info.IsDir() || ext == ".package" || ext == ".ts4script" {
			// Get relative path from mods folder
			relPath, err := filepath.Rel(modsPath, path)
			if err != nil {
				relPath = path
			}
			
			// Extract version information (if available)
			// This is a simple implementation - in real use, you'd want to parse
			// the mod files to extract their actual versions
			version := "Unknown"
			
			// Try to extract version from filename patterns like "ModName_v1.2.3"
			parts := strings.Split(info.Name(), "_")
			for _, part := range parts {
				if strings.HasPrefix(part, "v") && len(part) > 1 {
					// Something like v1.2.3
					version = part[1:] // Remove the 'v'
					break
				}
			}
			
			mod := models.Mod{
				Name:    info.Name(),
				Path:    relPath,
				Version: version,
				Size:    info.Size(),
				Updated: info.ModTime(),
			}
			
			mods = append(mods, mod)
		}
		
		return nil
	})
	
	if err != nil {
		return nil, err
	}
	
	return mods, nil
}

// GetModVersions returns a map of mod names to their available versions
func GetModVersions(backupPath string) (map[string][]string, error) {
	result := make(map[string][]string)
	
	if _, err := os.Stat(backupPath); os.IsNotExist(err) {
		return result, nil
	}
	
	entries, err := os.ReadDir(backupPath)
	if err != nil {
		return nil, err
	}
	
	for _, entry := range entries {
		if entry.IsDir() {
			modName := entry.Name()
			modPath := filepath.Join(backupPath, modName)
			
			// Get versions
			versionEntries, err := os.ReadDir(modPath)
			if err != nil {
				continue
			}
			
			var versions []string
			for _, versionEntry := range versionEntries {
				if versionEntry.IsDir() {
					versions = append(versions, versionEntry.Name())
				}
			}
			
			if len(versions) > 0 {
				result[modName] = versions
			}
		}
	}
	
	return result, nil
}