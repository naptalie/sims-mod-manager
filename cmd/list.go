package cmd

import (
	"fmt"
	
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
	
	"github.com/naptalie/sims4-mod-manager/internal/config"
	"github.com/naptalie/sims4-mod-manager/internal/core"
	"github.com/naptalie/sims4-mod-manager/internal/ui/styles"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all mods in your Sims 4 mods folder",
	Run: func(cmd *cobra.Command, args []string) {
		listMods()
	},
}

// Function to list all mods in the Sims 4 mods folder
func listMods() {
	fmt.Println(styles.TitleStyle.Render("📁 Mods in your Sims 4 folder:"))
	fmt.Println()
	
	mods, err := core.ScanModsFolder(config.AppConfig.SimsModsPath)
	if err != nil {
		fmt.Printf("Error scanning mods folder: %v\n", err)
		return
	}
	
	if len(mods) == 0 {
		fmt.Println(styles.NormalTextStyle.Render("No mods found in your Sims 4 mods folder."))
		return
	}
	
	// Set up the table
	columns := []table.Column{
		{Title: "Name", Width: 40},
		{Title: "Size", Width: 15},
		{Title: "Last Updated", Width: 20},
	}
	
	rows := []table.Row{}
	
	for _, mod := range mods {
		rows = append(rows, table.Row{
			mod.Name,
			fmt.Sprintf("%.2f MB", float64(mod.Size)/(1024*1024)),
			mod.Updated.Format("2006-01-02 15:04:05"),
		})
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
	fmt.Printf("\nTotal: %d mods\n", len(mods))
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all mods in your Sims 4 mods folder",
	Run: func(cmd *cobra.Command, args []string) {
		listMods()
	},
}

// Function to list all mods in the Sims 4 mods folder
func listMods() {
	fmt.Println(styles.TitleStyle.Render("📁 Mods in your Sims 4 folder:"))
	fmt.Println()
	
	mods, err := core.ScanModsFolder(config.AppConfig.SimsModsPath)
	if err != nil {
		fmt.Printf("Error scanning mods folder: %v\n", err)
		return
	}
	
	if len(mods) == 0 {
		fmt.Println(styles.NormalTextStyle.Render("No mods found in your Sims 4 mods folder."))
		return
	}
	
	// Set up the table
	columns := []table.Column{
		{Title: "Name", Width: 40},
		{Title: "Size", Width: 15},
		{Title: "Last Updated", Width: 20},
	}
	
	rows := []table.Row{}
	
	for _, mod := range mods {
		rows = append(rows, table.Row{
			mod.Name,
			fmt.Sprintf("%.2f MB", float64(mod.Size)/(1024*1024)),
			mod.Updated.Format("2006-01-02 15:04:05"),
		})
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
	fmt.Printf("\nTotal: %d mods\n", len(mods))
}