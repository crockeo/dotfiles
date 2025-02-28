#!/usr/bin/env zsh
# shellcheck disable=all
if [[ -f "/opt/homebrew/bin/brew" && -z "$HOMEBREW_REPOSITORY" ]]; then
    eval $(/opt/homebrew/bin/brew shellenv)
fi
export HOMEBREW_NO_AUTO_UPDATE=true

if [ -f ~/.fzf.zsh ]; then
    source ~/.fzf.zsh
fi

if [[ -d "$HOME/.cargo/env" ]]; then
    source "$HOME/.cargo/env"
fi

# Starting tmux
if [ "$EMACS" = "" ] && [ "$TERM_PROGRAM" != "tmux" ]
then
    tmux new-session -A -s foobar
fi

# Lines configured by zsh-newuser-install
HISTFILE=~/.zsh_histfile
HISTSIZE=1000
SAVEHIST=1000

function pjcd() {
    project=$(find-project "$1") || return
    cd "$project"
}

# Loading better auto-completion.
if [[ -f "$HOME/.zsh-completion" ]]; then
    fpath=(~/.zsh-completion $fpath)
fi
autoload -Uz compinit
compinit

# Setting the terminal prompt.
export PS1="%~$ "
if [[ ! -z "${IN_NIX_SHELL:-}" ]]; then
    export PS1="[nix] $PS1"
fi

# Making the colors all colored!
export CLICOLOR=1
export LSCOLORS=gxBxhxDxfxhxhxhxhxcxcx

# Adding my only local bin.
export PATH="$HOME/bin:$PATH"

# Making the terminal run with 256 colors.
export TERM=screen-256color

# Configuring golang
export GOPATH="$HOME/go"
export PATH="$PATH:$HOME/go/bin"

export EDITOR=hx

# Setting terminal mode to Emacs mode, so I can use fun things like ^A, ^E
bindkey -e

# Set up zoxide
eval "$(zoxide init zsh)"
alias cd=z

if [ -f ~/.company.zshrc ]; then
    source ~/.company.zshrc
fi

if [ -e '/nix/var/nix/profiles/default/etc/profile.d/nix-daemon.sh' ]; then
  . '/nix/var/nix/profiles/default/etc/profile.d/nix-daemon.sh'
fi

function venv() {
    if [[ -f ".venv/bin/activate" ]]; then
        source .venv/bin/activate
    elif [[ -f "venv/bin/activate" ]]; then
        source venv/bin/activate
    fi
}

if [[ -d "$HOME/.bun" ]]; then
    # bun completions
    [ -s "/Users/crockeo/.bun/_bun" ] && source "/Users/crockeo/.bun/_bun"

    # bun
    export BUN_INSTALL="$HOME/.bun"
    export PATH="$BUN_INSTALL/bin:$PATH"
fi

function cb() {
	git branch -vv --sort=-committerdate --color \
      | fzf --ansi --header Checkout \
      | rg '^\s*(\S+)' -o -r '$1' \
      | xargs git checkout
}
