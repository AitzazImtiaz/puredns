# puredns

`puredns` is a fast and efficient DNS brute-forcing tool written in Go. It is designed to be highly concurrent and flexible, allowing users to customize their workflow.

## Features

- Fast and concurrent DNS brute-forcing
- Wildcard detection and filtering
- Customizable resolvers and wordlists
- Multiple output formats (txt, json)
- Colored terminal output

## Installation

To install `puredns`, you need to have Go installed on your system. You can then build the project from source:

```bash
git clone https://github.com/aitzazimtiaz/puredns.git
cd puredns
./scripts/build.sh
```

## Usage

```bash
./puredns -d <domain> -w <wordlist> [options]
```

### Options

- `-d`: Target domain (required)
- `-w`: Path to wordlist file (required)
- `-r`: Path to resolvers file (default: `configs/resolvers.txt`)
- `-o`: Output file path
- `-q`: Quiet mode (no output to stdout)
- `-wildcard`: Enable wildcard detection (default: true)
- `-h`: Show help message
