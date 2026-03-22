#!/usr/bin/env bash
set -euo pipefail

REPO="yuntasha/cdui"
INSTALL_DIR="${HOME}/.local/bin"
INIT_LINE='eval "$(cdui init SHELL_NAME)"'

info()  { printf "\033[1;34m[info]\033[0m  %s\n" "$1"; }
ok()    { printf "\033[1;32m[ok]\033[0m    %s\n" "$1"; }
warn()  { printf "\033[1;33m[warn]\033[0m  %s\n" "$1"; }
error() { printf "\033[1;31m[error]\033[0m %s\n" "$1"; exit 1; }

detect_shell() {
    local sh
    sh="$(basename "$SHELL")"
    case "$sh" in
        zsh)  echo "zsh"  ;;
        bash) echo "bash" ;;
        *)    echo "zsh"  ;;
    esac
}

get_rc_file() {
    case "$1" in
        zsh)  echo "${HOME}/.zshrc"  ;;
        bash) echo "${HOME}/.bashrc" ;;
    esac
}

# --- Check Go ---
check_go() {
    if ! command -v go &>/dev/null; then
        error "Go가 설치되어 있지 않습니다. https://go.dev/dl 에서 설치해주세요."
    fi
    info "Go 확인: $(go version)"
}

# --- Build & Install ---
build_and_install() {
    local tmpdir
    tmpdir="$(mktemp -d)"
    trap 'rm -rf "$tmpdir"' EXIT

    info "소스 다운로드 중... (github.com/${REPO})"
    git clone --depth 1 "https://github.com/${REPO}.git" "$tmpdir/cdui" 2>/dev/null

    info "빌드 중..."
    (cd "$tmpdir/cdui" && go build -o cdui .)

    mkdir -p "$INSTALL_DIR"
    cp "$tmpdir/cdui/cdui" "$INSTALL_DIR/cdui"
    chmod +x "$INSTALL_DIR/cdui"
    ok "바이너리 설치 완료: ${INSTALL_DIR}/cdui"
}

# --- PATH 확인 ---
check_path() {
    if [[ ":$PATH:" != *":${INSTALL_DIR}:"* ]]; then
        warn "${INSTALL_DIR}이 PATH에 없습니다."
        local shell_name rc_file
        shell_name="$(detect_shell)"
        rc_file="$(get_rc_file "$shell_name")"
        local path_line="export PATH=\"\${HOME}/.local/bin:\${PATH}\""

        if [ -f "$rc_file" ] && grep -qF '.local/bin' "$rc_file"; then
            info "PATH 설정이 ${rc_file}에 이미 존재합니다."
        else
            echo "" >> "$rc_file"
            echo "# cdui - PATH" >> "$rc_file"
            echo "$path_line" >> "$rc_file"
            ok "PATH 설정 추가: ${rc_file}"
        fi
    else
        ok "PATH 확인 완료"
    fi
}

# --- Shell 연동 ---
setup_shell_init() {
    local shell_name rc_file init_line
    shell_name="$(detect_shell)"
    rc_file="$(get_rc_file "$shell_name")"
    init_line="eval \"\$(cdui init ${shell_name})\""

    if [ -f "$rc_file" ] && grep -qF "cdui init" "$rc_file"; then
        ok "쉘 연동이 ${rc_file}에 이미 설정되어 있습니다."
        return
    fi

    echo "" >> "$rc_file"
    echo "# cdui - shell integration" >> "$rc_file"
    echo "$init_line" >> "$rc_file"
    ok "쉘 연동 추가: ${rc_file}"
}

# --- Main ---
main() {
    echo ""
    echo "  ╔═══════════════════════════╗"
    echo "  ║   cdui installer          ║"
    echo "  ╚═══════════════════════════╝"
    echo ""

    check_go
    build_and_install
    check_path
    setup_shell_init

    echo ""
    ok "설치가 완료되었습니다!"
    info "새 터미널을 열거나 아래 명령어를 실행하세요:"
    echo ""
    echo "    source $(get_rc_file "$(detect_shell)")"
    echo ""
    info "사용법: cdui 를 입력하면 디렉토리 네비게이터가 실행됩니다."
    echo ""
}

main
