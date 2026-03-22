# cdui

Terminal UI for fast directory navigation. Browse and `cd` into directories interactively using a keyboard-driven TUI.

## Installation

### go install (recommended)

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

After installing the binary, add the following to your shell configuration so that `cdui` can change your working directory:

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

Run `cdui` to start the directory browser from the current directory, or pass a path:

```bash
cdui            # start from current directory
cdui ~/projects # start from ~/projects
```

### Key Bindings

| Key | Action |
|-----|--------|
| `j` / `↓` | Move cursor down |
| `k` / `↑` | Move cursor up |
| `l` / `→` / `Enter` | Open directory |
| `h` / `←` / `Backspace` | Go to parent |
| `Space` / `Tab` | Select current directory and `cd` |
| `g` / `Home` | Go to top |
| `G` / `End` | Go to bottom |
| `.` | Toggle hidden directories |
| `q` / `Esc` / `Ctrl+C` | Quit without changing directory |

## Requirements

- Go 1.24+
- zsh or bash

## License

MIT
