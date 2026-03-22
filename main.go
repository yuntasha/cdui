package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"

	"github.com/yuntasha/cdui/internal/navigator"
	"github.com/yuntasha/cdui/internal/shell"
)

func main() {
	// Handle "init" subcommand
	if len(os.Args) >= 2 && os.Args[1] == "init" {
		if len(os.Args) < 3 {
			fmt.Fprintln(os.Stderr, "usage: cdui init <bash|zsh>")
			os.Exit(1)
		}
		if err := shell.PrintInitScript(os.Args[2]); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		return
	}

	// Determine start directory
	startDir := "."
	if len(os.Args) >= 2 {
		startDir = os.Args[1]
	}

	m := navigator.New(startDir)

	// TUI renders to stderr so stdout is free for the selected path
	p := tea.NewProgram(m, tea.WithOutput(os.Stderr))

	finalModel, err := p.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}

	// Output selected path to stdout for shell wrapper to capture
	if result, ok := finalModel.(navigator.Model); ok && result.Selected != "" {
		fmt.Print(result.Selected)
	}
}
