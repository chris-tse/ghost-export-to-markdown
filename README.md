# ghost-export-to-markdown

A simple CLI to export Ghost posts to Markdown using the Ghost Content API.

## Installation

Run from NPM:

```bash
# Can alternatively use GHOST_URL and GHOST_API_KEY env vars
# Default export location is `ghost-export`
npx ghost-export-to-markdown --url=your.ghost.url --api-key=Your1Ghost2Api3Key --dir=path/to/export/dir
```

Build from source:

```bash
git clone https://github.com/chris-tse/ghost-export-to-markdown
go build -o bin/cli ./cli/main.go
```

## Development

Requirements:

- Go 1.24

```bash
go run ./cli/main.go
```
