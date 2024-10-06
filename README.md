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

### init

Init development environment.

Run: once

```bash
if [ ! -f .env ]; then
  cp .env.example .env
fi
```

### check

Check golang code.

Requires: init, fmt, vet, test
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
