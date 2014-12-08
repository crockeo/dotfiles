# vimconfig

This repo contains my configuration for vim.

### Wanna Install It?

Don't wanna think about anything? Just wanna install it? Download the repo and
run `install.sh`:

```bash
$ git clone https://github.com/crockeo/vimconfig.git
$ cd vimconfig/
$ ./install.sh
```

(Just so you know this will move your current vim files to `vim_old`.)

### Script Actions

* `update-repo.sh` updates the repository with the current system files.
* `update-system.sh` changes the system `.vim` and `.vimrc` to `.vim_old/` and
`.vimrc_old`, and then moves the `.vim/` and `.vimrc` from this repo to the `~/`
folder.
* `clean-system.sh` removes the `.vim_old/` and `.vimrc_old` files from the `~/`
folder. **This can result in a loss of your own vim config if not used
carefully!**
