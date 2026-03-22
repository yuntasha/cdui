cdui() {
    local result
    result="$(\command cdui "$@")"
    local exit_code=$?
    if [[ $exit_code -eq 0 ]] && [[ -n "$result" ]] && [[ -d "$result" ]]; then
        \builtin cd -- "$result"
    fi
}
