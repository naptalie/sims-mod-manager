package tui

import (
	"fmt"
	
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	
	"github.com/naptalie/sims4-mod-manager/internal/config"
	"github.com/naptalie/sims4-mod-manager/internal/core"
	"github.com/naptalie/sims4-mod-manager/internal/ui/styles"
)

// Custom message types for restore operations
type restoreStartMsg struct {
	modName string
}

type restoreVersionsLoadedMsg struct {
	versions []string
	err      error
}

type restoreFinishedMsg struct {
	success bool
	message string
}

// Function to start a restore operation
func startRestore(modName string) tea.Cmd {
	return func() tea.Msg {
		return restoreStartMsg{modName: modName}
	}
}

// Function to load available versions for a mod
func loadVersions(modName string) tea.Cmd {
	return func() tea.Msg {
		versions, err := core.GetAvailableVersions(modName, config.AppConfig.ModStoragePath)
		return restoreVersionsLoadedMsg{versions: versions, err: err}
	}
}

// Function to perform the restore in the background
func performRestore(modName, version string) tea.Cmd {
	return func() tea.Msg {
		err := core.RestoreMod(
			modName,
			version,
			config.AppConfig.SimsModsPath,
			config.AppConfig.ModStoragePath,
		)
		
		if err != nil {
			return restoreFinishedMsg{
				success: false,
				message: fmt.Sprintf("Error: %v", err),
			}
		}
		
		return restoreFinishedMsg{
			success: true,
			message: fmt.Sprintf("Mod '%s' restored to version %s successfully!", modName, version),
		}
	}
}

// VersionItem represents an item in the version list
type VersionItem struct {
	value string
}

// NewVersionItem creates a new version item
func NewVersionItem(value string) VersionItem {
	return VersionItem{value: value}
}

// Title returns the title of the version
func (i VersionItem) Title() string {
	return i.value
}

// Description returns additional information about the version
func (i VersionItem) Description() string {
	return ""
}

// FilterValue returns the value used for filtering
func (i VersionItem) FilterValue() string {
	return i.value
}

// RestoreModel represents the state for restore operations
type RestoreModel struct {
	state         string
	modName       string
	versions      []string
	selectedVersion string
	spinner       spinner.Model
	list          list.Model
	errorMsg      string
	successMsg    string
}

// Initialize a new restore model
func NewRestoreModel(modName string) RestoreModel {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(styles.PlumbobColor)
	
	return RestoreModel{
		state:         "loading-versions",
		modName:       modName,
		versions:      []string{},
		selectedVersion: "",
		spinner:       s,
		errorMsg:      "",
		successMsg:    "",
	}
}

// Init initializes the restore model
func (m RestoreModel) Init() tea.Cmd {
	return tea.Batch(
		m.spinner.Tick,
		loadVersions(m.modName),
	)
}

// Update handles messages for the restore model
func (m RestoreModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "esc":
			// Return to main menu
			mainModel, cmd := NewMainModel()
			return mainModel, cmd
		case "enter":
			if m.state == "select-version" {
				// Get selected version
				selectedItem := m.list.SelectedItem()
				if selectedItem != nil {
					versionItem, ok := selectedItem.(VersionItem)
					if ok {
						m.selectedVersion = versionItem.value
						m.state = "restoring"
						return m, performRestore(m.modName, m.selectedVersion)
					}
				}
			}
		}
		
		
	case restoreVersionsLoadedMsg:
		if msg.err != nil {
			m.errorMsg = msg.err.Error()
			m.state = "error"
			return m, nil
		}
		
		m.versions = msg.versions
		
		// Initialize the list with versions
		var items []list.Item
		for _, version := range m.versions {
			items = append(items, NewVersionItem(version))
		}
		
		// Create a new list
		l := list.New(items, list.NewDefaultDelegate(), 0, 0)
		l.Title = fmt.Sprintf("Available Versions for '%s'", m.modName)
		l.SetShowStatusBar(false)
		l.SetFilteringEnabled(false)
		l.Styles.Title = styles.TitleStyle
		l.Styles.NoItems = styles.NormalTextStyle
		
		m.list = l
		m.state = "select-version"
		return m, nil
		
	case restoreFinishedMsg:
		if msg.success {
			m.successMsg = msg.message
			m.state = "success"
		} else {
			m.errorMsg = msg.message
			m.state = "error"
		}
		return m, nil
		
	default:
		// Handle spinner updates for loading states
		if m.state == "loading-versions" || m.state == "restoring" {
			var spinnerCmd tea.Cmd
			m.spinner, spinnerCmd = m.spinner.Update(msg)
			return m, spinnerCmd
		}
		
		// Handle list updates for selection state
		if m.state == "select-version" {
			var listCmd tea.Cmd
			m.list, listCmd = m.list.Update(msg)
			return m, listCmd
		}
	}

	return m, cmd
}

// View renders the restore model
func (m RestoreModel) View() string {
	switch m.state {
	case "loading-versions":
		return fmt.Sprintf("\n  %s Loading available versions for '%s'...\n\n", 
			m.spinner.View(), m.modName)
			
	case "select-version":
		header := styles.TitleStyle.Render(fmt.Sprintf("Restore Mod: %s", m.modName))
		header += "\n" + styles.SubtitleStyle.Render("Select a version to restore")
		header += "\n\n" + styles.NormalTextStyle.Render("Press 'enter' to restore, 'esc' to go back")
		
		content := "\n\n" + m.list.View()
		
		return header + content
		
	case "restoring":
		return fmt.Sprintf("\n  %s Restoring mod '%s' to version '%s'...\n\n", 
			m.spinner.View(), m.modName, m.selectedVersion)
			
	case "success":
		return fmt.Sprintf("\n  ✅ %s\n\n  Press any key to return.\n", m.successMsg)
		
	case "error":
		return fmt.Sprintf("\n  ❌ %s\n\n  Press any key to return.\n", m.errorMsg)
		
	default:
		return "Something went wrong!"
	}
}	