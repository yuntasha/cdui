package navigator

import (
	"os"
	"sort"
	"strings"
)

type DirEntry struct {
	Name  string
	IsDir bool
}

// ReadDirs reads directory entries from the given path.
// If showHidden is false, entries starting with '.' are filtered out.
// The result always starts with ".." unless the path is "/".
func ReadDirs(path string, showHidden bool) []DirEntry {
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil
	}

	var dirs []DirEntry

	// Add parent directory unless at root
	if path != "/" {
		dirs = append(dirs, DirEntry{Name: "..", IsDir: true})
	}

	var filtered []DirEntry
	for _, e := range entries {
		name := e.Name()

		// Skip hidden directories if not showing them
		if !showHidden && strings.HasPrefix(name, ".") {
			continue
		}

		// Include directories directly
		if e.IsDir() {
			filtered = append(filtered, DirEntry{Name: name, IsDir: true})
			continue
		}

		// Check if symlink points to a directory
		if e.Type()&os.ModeSymlink != 0 {
			info, err := os.Stat(path + "/" + name)
			if err == nil && info.IsDir() {
				filtered = append(filtered, DirEntry{Name: name, IsDir: true})
			}
		}
	}

	// Case-insensitive sort
	sort.Slice(filtered, func(i, j int) bool {
		return strings.ToLower(filtered[i].Name) < strings.ToLower(filtered[j].Name)
	})

	dirs = append(dirs, filtered...)
	return dirs
}

// CountDirs counts the number of directories (excluding "..") in the given path.
func CountDirs(path string, showHidden bool) int {
	entries, err := os.ReadDir(path)
	if err != nil {
		return 0
	}

	count := 0
	for _, e := range entries {
		name := e.Name()
		if !showHidden && strings.HasPrefix(name, ".") {
			continue
		}
		if e.IsDir() {
			count++
			continue
		}
		if e.Type()&os.ModeSymlink != 0 {
			info, err := os.Stat(path + "/" + name)
			if err == nil && info.IsDir() {
				count++
			}
		}
	}
	return count
}
