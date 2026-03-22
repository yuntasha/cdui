package navigator

import (
	"fmt"
	"path/filepath"
	"strings"

	tea "charm.land/bubbletea/v2"
)

type Model struct {
	currentDir  string
	entries     []DirEntry
	cursor      int
	offset      int
	visibleRows int
	showHidden  bool
	Selected    string // The selected directory path (empty if cancelled)
	quitting    bool
	width       int
}

func New(startDir string) Model {
	absDir, err := filepath.Abs(startDir)
	if err != nil {
		absDir = startDir
	}

	m := Model{
		currentDir:  absDir,
		visibleRows: 20,
	}
	m.loadEntries()
	return m
}

func (m *Model) loadEntries() {
	m.entries = ReadDirs(m.currentDir, m.showHidden)
	m.cursor = 0
	m.offset = 0
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		// Reserve lines for header(1) + subtitle(1) + blank(1) + help(1) + blank(1) = 5
		m.visibleRows = msg.Height - 5
		if m.visibleRows < 3 {
			m.visibleRows = 3
		}
		m.clampViewport()
		return m, nil

	case tea.KeyPressMsg:
		// Select and quit (cd to this directory)
		if msg.Code == tea.KeySpace || msg.Code == tea.KeyTab {
			m.Selected = m.currentDir
			m.quitting = true
			return m, tea.Quit
		}

		switch msg.String() {
		// Quit without selection
		case "q", "esc", "ctrl+c":
			m.quitting = true
			return m, tea.Quit

		// Move cursor up
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
				m.clampViewport()
			}

		// Move cursor down
		case "down", "j":
			if m.cursor < len(m.entries)-1 {
				m.cursor++
				m.clampViewport()
			}

		// Enter directory
		case "enter", "right", "l":
			if len(m.entries) > 0 {
				m.enterDir()
			}

		// Go to parent directory
		case "backspace", "left", "h":
			m.goParent()

		// Toggle hidden directories
		case ".":
			m.showHidden = !m.showHidden
			m.loadEntries()

		// Go to top
		case "g", "home":
			m.cursor = 0
			m.offset = 0

		// Go to bottom
		case "G", "end":
			if len(m.entries) > 0 {
				m.cursor = len(m.entries) - 1
				m.clampViewport()
			}
		}
	}

	return m, nil
}

func (m *Model) enterDir() {
	entry := m.entries[m.cursor]
	var newDir string
	if entry.Name == ".." {
		newDir = filepath.Dir(m.currentDir)
	} else {
		newDir = filepath.Join(m.currentDir, entry.Name)
	}

	// Verify the directory is accessible
	newEntries := ReadDirs(newDir, m.showHidden)
	if newEntries != nil {
		m.currentDir = newDir
		m.entries = newEntries
		m.cursor = 0
		m.offset = 0
	}
}

func (m *Model) goParent() {
	parent := filepath.Dir(m.currentDir)
	if parent == m.currentDir {
		return // Already at root
	}

	oldDir := filepath.Base(m.currentDir)
	m.currentDir = parent
	m.loadEntries()

	// Try to position cursor on the directory we came from
	for i, e := range m.entries {
		if e.Name == oldDir {
			m.cursor = i
			m.clampViewport()
			break
		}
	}
}

func (m *Model) clampViewport() {
	if m.cursor < m.offset {
		m.offset = m.cursor
	}
	if m.cursor >= m.offset+m.visibleRows {
		m.offset = m.cursor - m.visibleRows + 1
	}
}

func (m Model) View() tea.View {
	if m.quitting {
		return tea.NewView("")
	}

	var b strings.Builder

	// Header: current path
	header := m.currentDir
	if m.width > 0 {
		headerStyle := HeaderStyle.Width(m.width)
		b.WriteString(headerStyle.Render(header))
	} else {
		b.WriteString(HeaderStyle.Render(header))
	}
	b.WriteString("\n")

	// Subtitle: directory count and hidden status
	dirCount := CountDirs(m.currentDir, m.showHidden)
	hiddenStatus := "hidden dirs hidden"
	if m.showHidden {
		hiddenStatus = "hidden dirs shown"
	}
	subtitle := fmt.Sprintf(" %d directories | %s", dirCount, hiddenStatus)
	b.WriteString(SubtitleStyle.Render(subtitle))
	b.WriteString("\n\n")

	// Directory list with scrolling
	if len(m.entries) == 0 {
		b.WriteString(SubtitleStyle.Render("  (empty directory)"))
		b.WriteString("\n")
	} else {
		end := m.offset + m.visibleRows
		if end > len(m.entries) {
			end = len(m.entries)
		}

		// Scroll up indicator
		if m.offset > 0 {
			b.WriteString(ScrollIndicatorStyle.Render("... more above"))
			b.WriteString("\n")
		}

		for i := m.offset; i < end; i++ {
			entry := m.entries[i]
			name := entry.Name
			if entry.IsDir && name != ".." {
				name += "/"
			} else if name == ".." {
				name += "/"
			}

			if i == m.cursor {
				b.WriteString(CursorStyle.Render(fmt.Sprintf("  > %s", name)))
			} else if entry.Name == ".." {
				b.WriteString(ParentStyle.Render(fmt.Sprintf("    %s", name)))
			} else {
				b.WriteString(NormalStyle.Render(fmt.Sprintf("    %s", name)))
			}
			b.WriteString("\n")
		}

		// Scroll down indicator
		if end < len(m.entries) {
			b.WriteString(ScrollIndicatorStyle.Render("... more below"))
			b.WriteString("\n")
		}
	}

	// Help bar
	b.WriteString("\n")
	help := " enter: open | space: select & cd | backspace: parent | .: hidden | q: quit"
	b.WriteString(HelpStyle.Render(help))

	return tea.NewView(b.String())
}
