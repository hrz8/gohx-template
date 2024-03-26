# gohx

Just a template for fullstack app using `Golang` + `htmx` + `templ`.

## Pre

### Install `templ`:
```bash
go install github.com/a-h/templ/cmd/templ@latest
```

### Install `task`
```bash
go install github.com/go-task/task/v3/cmd/task@latest
```

### Install `air`
```bash
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
