# Pokédex CLI

An interactive command-line Pokédex built with Go and [PokéAPI](https://pokeapi.co/). Explore Pokémon locations, attempt to catch Pokémon, and inspect your collection from the terminal.

Built as part of the [Boot.dev](https://www.boot.dev/) curriculum to practice HTTP requests, JSON decoding, caching, concurrency, and testing in Go.

## Features

- **Explore locations:** Browse location areas with forward and backward pagination.
- **Discover Pokémon:** List the Pokémon found in a specific location.
- **Catch Pokémon:** Attempt catches with a probability based on each Pokémon’s base experience.
- **Inspect your collection:** View a caught Pokémon’s height, weight, stats, and types.
- **Cache API responses:** Reuse responses for locations and Pokémon to reduce repeated network requests.
- **Automatic cache cleanup:** Use a background goroutine and mutex-protected cache to periodically remove older entries.

## Requirements

- Go **1.26.2 or later**, as specified in `go.mod`.
- An internet connection for requests to PokéAPI.

No API key, database, or configuration file is required.

## Installation

Clone the repository:

```bash
git clone https://github.com/Ha0cH/pokedex.git
cd pokedex
```

Run the application:

```bash
go run .
```

Alternatively, build and run an executable:

```bash
go build -o pokedex .
./pokedex
```

On Windows:

```powershell
go build -o pokedex.exe .
.\pokedex.exe
```

## Usage

When the application starts, it displays an interactive prompt:

```text
Pokedex >
```

Enter commands at this prompt.

| Command | Description |
|---|---|
| `help` | Display the available commands. |
| `map` | Display the next page of location areas. |
| `mapb` | Display the previous page of location areas. |
| `explore <area-name>` | List Pokémon found in a location area. |
| `catch <pokemon-name>` | Attempt to catch a Pokémon. |
| `inspect <pokemon-name>` | Show details about a Pokémon you have caught. |
| `pokedex` | List all Pokémon caught during the current session. |
| `exit` | Close the application. |

Use hyphenated location names as returned by `map`, and a single space between a command and its argument.

### Example

Browse locations:

```text
map
```

Explore a location:

```text
explore pastoria-city-area
```

Attempt to catch a Pokémon:

```text
catch tentacool
```

Catches are probabilistic, so a Pokémon may escape. After a successful catch:

```text
inspect tentacool
pokedex
```

You can also attempt to catch a Pokémon directly by name without exploring a location first.

## How It Works

### API Requests

The application retrieves location and Pokémon data from PokéAPI, then decodes JSON responses into Go structs. HTTP requests use a 10-second timeout, and request errors are displayed in the terminal.

### Caching

API responses are cached in memory using their URLs as keys. Cache access is protected by a mutex because the command loop and background cleanup goroutine share the same data.

The application checks for old entries every five minutes and removes entries older than five minutes during cleanup. Expiration is handled by this periodic cleanup rather than checked on every read.

### Catching Pokémon

Catch probability decreases as a Pokémon’s base experience increases, with a minimum chance of 5%. Successfully caught Pokémon are stored in an in-memory collection, and duplicate catches are detected.

**Your collection and cached responses are not saved when you exit.** Each new session starts with an empty Pokédex.

## Tests

Run the tests:

```bash
go test ./...
```

The repository includes tests for:

- Input normalization and command tokenization.
- Adding and retrieving cache entries.
- Removing old cache entries through background cleanup.

To run tests with Go’s race detector:

```bash
go test -race ./...
```

## Project Structure

```text
pokedex/
├── main.go                  # Startup, command registration, and input loop
├── commands.go              # Command handlers and API requests
├── config.go                # Shared application configuration
├── json_types.go            # API response structures
├── repl.go                  # Input parsing and command dispatch
├── repl_test.go             # Input parsing tests
├── go.mod                   # Go module and version
└── internal/
    └── pokecache/
        ├── pokecache.go      # In-memory cache and background cleanup
        └── pokecache_test.go # Cache tests
```

## What I Practiced

- Calling REST APIs with Go’s `net/http` package.
- Decoding JSON into typed structs.
- Building an interactive command-line application.
- Using maps to manage commands, cached responses, and application state.
- Coordinating shared data access with goroutines and mutexes.
- Writing table-driven tests with Go’s `testing` package.

## Acknowledgments

- [Boot.dev](https://www.boot.dev/) for the project curriculum.
- [PokéAPI](https://pokeapi.co/) for the Pokémon data.