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

# Adding the go path.
export GOPATH="$HOME/Documents/Development/Go"

# Adding my only local bin.
export PATH="$PATH:$HOME/usrbin"

# Adding the Go language bin
export PATH="$PATH:$GOPATH/bin"

export PATH="$PATH:$HOME/Library/Haskell/bin"

# Making the terminal run with 256 colors.
export TERM=screen-256color

# Starting tmux
if [ "$TMUX" = "" ]
then
    tmux attach
fi
