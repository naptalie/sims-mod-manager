package cmd

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"

	"github.com/naptalie/sims4-mod-manager/internal/config"
	"github.com/naptalie/sims4-mod-manager/internal/core"
	"github.com/naptalie/sims4-mod-manager/internal/ui/styles"
)

var versionsCmd = &cobra.Command{
	Use:   "versions",
	Short: "List all mod versions",
	Run: func(cmd *cobra.Command, args []string) {
		listModVersions()
	},
}

// Function to list all mod versions
func listModVersions() {
	fmt.Println(styles.TitleStyle.Render("🔄 Mod Versions:"))
	fmt.Println()
	
	// Get versions from backup folder structure
	modVersions, err := core.GetModVersions(config.AppConfig.ModStoragePath)
	if err != nil {
		fmt.Printf("Error reading backup folder: %v\n", err)
		return
	}
	
	if len(modVersions) == 0 {
		fmt.Println(styles.NormalTextStyle.Render("No mod backups found. Use 'backup' command to create backups."))
		return
	}
	
	// Set up the table
	columns := []table.Column{
		{Title: "Mod Name", Width: 40},
		{Title: "Available Versions", Width: 30},
	}
	
	rows := []table.Row{}
	
	for modName, versions := range modVersions {
		versionStr := strings.Join(versions, ", ")
		rows = append(rows, table.Row{modName, versionStr})
	}
	
	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(10),
	)
	
	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(styles.PlumbobColor).
		BorderBottom(true).
		Bold(true)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("#ffffff")).
		Background(styles.PlumbobColor).
		Bold(true)
	t.SetStyles(s)
	
	fmt.Println(styles.BoxStyle.Render(t.View()))
	fmt.Printf("\nTotal: %d mods with backups\n", len(modVersions))
}