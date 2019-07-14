set nocompatible

"""""""""""""""""""""""
" Plugin Installation "
call plug#begin('~/.local/share/nvim/plugged')

" General-use plugins
Plug 'scrooloose/nerdcommenter'
Plug 'godlygeek/tabular'
Plug 'airblade/vim-gitgutter'
Plug 'kien/ctrlp.vim'
Plug 'tomasr/molokai'
Plug 'roxma/nvim-yarp'
Plug 'ncm2/ncm2'
Plug 'ncm2/ncm2-bufword'
Plug 'ncm2/ncm2-path'

" Code formatting tooling
Plug 'google/vim-maktaba'
Plug 'google/vim-codefmt'
Plug 'google/vim-glaive'
Plug 'syml/rust-codefmt'
Plug 'ambv/black'

" C++ IDE Support
Plug 'Rip-Rip/clang_complete'

" Golang IDE Support
Plug 'fatih/vim-go'
Plug 'ncm2/ncm2-go'

" Python IDE Support
Plug 'ncm2/ncm2-jedi'

" Rust IDE Support
Plug 'rust-lang/rust.vim'

call plug#end()

""""""""""""""""""
" General config "
filetype plugin indent on
syntax on
set nobackup
set noswapfile
set number
set so=5
set nowrap
set expandtab
set tabstop=2
set softtabstop=2
set shiftwidth=2
set colorcolumn=121
set textwidth=120
set backspace=indent,eol,start
let mapleader=" "
imap <C-f> <ESC>
map <C-a> <Home>
map <C-e> <End>
imap <C-a> <Home>
imap <C-e> <End>
map ; :b#<CR>
map <Leader><Left> :tabp<CR>
map <Leader><Right> :tabn<CR>
autocmd FileType * setlocal formatoptions-=c formatoptions-=r formatoptions-=o
set nofoldenable

""""""""""""""""""""""""
" Plugin Configuration "

" CtrlP Configuration
let g:ctrlp_map = '<C-p>'
let g:ctrlp_user_command = ['.git/', 'git --git-dir=%s/.git ls-files -oc --exclude-standard']

" Black Configuration
let g:black_linelength=120
let g:black_skip_string_normalization=1

" General IDE Configuration
autocmd BufEnter * call ncm2#enable_for_buffer()
set completeopt=menuone,noselect,noinsert
set shortmess+=c
let ncm2#popup_delay = 5
let ncm2#complete_length = [[1, 1]]
let g:ncm2#matcher = 'substrfuzzy'

" C++ IDE Configuration
let g:clang_library_path='/Library/Developer/CommandLineTools/usr/lib'
let g:clang_format#command='/usr/local/opt/llvm@7/bin/clang-format'
autocmd BufEnter *.go call ncm2#override_source('bufword', {'scope_blacklist': ["golang"]})

" Python IDE Configuration
autocmd BufEnter *.py call ncm2#override_source('bufword', {'scope_blacklist': ["python"]})

" Color Scheme Configuration
let g:molokai_original = 1
set t_Co=256
try
    colorscheme molokai
catch /^Vim\%((\a\+)\)\=:E185/
    colorscheme slate
endtry

" Formatting Configuration
function! DynFormat()
  let buf_ft = &filetype
  if buf_ft ==# "python"
    Black
  else
    FormatCode
  endif
endfunction

map <C-f> :call DynFormat()<CR>
