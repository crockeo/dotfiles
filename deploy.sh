#!/usr/bin/env sh

cd $(dirname $0)

mkdir -p ~/bin
mkdir -p ~/.config/kitty

# git
ln -s $PWD/.gitconfig ~/.gitconfig

# kitty
ln -s $PWD/kitty.conf ~/.config/kitty/kitty.conf
ln -s $PWD/current-theme.conf ~/.config/kitty/current-theme.conf

# tmux
ln -s $PWD/.tmux.conf ~/.tmux.conf

# shell
ln -s $PWD/.zshrc ~/.zshrc

# misc tools
ln -s $PWD/make-packing-checklist ~/bin
