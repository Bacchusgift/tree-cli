# Tree CLI

A command-line tool written in Go for displaying directory structures in a tree-like format, similar to the Unix `tree` command.

## Features
- Display directory trees with customizable depth
- Option to show hidden files
- List directories only
- Show full file paths
- Flat (no-indent) display mode
- Output results to a file
- Colored output for better readability

## Command Line Options
- `-L` Set the maximum display depth
- `-a` Show all files, including hidden files
- `-d` List directories only
- `-f` Show full paths
- `-i` Flat display (no indentation)
- `-o` Output to file

## Usage
```sh
# Display the current directory tree
./tree-cli

# Display up to 2 levels deep
./tree-cli -L 2

# Show all files including hidden ones
./tree-cli -a

# Only list directories
./tree-cli -d

# Show full paths
./tree-cli -f

# Flat display (no indentation)
./tree-cli -i

# Output to a file
./tree-cli -o output.txt
```

## Build
Make sure you have Go installed, then run:
```sh
go build -o tree-cli main.go
```

## License
MIT
