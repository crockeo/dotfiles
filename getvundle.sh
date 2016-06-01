#!/usr/bin/env bash
# Name  : getvundle.sh
# Author: Cerek Hillen
#
# Description:
#   After the .vimrc has been deployed, create a bundle folder and install
#   Vundle into it. Then install all other plugins specified in the .vimrc.
mkdir -p ~/.vim/bundle
cd ~/.vim/bundle
git clone http://github.com/VundleVim/Vundle.vim
vim +PluginInstall +qall
