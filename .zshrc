eval $(/opt/homebrew/bin/brew shellenv)

# Starting tmux
if [ "$EMACS" = "" ] && [ "$TMUX" = "" ]
then
    tmux new-session -A -s foobar
fi

# Lines configured by zsh-newuser-install
HISTFILE=~/.zsh_histfile
HISTSIZE=1000
SAVEHIST=1000

# Loading better auto-completion.
autoload -Uz compinit
compinit

# Setting the terminal prompt.
PS1="%~$ "

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

# Setting terminal mode to Emacs mode, so I can use fun things like ^A, ^E
bindkey -e
