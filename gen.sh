#!/bin/bash
#!/bin/zsh

# swap for shell lint, for it does not support zsh now

# install gum yq:
# go install github.com/charmbracelet/gum@latest
# go install github.com/mikefarah/yq/v4@latest
# go install github.com/ogen-go/ogen/cmd/ogen@latest

function env_check() {
    if ! command -v ogen >/dev/null; then
        echo "ogen not found, please install with 'go install github.com/ogen-go/ogen/cmd/ogen@latest'"
        exit 1
    fi
    if ! command -v gum >/dev/null; then
        echo "gum not found, please install with 'go install github.com/charmbracelet/gum@latest'"
        exit 1
    fi
    if ! command -v yq >/dev/null; then
        echo "yq not found, please install with 'go install github.com/mikefarah/yq/v4@latest' first"
        exit 1
    fi
    echo "shell env is ready."
}

function choose_target() {
    local target
    target=$(gum choose "a" "b" "c")
    echo "$target"
}

function input_version() {
    local version
    version=$(
        gum input \
            --header='please input a version for this spec. eg: v1' \
            --placeholder="v1" \
            --value="v1"
    )
    echo "$version"
}

function entgo_code_gen() {
    local target=$1
    gum spin --show-output --spinner dot --title "Generating entgo Code..." \
        -- go run cmd/gen/main.go "${target}" &&
        echo "code gen ok" || echo 'code gen failed'
}

function main() {
    gum format <<EOF
# 代码生成工具
## Usage
* 按照提示进行即可.
EOF
    if [[ $# -eq 0 ]]; then
        action=$(gum choose \
            "entgo code gen core" \
            "Run ogen ogm for vip & version" \
            "quit")
    else
        action="entgo code gen"
    fi

    case $action in
        "entgo code gen core")
            entgo_code_gen "core"
            ;;
        "quit")
            exit
            ;;
        *)
            echo 'never here, please check source'
            ;;
    esac
}

env_check
main "$@"
