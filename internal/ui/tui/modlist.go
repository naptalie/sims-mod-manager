package tui

import (
	"fmt"
	"strings"
	
	tea "github.com/charmbracelet/bubbletea"

	"github.com/naptalie/sims4-mod-manager/internal/config"
	"github.com/naptalie/sims4-mod-manager/internal/core"
	"github.com/naptalie/sims4-mod-manager/internal/models"
)

// ModItem represents an item in the mod list
type ModItem struct {
	mod models.Mod
}

// NewModItem creates a new mod item
func NewModItem(mod models.Mod) ModItem {
	return ModItem{mod: mod}
}

// Title returns the title of the mod
func (i ModItem) Title() string {
	return i.mod.Name
}

// Description returns additional information about the mod
func (i ModItem) Description() string {
	size := "unknown size"
	if i.mod.Size > 0 {
		size = fmt.Sprintf("%.2f MB", float64(i.mod.Size)/(1024*1024))
	}
	
	updated := "unknown date"
	if !i.mod.Updated.IsZero() {
		updated = i.mod.Updated.Format("2006-01-02")
	}
	
	return fmt.Sprintf("Size: %s, Updated: %s", size, updated)
}

// FilterValue returns the value used for filtering
func (i ModItem) FilterValue() string {
	return i.mod.Name
}

// loadMods loads the mods from the filesystem
func loadMods() tea.Cmd {
	return func() tea.Msg {
		mods, err := core.ScanModsFolder(config.AppConfig.SimsModsPath)
		return modsLoadedMsg{mods: mods, err: err}
	}
}

// filterMods filters the mods based on the search term
func filterMods(mods []models.Mod, searchTerm string) []models.Mod {
	if searchTerm == "" {
		return mods
	}
	
	var filtered []models.Mod
	for _, mod := range mods {
		if strings.Contains(strings.ToLower(mod.Name), strings.ToLower(searchTerm)) {
			filtered = append(filtered, mod)
		}
	}
	
	return filtered
}