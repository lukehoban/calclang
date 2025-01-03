# Calclang

A simple calculator language that supports basic arithmetic operations.

## Usage

Calclang supports both a CLI interface and a REST API server.

### CLI Usage

To run the CLI version:

```bash
go run cmd/cli/main.go
```

Then enter expressions in the format:
```
ADD 1 2
SUB 5 3
```

### REST API Usage

To run the server:

```bash
go run cmd/server/main.go
```

The server will start on port 8080. You can then make POST requests to `/calculate` with a JSON body:

```bash
curl -X POST http://localhost:8080/calculate \
  -H "Content-Type: application/json" \
  -d '{"expression": "ADD 1 2"}'
```

Response format:
```json
{
  "result": 3
}
```

Or in case of an error:
```json
{
  "error": "Invalid expression"
}
```

## Supported Operations

- `ADD x y`: Adds two numbers
- `SUB x y`: Subtracts y from x