// Update handles all the standard bubble tea update things
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "esc":
			if m.State != "main" {
				m.State = "main"
				m.ErrorMsg = ""
				m.SuccessMsg = ""
				return m, cmd
			}
		case "b":
			// Only respond to 'b' key in main view
			if m.State == "main" {
				// Get currently selected mod, if any
				selectedItem := m.List.SelectedItem()
				if selectedItem != nil {
					modItem, ok := selectedItem.(ModItem)
					if ok {
						m.SelectedMod = modItem.mod.Name
						// Start backup process for selected mod
						return m, startBackup(m.SelectedMod)
					}
				}
			}
		}
		
	case modsLoadedMsg:
		if msg.err != nil {
			m.ErrorMsg = msg.err.Error()
			package tui
	
		}
	}
}

import (
	"strings"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/yourusername/sims4-mod-manager/internal/models"
	"github.com/yourusername/sims4-mod-manager/internal/ui/styles"
)

// Model represents the application state for the TUI
type Model struct {
	State         string
	List          list.Model
	Spinner       spinner.Model
	TextInput     textinput.Model
	Mods          []models.Mod
	SelectedMod   string
	Versions      []string
	ErrorMsg      string
	SuccessMsg    string
	SearchInput   textinput.Model
	FilteredMods  []models.Mod
}

// Custom message types
type modsLoadedMsg struct {
	mods []models.Mod
	err  error
}

// NewMainModel creates a new main model ready to load mods
func NewMainModel() (Model, tea.Cmd) {
	// Create a spinner
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(styles.PlumbobColor)
	
	// Create a text input for search
	ti := textinput.New()
	ti.Placeholder = "Search mods..."
	ti.Focus()
	ti.CharLimit = 50
	ti.Width = 30
	
	// Initialize the model
	m := Model{
		State:       "loading",
		Spinner:     s,
		SearchInput: ti,
		Mods:        []models.Mod{},
		FilteredMods: []models.Mod{},
	}
	
	return m, loadMods()
}

// Initialize the model
func InitialModel() tea.Model {
	m, _ := NewMainModel()
	return m
}

// Init is the first function that will be called when your program starts
func (m Model) Init() tea.Cmd {
	if m.State == "loading" {
		return tea.Batch(
			m.Spinner.Tick,
			loadMods(),
		)
	}
	return nil
}

