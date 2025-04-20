package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/naptalie/sims4-mod-manager/internal/ui/tui"
)

var uiCmd = &cobra.Command{
	Use:   "ui",
	Short: "Launch the TUI interface",
	Run: func(cmd *cobra.Command, args []string) {
		launchUI()
	},
}

// Launch the UI
func launchUI() {
	// Create and initialize the model
	m := tui.InitialModel()
	
	// Start the Bubble Tea program
	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running UI: %v\n", err)
		os.Exit(1)
	}
}