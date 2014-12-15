" Setting no compatible
set nocompatible

" Making sure to turn the filetype detection off (per the Vundle docs'
" request)
filetype off

" Setting up Vundle
set rtp+=~/.vim/bundle/Vundle.vim
call vundle#begin()

" Setting up plugins.
Plugin 'gmarik/Vundle.vim'

Plugin 'scrooloose/nerdtree'
Plugin 'scrooloose/nerdcommenter'
Plugin 'jistr/vim-nerdtree-tabs'
Plugin 'tikhomirov/vim-glsl'
Plugin 'fatih/vim-go'
Plugin 'digitaltoad/vim-jade'
Plugin 'godlygeek/tabular'
Plugin 'wlangstroth/vim-racket'

" Done with Vundle setup
call vundle#end()

" Setting filetype-specific stuffs
filetype plugin on
filetype indent on
filetype on

" Setting syntax highlighting
syntax enable

" If filetype isn't work, use default autoindent
set autoindent

" Adding line numbers
set number

" Setting buffer around keys for scrolling
set so=100

" Disabling word wrapping
set nowrap

" Chaning the tab width to 4.
set expandtab
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
map <S-Up> <C-w><Up>
map <S-Down> <C-w><Down>
map <S-Left> <C-w><Left>
map <S-Right> <C-w><Right>

" Moving around tabs
map <Leader><Left> :tabp<CR>
map <Leader><Right> :tabn<CR>

" Maping another way to exit insert mode.
imap <C-f> <ESC>

" Reloading the .vimrc
map <Leader><Leader> :source $MYVIMRC<CR>

" Switching to the last-used buffer.
map ; :b#<CR>

map ' :tabe<CR>
map " :q<CR>

" Toggling NERDTree in the current tab.
map <C-n> :NERDTreeTabsToggle<CR>

" GVim Specific Settings
if has("gui_running")
  set guioptions-=T
  set guioptions-=r
endif

" A line at column 81 to keep one from writing more than terminal width.
set colorcolumn=81

" Setting the original molokai theme.
let g:molokai_original = 1
