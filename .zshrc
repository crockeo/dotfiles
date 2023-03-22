if [[ -f "/opt/homebrew/bin/brew" && -z "$HOMEBREW_REPOSITORY" ]]; then
    eval $(/opt/homebrew/bin/brew shellenv)
fi

# Starting tmux
if [ "$EMACS" = "" ] && [ "$TMUX" = "" ]
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
autoload -Uz compinit
compinit

# Setting the terminal prompt.
PS1="%~$ "
if [ -n "$IN_NIX_SHELL" ]; then
    PS1="[nix] ${PS1}"
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

if [ -f ~/.company.zshrc ]; then
    source ~/.company.zshrc
fi

if [ -e '/nix/var/nix/profiles/default/etc/profile.d/nix-daemon.sh' ]; then
  . '/nix/var/nix/profiles/default/etc/profile.d/nix-daemon.sh'
fi
