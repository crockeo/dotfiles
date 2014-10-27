#!/bin/sh

mv ~/.vim/ ~/.vim_old/
mv ~/.vimrc ~/.vimrc_old

cp -r ./.vim ~/
cp ./.vimrc ~/
