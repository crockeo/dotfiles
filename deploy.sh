#!/usr/bin/env sh

mkdir -p ~/config/nvim

ln .zshrc ~/.zshrc
ln .tmux.conf ~/.tmux.conf
ln init.vim ~/config/nvim/init.vim

bash ./getplug.sh
