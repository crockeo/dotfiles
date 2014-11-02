# vimconfig

This repo contains my configuration for vim.

### Script Actions

* `update-repo.sh` updates the repository with the current system files.
* `update-system.sh` changes the system `.vim` and `.vimrc` to `.vim_old/` and
`.vimrc_old`, and then moves the `.vim/` and `.vimrc` from this repo to the `~/`
folder.
* `clean-system.sh` removes the `.vim_old/` and `.vimrc_old` files from the `~/`
folder. **This can result in a loss of your own vim config if not used
carefully!**
