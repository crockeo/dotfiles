#!/bin/bash

./update-system.sh
mkdir -p ~/.vim/bundle
cd ~/.vim/bundle/
git clone git@github.com:gmarik/Vundle.vim
vim +PluginInstall +qall
