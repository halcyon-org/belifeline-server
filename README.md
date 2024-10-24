# Kizuna

## Dev dependencies

- [Docker](https://www.docker.com/)
- [Go](https://go.dev//)
- [xc](https://xcfile.dev/)
- [golangci-lint](https://golangci-lint.run/)
- [sqlfluff](https://www.sqlfluff.com/)

## Tasks

> [!NOTE]
> You can use `xc`(<https://xcfile.dev/>) to run the commands
>
> See <https://xcfile.dev/getting-started/#installation> for installation instructions

### install

Install the dependencies.

Run: once

```bash
go mod tidy
go install golang.org/x/tools/go/analysis/passes/defers/cmd/defers@latest
go install golang.org/x/tools/go/analysis/passes/fieldalignment/cmd/fieldalignment@latest
go install golang.org/x/tools/go/analysis/passes/findcall/cmd/findcall@latest
go install golang.org/x/tools/go/analysis/passes/nilness/cmd/nilness@latest
go install golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow@latest
```

### init

Init development environment.

Requires: install
Run: once

```bash
if [ ! -f .env ]; then
  cp .env.example .env
fi
```

### check

Check golang code.

Requires: init, fmt, vet, lint, test
RunDeps: async

### fmt

Format golang code.

Requires: init

```bash
go fmt ./...
```

### lint

Lint golang code.

Requires: init

```bash
golangci-lint run
```

### vet

Vet golang code.

Requires: init

```bash
go vet ./...
go vet -vettool=$(command -v defers) ./...
go vet -vettool=$(command -v fieldalignment) ./...
go vet -vettool=$(command -v nilness) ./...
go vet -vettool=$(command -v shadow) ./...
```

### test

Test golang code.

Requires: init

```bash
go test ./...
```

### build

Build golang code.

Requires: init, gen

```bash
go build -o bin/kizuna cmd/kizuna/main.go
```

### gen

Generate golang code.

Requires: init

```bash
go generate ./...
```

### up

Up the development environment.

Requires: init

```bash
docker compose up -d
```

### migrate

Dry run the migrations.

Requires: init, up

```bash
go run cmd/migration/main.go --dry | sqlfluff fix - --dialect postgres
```

### migrate:apply

Apply the migrations.

Requires: init, up

```bash
go run cmd/migration/main.go
```

### down

Down the development environment.

Requires: init

```bash
docker compose down
```

### down:all

Down the development environment and remove volumes.

Requires: init

```bash
docker compose down -v
```

### run

Run the development environment.

Requires: init, up, migrate:apply

```bash
go run cmd/kizuna/main.go
```

### cli

Run the cli.

Requires: init, up, migrate:apply

```bash
echo "RUN: go run cmd/kizuna-cli/main.go"
```
