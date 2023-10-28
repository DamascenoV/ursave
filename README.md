# UrSave
## Overview
UrlSave is a command-line tool developed in Go that allows you to save and manage your favorite URLs directly from your terminal. It provides a convenient way to store, list and open URLs, making it easy to access your favorite websites.

## Features
- Save URLs with a custom alias for easy access.
- List all saved URLs.
- Open a saved URL in your default browser.
- Delete sabed URLs.
- Edit saved URLs.
- Search for saved URLs with fuzzy-finder.

## Prerequisites
- Go 1.18 or higher

## Usage
To use the UrSave tool in the Terminal, you can run the following commands:

- To add a Url:
```bash
$ ursave add -n <name/alias> -u <url>
```
- To list all saved urls with fzf:
```bash
$ ursave list
```
- To open a saved url:
```bash
$ ursave -o <alias>
```
- To delete a saved url:
```bash
$ ursave delete -n <name/alias>
```
- To edit a saved url:
```bash
$ ursave edit -n <name/alias>
```
## License
This project is licensed under the MIT License

## Support
If you encounter any issues or have questions, please open an issue in the GitHub repository.

Enjoy saving and managing your favorite URLs from the terminal!
