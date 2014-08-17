" Setting no compatible
set nocompatible

" Setting syntax highlighting
syntax enable

" Setting filetype-specific stuffs
filetype plugin on
filetype indent on
filetype on

" If filetype isn't work, use default autoindent
set autoindent

" Adding line numbers
set number

" Setting buffer around keys for scrolling
set so=4

" Disabling word wrapping
set nowrap

" Changing tab-width to 2
set tabstop=2
set shiftwidth=2
set smarttab

" Stopping vim from making backups and swaps
set nobackup
set noswapfile

" Changing colorscheme to molokai
set t_Co=256
colorscheme molokai

" Telling Vim to confirm instead of fail commands
set confirm

" Disabling modelines
set nomodeline

" Backspacing over line breaks and the such
set backspace=indent,eol,start

" Adding save/reload hotkeys
map < :tabp<Enter>
map > :tabn<Enter>
map <C-w> :w<Enter>
map <C-e> :q<Enter>
map <C-r> :wq<Enter>
map <C-t> :so $MYVIMRC<Enter>

" GVim Specific Settings
if has("gui_running")
  set guioptions-=T
  set guioptions-=r
endif
