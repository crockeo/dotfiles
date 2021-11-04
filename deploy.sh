#!/usr/bin/env sh

cd $(dirname $0)

mkdir -p ~/.config/kitty

ln -s $PWD/.gitconfig ~/.gitconfig
ln -s $PWD/kitty.conf ~/.config/kitty/kitty.conf
ln -s $PWD/.tmux.conf ~/.tmux.conf
ln -s $PWD/.zshrc ~/.zshrc
