# BASH HISTORY PARSER

This tool takes a bash history file (e.g .bash_history, .zsh_history) as input and tries to identify the working directory where commands have been run.

## Why

Usually, after exploiting a path traversal vulnerability, can be useful to inspect the bash history of the users in the system. This parser will make your life easier, and tries to identify paths inside it.

## Usage

```
Usage:
  bash-history-parser [flags]

Flags:
  -d, --dir string    Starting dir, usually the user home directory (required) (default "home/parser")
  -f, --file string   .bash_history file to parse (required)
  -h, --help          help for bash-history-parser
```

## Limitations

It is not possible to identify the right path when the user creates a new terminal session, a result if the user uses only relative ```cd``` commands it will not possible to identify the right path.

Example

Home: ```/home/parser```

File: ```.bash_history```
```bash
cd mydir   --> /home/parser/mydir        [CORRECT]
# user starts a new terminal session
cd mydir   --> /home/parser/mydir/mydir  [INCORRECT - new session cannot be identified]
# user starts a new terminal session
cd ~/mydir --> /home/parser/mydir        [CORRECT]
```
