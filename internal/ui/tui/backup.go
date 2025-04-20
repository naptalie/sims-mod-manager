package tui

import (
	"fmt"
	
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	
	"github.com/yourusername/sims4-mod-manager/internal/config"
	"github.com/yourusername/sims4-mod-manager/internal/core"
	"github.com/yourusername/sims4-mod-manager/internal/ui/styles"
)

// Custom message types for backup operations
type backupStartMsg struct {
	modName string
}

type backupFinishedMsg struct {
	success bool
	message string
}

// Function to start a backup operation
func startBackup(modName string) tea.Cmd {
	return func() tea.Msg {
		return backupStartMsg{modName: modName}
	}
}

// Function to perform the backup in the background
func performBackup(modName string) tea.Cmd {
	return func() tea.Msg {
		var err error
		
		if modName == "" {
			// Backup all mods
			_, failedMods, err := core.BackupAllMods(
				config.AppConfig.SimsModsPath,
				config.AppConfig.ModStoragePath,
			)
			
			if err != nil {
				return backupFinishedMsg{
					success: false,
					message: fmt.Sprintf("Error: %v", err),
				}
			}
			
			if len(failedMods) > 0 {
				return backupFinishedMsg{
					success: true,
					message: fmt.Sprintf("Backup completed with some failures (%d mods failed)", len(failedMods)),
				}
			}
			
			return backupFinishedMsg{
				success: true,
				message: "All mods backed up successfully!",
			}
		} else {
			// Backup a specific mod
			err = core.BackupMod(
				modName,
				config.AppConfig.SimsModsPath,
				config.AppConfig.ModStoragePath,
			)
			
			if err != nil {
				return backupFinishedMsg{
					success: false,
					message: fmt.Sprintf("Error: %v", err),
				}
			}
			
			return backupFinishedMsg{
				success: true,
				message: fmt.Sprintf("Mod '%s' backed up successfully!", modName),
			}
		}
	}
}

// BackupModel represents the state for backup operations
type BackupModel struct {
	spinner   spinner.Model
	modName   string
	status    string
	completed bool
	success   bool
}

// Initialize a new backup model
func NewBackupModel(modName string) BackupModel {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(styles.PlumbobColor)
	
	status := "Backing up all mods..."
	if modName != "" {
		status = fmt.Sprintf("Backing up mod '%s'...", modName)
	}
	
	return BackupModel{
		spinner:   s,
		modName:   modName,
		status:    status,
		completed: false,
		success:   false,
	}
}

// Init initializes the backup model
func (m BackupModel) Init() tea.Cmd {
	return tea.Batch(
		m.spinner.Tick,
		performBackup(m.modName),
	)
}

// Update handles messages for the backup model
func (m BackupModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			// Return to main view
			mainModel, cmd := NewMainModel()
			return mainModel, cmd
		}
		
	case backupFinishedMsg:
		m.completed = true
		m.success = msg.success
		m.status = msg.message
		return m, nil
		
	default:
		var spinnerCmd tea.Cmd
		m.spinner, spinnerCmd = m.spinner.Update(msg)
		return m, spinnerCmd
	}

	return m, cmd
}

// View renders the backup model
func (m BackupModel) View() string {
	if m.completed {
		if m.success {
			return fmt.Sprintf("\n  ✅ %s\n\n  Press any key to return.\n", m.status)
		}
		return fmt.Sprintf("\n  ❌ %s\n\n  Press any key to return.\n", m.status)
	}
	
	return fmt.Sprintf("\n  %s %s\n\n", m.spinner.View(), m.status)
}