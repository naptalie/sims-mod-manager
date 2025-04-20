package main

import (
	"fmt"
	"github.com/naptalie/sims4-mod-manager/cmd"
	"github.com/naptalie/sims4-mod-manager/internal/config"
	"github.com/naptalie/sims4-mod-manager/internal/ui/styles"
)

func main() {
	// Initialize config
	config.InitConfig()
	
	// Display ASCII Art and Title
	fmt.Println(styles.RenderPlumbob())
	fmt.Println(styles.TitleStyle.Render("Sims 4 Mod Manager"))
	fmt.Println(styles.SubtitleStyle.Render("Version control for your Simming adventures!"))
	fmt.Println()
	
	// Execute the root command
	cmd.Execute()
}