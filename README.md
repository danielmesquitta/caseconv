# Case Conversion CLI

A fast and simple CLI tool for converting text between different case formats.

## Features

- Convert text to various case formats (camelCase, PascalCase, snake_case, kebab-case, etc.)
- Go-specific case conversions
- Supports both command-line arguments and stdin input (pipes)
- Built with Go using Cobra CLI framework and ettle/strcase library

## Installation

### Using Go Install

```bash
go install github.com/danielmesquitta/caseconv@latest
```

## Usage

### Basic Syntax

```bash
# Using command line arguments
caseconv <command> "your text here"

# Using pipes (stdin)
echo "your text here" | caseconv <command>
```

### Available Commands

| Command | Description | Example Input | Example Output |
|---------|-------------|---------------|----------------|
| `camel` | Convert to camelCase | "hello world" | "helloWorld" |
| `pascal` | Convert to PascalCase | "hello world" | "HelloWorld" |
| `snake` | Convert to snake_case | "HelloWorld" | "hello_world" |
| `kebab` | Convert to kebab-case | "HelloWorld" | "hello-world" |
| `uppersnake` | Convert to UPPER_SNAKE_CASE | "hello world" | "HELLO_WORLD" |
| `upperkebab` | Convert to UPPER-KEBAB-CASE | "hello world" | "HELLO-WORLD" |

### Go-Specific Commands

These commands provide Go-specific handling for identifiers and special cases:

| Command | Description | Use Case | Example Input | Example Output |
|---------|-------------|---------------|----------------|----------------|
| `gocamel` | Go-style camelCase | Go variable names | user id | userID
| `gopascal` | Go-style PascalCase | Go type/function names | user id | UserID

## Examples

### Command Line Arguments

```bash
# Convert to camelCase
$ caseconv camel "hello world example"
helloWorldExample

# Convert to PascalCase
$ caseconv pascal "hello world example"
HelloWorldExample

# Convert to snake_case
$ caseconv snake "HelloWorldExample"
hello_world_example

# Convert to kebab-case
$ caseconv kebab "HelloWorldExample"
hello-world-example

# Convert to UPPER_SNAKE_CASE
$ caseconv uppersnake "hello world example"
HELLO_WORLD_EXAMPLE
```

### Using Pipes

```bash
# From echo
$ echo "hello world example" | caseconv camel
helloWorldExample

# From file
$ cat file.txt | caseconv snake

# Chain with other commands
$ curl -s api.example.com/data | jq -r '.name' | caseconv pascal
```

### Multiple Words

```bash
# Multiple words as arguments
$ caseconv camel hello world example
helloWorldExample

# Quoted string (recommended)
$ caseconv camel "hello world example"
helloWorldExample
```

### Batch Processing

```bash
# Convert multiple strings
$ echo -e "hello world\nfoo bar\nbaz qux" | while read line; do caseconv pascal "$line"; done
HelloWorld
FooBar
BazQux
```

## Help

Get help for any command:

```bash
# General help
$ caseconv --help

# Help for specific command
$ caseconv camel --help
```

## Dependencies

- [github.com/spf13/cobra](https://github.com/spf13/cobra) - CLI framework
- [github.com/ettle/strcase](https://github.com/ettle/strcase) - Case conversion library

## License

This project is licensed under the terms specified in the LICENSE file.