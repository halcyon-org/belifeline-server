# Kizuna

## Tasks

> [!NOTE]
> You can use `xc`(<https://xcfile.dev/>) to run the commands
> See <https://xcfile.dev/getting-started/#installation> for installation instructions

### init

Init development environment.

Run: once

```bash
echo "init"
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
