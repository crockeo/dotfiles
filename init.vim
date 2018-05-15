" Disabling backwards compatibility.
set nocompatible

" Setting up Plug.
call plug#begin('~/.local/share/nvim/plugged')

Plug 'scrooloose/nerdcommenter'
Plug 'godlygeek/tabular'
Plug 'airblade/vim-gitgutter'
Plug 'kien/ctrlp.vim'
Plug 'rust-lang/rust.vim'
Plug 'tomasr/molokai'

call plug#end()

let g:ctrlp_map = '<C-p>'
let g:ctrlp_user_command = ['.git/', 'git --git-dir=%s/.git ls-files -oc --exclude-standard']

" Enabling filetype-based functionality.
filetype plugin on
filetype indent on
filetype on

" Enabling syntax highlighting.
syntax on

" Removing the backup file and swap files.
set nobackup
set noswapfile

" Adding line numbers
set number

" Adding a buffer around the cursor when scrolling.
set so=5

" Disabling word wrapping
set nowrap

" Chaning the tab width to 4.
set expandtab
set tabstop=4
set softtabstop=4
set shiftwidth=4

" A line at column 81 to keep one from writing more than terminal width.
set colorcolumn=81

" Allowing you to backspace in a close-to-sane way.
set backspace=indent,eol,start

" Mapping the leader to space.
let mapleader=" "

" Moving around tabs
map <Leader><Left> :tabp<CR>
map <Leader><Right> :tabn<CR>

" Maping another way to exit insert mode.
imap <C-f> <ESC>

" Moving to the back and front of a line, respectively.
map <C-a> <Home>
map <C-e> <End>
imap <C-a> <Home>
imap <C-e> <End>

" Switching to the last-used buffer.
map ; :b#<CR>

" Disabling auto-commenting the next line.
autocmd FileType * setlocal formatoptions-=c formatoptions-=r formatoptions-=o

" Disabling folding of functions.
set nofoldenable

" Changing colorscheme to molokai if it exists. Otherwise using the slate color
" scheme.
let g:molokai_original = 1
set t_Co=256
try
    colorscheme molokai
catch /^Vim\%((\a\+)\)\=:E185/
    colorscheme slate
endtry
