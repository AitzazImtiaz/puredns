# Architecture

The `puredns` project is organized into several modules, each with a specific responsibility.

- **cmd**: Contains the main application entry point and CLI flag handling.
- **internal**: Core logic of the application.
  - **brute**: Main brute-force logic and concurrency engine.
  - **dns**: DNS resolution, wildcard detection, and record helpers.
  - **input**: Wordlist and resolver parsing.
  - **output**: Output writers, terminal printer, and result formatter.
  - **utils**: Logging and other utility functions.
- **configs**: Default configuration files.
- **scripts**: Build, benchmark, and test scripts.
- **test**: Unit and integration tests.

## Data Flow

1. The `main` function in `cmd/puredns/main.go` parses the command-line arguments.
2. The `brute.Run` function is called, which reads the wordlist and resolvers.
3. Wildcard detection is performed by the `dns.DetectWildcard` function.
4. A pool of `BruteForceWorker` goroutines is created to perform the DNS lookups.
5. The workers resolve subdomains using the `dns.Resolve` function.
6. Results are filtered to remove wildcard IPs and sent to the results channel.
7. The `main` function receives the results and prints them to the console and/or writes them to a file.
