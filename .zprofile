eval $(/opt/homebrew/bin/brew shellenv)

# Adding my only local bin.
export PATH="$HOME/bin:$PATH"

# Adding GOPATH bin (for globally installed tools, like gopls)
export GOPATH="$HOME/go"
export PATH="$PATH:$HOME/go/bin"