// Update handles all the standard bubble tea update things
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "esc":
			if m.State != "main" {
				m.State = "main"
				m.ErrorMsg = ""
				m.SuccessMsg = ""
				return m, cmd
			}
		case "b":
			// Only respond to 'b' key in main view
			if m.State == "main" {
				// Get currently selected mod, if any
				selectedItem := m.List.SelectedItem()
				if selectedItem != nil {
					modItem, ok := selectedItem.(ModItem)
					if ok {
						m.SelectedMod = modItem.mod.Name
						// Start backup process for selected mod
						return m, startBackup(m.SelectedMod)
					}
				}
			}
		case "a":
			// Backup all mods
			if m.State == "main" {
				m.SelectedMod = ""
				return m, startBackup("")
			}
		case "r":
			// Restore selected mod
			if m.State == "main" {
				// Get currently selected mod, if any
				selectedItem := m.List.SelectedItem()
				if selectedItem != nil {
					modItem, ok := selectedItem.(ModItem)
					if ok {
						m.SelectedMod = modItem.mod.Name
						// Start restore process for selected mod
						return m, startRestore(m.SelectedMod)
					}
				}
			}
		}
		
	case modsLoadedMsg:
		if msg.err != nil {
			m.ErrorMsg = msg.err.Error()
			m.State = "error"
			return m, nil
		}
		
		m.Mods = msg.mods
		m.FilteredMods = msg.mods
		
		// Initialize the list with mods
		var items []list.Item
		for _, mod := range msg.mods {
			items = append(items, NewModItem(mod))
		}
		
		// Create a new list
		l := list.New(items, list.NewDefaultDelegate(), 0, 0)
		l.Title = "Your Sims 4 Mods"
		l.SetShowStatusBar(false)
		l.SetFilteringEnabled(false)
		l.Styles.Title = styles.TitleStyle
		l.Styles.NoItems = styles.NormalTextStyle
		
		m.List = l
		m.State = "main"
		return m, nil
		
	case backupStartMsg:
		// Switch to backup state
		m.State = "backing-up"
		m.SelectedMod = msg.modName
		
		// Create a backup model and initialize it
		backupModel := NewBackupModel(msg.modName)
		return backupModel, backupModel.Init()
		
	case backupFinishedMsg:
		// Handle backup completion
		if msg.success {
			m.SuccessMsg = msg.message
		} else {
			m.ErrorMsg = msg.message
		}
		
		// Return to main state
		m.State = "main"
		return m, nil
		
	case restoreStartMsg:
		// Switch to restore state
		m.State = "restoring"
		m.SelectedMod = msg.modName
		
		// Create a restore model and initialize it
		restoreModel := NewRestoreModel(msg.modName)
		return restoreModel, restoreModel.Init()
		
	case restoreFinishedMsg:
		// Handle restore completion
		if msg.success {
			m.SuccessMsg = msg.message
		} else {
			m.ErrorMsg = msg.message
		}
		
		// Return to main state and refresh mod list
		m.State = "main"
		return m, loadMods()()
	}
	
	// Different updates based on state
	switch m.State {
	case "loading":
		// Update the spinner
		var spinnerCmd tea.Cmd
		m.Spinner, spinnerCmd = m.Spinner.Update(msg)
		return m, spinnerCmd
		
	case "main":
		// Update the list
		var listCmd tea.Cmd
		m.List, listCmd = m.List.Update(msg)
		cmd = listCmd
		
		// Handle search input
		var searchCmd tea.Cmd
		m.SearchInput, searchCmd = m.SearchInput.Update(msg)
		cmd = tea.Batch(cmd, searchCmd)
		
		// Apply search filtering
		searchTerm := m.SearchInput.Value()
		if searchTerm != "" {
			// Filter has changed, update the list
			filtered := filterMods(m.Mods, searchTerm)
			m.FilteredMods = filtered
			
			// Update list items
			var items []list.Item
			for _, mod := range filtered {
				items = append(items, NewModItem(mod))
			}
			
			m.List.SetItems(items)
		} else if len(m.FilteredMods) != len(m.Mods) {
			// Reset to full list if search is cleared
			m.FilteredMods = m.Mods
			var items []list.Item
			for _, mod := range m.Mods {
				items = append(items, NewModItem(mod))
			}
			m.List.SetItems(items)
		}
		
		return m, cmd
	}
	
	return m, cmd
}

// View renders the current UI based on the model
func (m Model) View() string {
	// Different views based on state
	switch m.State {
	case "loading":
		return renderLoading(m)
	case "error":
		return renderError(m)
	case "main":
		return renderMain(m)
	default:
		return "Something went wrong!"
	}
}

// Helper functions to render different UI states
func renderLoading(m Model) string {
	return "\n  " + m.Spinner.View() + " Loading mods...\n\n"
}

func renderError(m Model) string {
	return "\n  " + m.ErrorMsg + "\n\n  Press any key to quit.\n"
}

func renderMain(m Model) string {
	// Build the header
	header := styles.TitleStyle.Render("The Sims 4 Mod Manager")
	header += "\n" + styles.SubtitleStyle.Render("Version control for your Simming adventures!")
	header += "\n\n" + styles.NormalTextStyle.Render("Press 'q' to quit, 'b' to backup selected mod, 'a' to backup all mods, 'r' to restore")
	
	// Build the search input
	search := "\n\n  🔍 " + m.SearchInput.View()
	
	// Build the main content
	content := "\n\n" + m.List.View()
	
	// Build the footer with messages
	footer := "\n\n"
	if m.ErrorMsg != "" {
		footer += lipgloss.NewStyle().Foreground(lipgloss.Color("#ff0000")).Render("  " + m.ErrorMsg)
	}
	if m.SuccessMsg != "" {
		footer += lipgloss.NewStyle().Foreground(styles.GreenColor).Render("  " + m.SuccessMsg)
	}
	
	return header + search + content + footer
}