# cdui

Terminal UI for fast directory navigation. Browse directories and files interactively, `cd` into directories, and open files in vim — all from a keyboard-driven TUI.

## Installation

### Quick Install (recommended)

```bash
bash <(curl -s https://raw.githubusercontent.com/yuntasha/cdui/main/install.sh)
```

This script automatically handles everything: build, binary install, PATH setup, and shell integration.

### go install

```bash
go install github.com/yuntasha/cdui@latest
```

### Build from source

```bash
git clone https://github.com/yuntasha/cdui.git
cd cdui
make install
```

The binary is installed to `~/.local/bin/`. Make sure this directory is in your `PATH`.

## Shell Setup

If you used the quick install script, this is already done. Otherwise, add the following to your shell configuration:

**zsh** (`~/.zshrc`):

```bash
eval "$(cdui init zsh)"
```

**bash** (`~/.bashrc`):

```bash
eval "$(cdui init bash)"
```

Restart your shell or run `source ~/.zshrc` (or `~/.bashrc`) to apply.

## Usage

Run `cdui` to start the browser from the current directory, or pass a path:

```bash
cdui            # start from current directory
cdui ~/projects # start from ~/projects
```

### Key Bindings

| Key | Action |
|-----|--------|
| `j` / `↓` | Move cursor down |
| `k` / `↑` | Move cursor up |
| `l` / `→` / `Enter` | Open directory / open file in vim |
| `h` / `←` / `Backspace` | Go to parent |
| `Space` / `Tab` | Select current directory and `cd` |
| `g` / `Home` | Go to top |
| `G` / `End` | Go to bottom |
| `.` | Toggle hidden entries |
| `q` / `Esc` / `Ctrl+C` | Quit without changing directory |

## Requirements

- Go 1.24+
- zsh or bash

## License

MIT
