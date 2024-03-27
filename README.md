# gohx

Just a template for Fullstack Web App using `Golang` + `htmx` + `templ` with SPA-like router.

## Tools

```bash
# Install sqlc
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
# Install templ
go install github.com/a-h/templ/cmd/templ@latest
# Install task
go install github.com/go-task/task/v3/cmd/task@latest
# Install air
go install github.com/cosmtrek/air@latest
```

## Setup

### Install deps
```bash
go mod tidy
yarn install
```

### Build templ
```bash
templ generate
```

### Build assets
```bash
node esbuild.mjs
```

## Dev
```bash
task watch
task watch:fe
```
