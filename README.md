# Brag

## Description

Brag is a command-line interface (CLI) tool that allows you to write and manage entries in a brag document. Use it to keep track of your accomplishments, milestones, or anything you want to brag about!

## Installation

1. Make sure you have Go installed on your system. You can download it from the official Go website: https://golang.org/

2. Clone the brag repository to your local machine:
`git clone [https://github.com/BrunoPansani/brag.git](https://github.com/BrunoPansani/brag.git)`

3. Navigate to the project directory:
`cd brag`

4. Build the Go executable:
`go build ./cmd/brag`

5. (Optional) Add the `brag` executable to your system's PATH to run it from anywhere.
   1. In the terminal, run `pwd` to get the current directory.
   2. Use a text editor to open your shell configuration file. For macOS, it is typically `~/.bash_profile` or `~/.zshrc`. For Linux, it can be `~/.bashrc` or `~/.zshrc`.
   3. Add the following line to the configuration file, replacing `/path/to/brag` with the actual path to the brag executable:
`export PATH="/path/to/brag:$PATH"`
   4. Run `source ~/.bash_profile` (or the relevant configuration file) to apply the changes to the current session. Alternatively, you can close and reopen the terminal window.
   5. Type brag in the terminal to verify that the command is recognized and executable from anywhere.

## Usage

The `brag` CLI provides several commands to manage your brag document. Here are the available commands:

- `init`: Initializes the brag document.
- `add <entry>`: Adds a new entry to the brag document.
- `list`: Lists all entries in the brag document.
- `remove <id>`: Removes the entry with the specified ID.
- `clear`: Clears all entries from the brag document.
- `export <format>`: Exports the brag document to the specified format (txt, csv, json).
- `help`: Displays help information.

To run a command, open a terminal or command prompt and navigate to the project directory.

Here are some examples of how to use the `brag` CLI:

- Initialize the brag document:

`./brag init`

- Add a new entry to the brag document:`

`./brag add "I completed a challenging project today."`

- List all entries in the brag document:

`./brag list`

- Remove an entry from the brag document:

`./brag remove 1`

- Clear all entries from the brag document:

`./brag clear`

- Export the brag document to a specific format (txt, csv, json):

`./brag export txt`

`./brag export csv`

`./brag export json`

- Display help information:

`./brag help`

Note: If you added the `brag` executable to your system's PATH, you can simply use `brag` instead of `./brag` in the above commands.

That's it! You're now ready to start using the `brag` CLI to manage your brag document.

