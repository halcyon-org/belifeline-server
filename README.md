# Kizuna

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

### down

Down the development environment.

Requires: init

```bash
docker compose down
```

### run

Run the development environment.

Requires: init, up

```bash
go run cmd/kizuna/main.go
```
