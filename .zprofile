#!/usr/bin/env zsh
# shellcheck disable=all
if [[ -f "/opt/homebrew/bin/brew" && -z "$HOMEBREW_REPOSITORY" ]]; then
    eval $(/opt/homebrew/bin/brew shellenv)
fi

if [ -f ~/.company.zprofile ]; then
    source ~/.company.zprofile
fi

# Adding my only local bin.
export PATH="$HOME/bin:$PATH"

# Adding GOPATH bin (for globally installed tools, like gopls)
export GOPATH="$HOME/go"
export PATH="$PATH:$HOME/go/bin"
