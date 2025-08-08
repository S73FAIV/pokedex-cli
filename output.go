package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

var (
	nameStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFD700")).Bold(true) // gold, bold
	labelStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#00CED1")).Bold(true) // dark turquoise
	valueStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFFFF"))
	boxStyle   = lipgloss.NewStyle().Padding(1, 2).Border(lipgloss.NormalBorder()).BorderForeground(lipgloss.Color("#5F5FFF"))
	//style = lipgloss.ColorProfile().Color()
)

func PrintPokemon(p Pokemon) {
	content := fmt.Sprintf(
		"%s: %s\n%s: %d dm\n%s: %d hg\n",
		labelStyle.Render("Name"), nameStyle.Render(p.Name),
		labelStyle.Render("Height"), p.Height,
		labelStyle.Render("Weight"), p.Weight,
	)

	box := boxStyle.Render(content)
	fmt.Println(box)
}
