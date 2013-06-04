" Setting no compatible
set nocompatible

" Setting syntax highlighting
syntax enable

" Setting filetype-specific stuffs
filetype plugin on
filetype indent on
filetype on

" Adding line numbers
set number

" Setting buffer around keys for scrolling
set so=4

" Disabling word wrapping
set nowrap

" Changing tab-width to 4
set tabstop=4
set shiftwidth=4
set noexpandtab

" Stopping vim from making those damn backups
set nobackup

" Changing colorscheme to molokai
set t_Co=256
colorscheme molokai

" Adding save/reload hotkeys
map r :so $MYVIMRC<Enter>
map s :w<Enter>

" GVim Specific Settings
if has("gui_running")
    set guioptions-=T
    set guioptions-=r
endif
