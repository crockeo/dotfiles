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
set so=100

" Disabling word wrapping
set nowrap

" Chaning the tab width to 4.
set tabstop=4
set shiftwidth=4

" Setting smart tabs
set smarttab

" Stopping vim from making backups and swaps
set nobackup
set noswapfile

" Changing colorscheme to molokai
set t_Co=256
colorscheme molokai

" Setting pathogen to work
execute pathogen#infect()

" Setting the <Leader> to ' '
let mapleader=" "

" Disabling modelines
set nomodeline

" Backspacing over line breaks and the such
set backspace=indent,eol,start

" Moving around new buffers
map <C-c> :vsp<Enter>
map <C-Up> <C-w><Up>
map <C-Down> <C-w><Down>
map <C-Left> <C-w><Left>
map <C-Right> <C-w><Right>

" Moving around tabs
map <Leader><Left> :tabp<Enter>
map <Leader><Right> :tabn<Enter>

" Saving a file
map <C-z> :w<Enter>

" Quitting a buffer
map <C-x> :q<Enter>

" Opening a new tab and opening NERDTree
map <C-n> :tabe<Enter>

" GVim Specific Settings
if has("gui_running")
  set guioptions-=T
  set guioptions-=r
endif

" Filetype associations
au BufRead,BufNewFile *.tpp set filetype=cpp
au BufRead,BufNewFile *.jade set filetype=jade
