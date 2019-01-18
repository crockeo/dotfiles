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
export PATH="$PATH:$HOME/bin"

# Making the terminal run with 256 colors.
export TERM=screen-256color

# Aliasing old vims to new vims
alias vi="nvim"
alias vim="nvim"

# Setting the default editor to nvim
export EDITOR=nvim

# Configuring NVM.
export NVM_DIR="$HOME/.nvm"
[ -s "/usr/local/opt/nvm/nvm.sh" ] && . "/usr/local/opt/nvm/nvm.sh"
[ -s "/usr/local/opt/nvm/etc/bash_completion" ] && . "/usr/local/opt/nvm/etc/bash_completion"

# Configuring pyenv
eval "$(pyenv init -)"

# Configuring golang
export GOPATH="$HOME/go"
export PATH="$PATH:$HOME/go/bin"

# Setting terminal mode to Emacs mode, so I can use fun things like ^A, ^E
bindkey -e

# Starting tmux
if [ "$TMUX" = "" ]
then
    tmux attach
fi
