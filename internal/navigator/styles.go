package navigator

import (
	"charm.land/lipgloss/v2"
)

var (
	HeaderStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FFFFFF")).
			Background(lipgloss.Color("#5F5FD7")).
			Padding(0, 1)

	SubtitleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#888888")).
			Padding(0, 1)

	CursorStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#04B575"))

	ParentStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#5F87FF"))

	NormalStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#DDDDDD"))

	ScrollIndicatorStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#666666")).
				Padding(0, 3)

	HelpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#626262")).
			Padding(0, 1)
)
