# Calc Lang

A simple calculator language that supports basic arithmetic operations.

## Features

- Basic arithmetic operations (ADD, SUB)
- Available as both CLI and REST API
- Simple expression syntax

## Building

```bash
# Build CLI
go build -o calc-cli ./cmd/cli

# Build Server
go build -o calc-server ./cmd/server
```

## Using the CLI

Run the CLI:

```bash
./calc-cli
```

Example usage:
```
> ADD 5 3
8
> SUB 10 4
6
```

## Using the REST API

Start the server:

```bash
./calc-server
```

The server will start on port 8080.

Make calculations using HTTP POST requests:

```bash
curl -X POST http://localhost:8080/calculate \
  -H "Content-Type: application/json" \
  -d '{"expression": "ADD 5 3"}'
```

Response:
```json
{"value": 8}
```

Error response:
```json
{"error": "invalid expression format"}
```

## Expression Syntax

Expressions follow the format:
```
OPERATION NUMBER NUMBER
```

Supported operations:
- ADD: Adds two numbers
- SUB: Subtracts the second number from the first