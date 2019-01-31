# Starting tmux
if [ "$TMUX" = "" ]
then
    tmux attach
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

# Aliasing old vims to new vims
alias vi="nvim"
alias vim="nvim"

# Setting the default editor to nvim
export EDITOR=nvim

# Configuring NVM. Uncomment if you need NVM. Causes shell to open very slowly.
#export NVM_DIR="$HOME/.nvm"
#[ -s "/usr/local/opt/nvm/nvm.sh" ] && . "/usr/local/opt/nvm/nvm.sh"
#[ -s "/usr/local/opt/nvm/etc/bash_completion" ] && . "/usr/local/opt/nvm/etc/bash_completion"

# Configuring pyenv
if type "pyenv" > /dev/null; then
  eval "$(pyenv init -)"
fi

# Configuring golang
export GOPATH="$HOME/go"
export PATH="$PATH:$HOME/go/bin"

# Creating a path to clang-format version 7.
export CLANG_FORMAT="/usr/local/opt/llvm@7/bin/clang-format"

# Setting terminal mode to Emacs mode, so I can use fun things like ^A, ^E
bindkey -e
