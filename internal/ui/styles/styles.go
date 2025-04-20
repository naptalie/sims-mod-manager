package styles

import (
	"github.com/charmbracelet/lipgloss"
)

// Sims 4 Theme Colors
var (
	GreenColor      = lipgloss.Color("#3fdb57")
	BlueColor       = lipgloss.Color("#3f93db")
	LightBlueColor  = lipgloss.Color("#87cefa")
	PlumbobColor    = lipgloss.Color("#17f174")
	BackgroundColor = lipgloss.Color("#2d2d2d")
	
	// Styles
	TitleStyle = lipgloss.NewStyle().
			Foreground(PlumbobColor).
			Bold(true).
			MarginLeft(2)
	
	SubtitleStyle = lipgloss.NewStyle().
			Foreground(BlueColor).
			Italic(true).
			MarginLeft(4)
	
	NormalTextStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#ffffff"))
	
	HighlightStyle = lipgloss.NewStyle().
			Foreground(PlumbobColor).
			Bold(true)
	
	BoxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(PlumbobColor).
			Padding(1).
			BorderTop(true).
			BorderLeft(true).
			BorderRight(true).
			BorderBottom(true)
			
	SearchStyle = lipgloss.NewStyle().
			Border(lipgloss.NormalBorder()).
			BorderForeground(PlumbobColor).
			Padding(0, 1).
			BorderBottom(true)
)

// PLUMBOB ASCII ART
const plumbobArt = `
     /\
    /  \
   /    \
  /      \
 /        \
/          \
\          /
 \        /
  \      /
   \    /
    \  /
     \/
`

// RenderPlumbob renders the plumbob ASCII art with the correct color
func RenderPlumbob() string {
	return lipgloss.NewStyle().Foreground(PlumbobColor).Render(plumbobArt)
}