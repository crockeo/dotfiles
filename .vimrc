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

" Setting the <Leader> to ' '
let mapleader=" "

" Disabling modelines
set nomodeline

" Backspacing over line breaks and the such
set backspace=indent,eol,start

" Moving around new buffers
map <C-Up> <C-w><Up>
map <C-Down> <C-w><Down>
map <C-Left> <C-w><Left>
map <C-Right> <C-w><Right>

" Moving around tabs
map <Leader><Left> :tabp<CR>
map <Leader><Right> :tabn<CR>

" Switching to the last-used buffer.
map ; :b#<CR>

map ' :tabe<CR>
map " :q<CR>

" Toggling NERDTree in the current tab.
map <C-n> :NERDTreeToggle<CR>

" GVim Specific Settings
if has("gui_running")
  set guioptions-=T
  set guioptions-=r
endif

" A line at column 81 to keep one from writing more than terminal width.
set colorcolumn=81

" Setting pathogen to work
execute pathogen#infect('~/.vim/bundle/{}')
