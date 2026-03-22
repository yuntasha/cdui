package shell

import (
	_ "embed"
	"fmt"
)

//go:embed templates/bash.sh
var bashScript string

//go:embed templates/zsh.sh
var zshScript string

// PrintInitScript prints the shell wrapper function for the given shell.
func PrintInitScript(shell string) error {
	switch shell {
	case "bash":
		fmt.Print(bashScript)
	case "zsh":
		fmt.Print(zshScript)
	default:
		return fmt.Errorf("unsupported shell: %s (supported: bash, zsh)", shell)
	}
	return nil
}
