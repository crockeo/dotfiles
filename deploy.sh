#!/usr/bin/env sh

cd $(dirname $0)

mkdir -p ~/.config/nvim

ln .zshrc ~/.zshrc
ln .tmux.conf ~/.tmux.conf
ln init.vim ~/.config/nvim/init.vim
ln .gitconfig ~/.gitconfig

curl -fLo ~/.local/share/nvim/site/autoload/plug.vim --create-dirs \
	https://raw.githubusercontent.com/junegunn/vim-plug/master/plug.vim
nvim -c "PlugInstall"
