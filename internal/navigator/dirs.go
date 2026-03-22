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

	var filteredDirs []DirEntry
	var filteredFiles []DirEntry
	for _, e := range entries {
		name := e.Name()

		// Skip hidden entries if not showing them
		if !showHidden && strings.HasPrefix(name, ".") {
			continue
		}

		// Include directories
		if e.IsDir() {
			filteredDirs = append(filteredDirs, DirEntry{Name: name, IsDir: true})
			continue
		}

		// Check if symlink points to a directory
		if e.Type()&os.ModeSymlink != 0 {
			info, err := os.Stat(path + "/" + name)
			if err == nil && info.IsDir() {
				filteredDirs = append(filteredDirs, DirEntry{Name: name, IsDir: true})
				continue
			}
		}

		// Include files
		filteredFiles = append(filteredFiles, DirEntry{Name: name, IsDir: false})
	}

	// Case-insensitive sort (directories first, then files)
	sort.Slice(filteredDirs, func(i, j int) bool {
		return strings.ToLower(filteredDirs[i].Name) < strings.ToLower(filteredDirs[j].Name)
	})
	sort.Slice(filteredFiles, func(i, j int) bool {
		return strings.ToLower(filteredFiles[i].Name) < strings.ToLower(filteredFiles[j].Name)
	})

	dirs = append(dirs, filteredDirs...)
	dirs = append(dirs, filteredFiles...)
	return dirs
}

// CountEntries counts the number of directories and files (excluding "..") in the given path.
func CountEntries(path string, showHidden bool) (dirs int, files int) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return 0, 0
	}

	for _, e := range entries {
		name := e.Name()
		if !showHidden && strings.HasPrefix(name, ".") {
			continue
		}
		if e.IsDir() {
			dirs++
			continue
		}
		if e.Type()&os.ModeSymlink != 0 {
			info, err := os.Stat(path + "/" + name)
			if err == nil && info.IsDir() {
				dirs++
				continue
			}
		}
		files++
	}
	return dirs, files
}
