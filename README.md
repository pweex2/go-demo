# go-demo

A minimal Go web service demo with GitHub Actions CI.

## 🔧 Features

- Simple HTTP server (`/hello`)
- Unit tests
- CI with GitHub Actions
- Artifact upload (compiled binary)

## 🚀 Quick Start

```bash
go run main.go
curl http://localhost:8080/hello
```

## 🧪 Run Tests

```bash
go test ./...
```

## 📦 Build

```bash
go build -o bin/go-demo
```

## 🛠 Makefile Build (Recommended)

You can also use the provided Makefile to simplify builds:

```bash
make build
```

The Makefile will auto-detect your OS:

- On macOS: builds a native binary
- On Linux (e.g. CI or container): builds a Linux-compatible binary

The result is placed in `bin/go-demo`.

```bash
make build-mac
```

can be used to build a macOS-compatible binary.

```bash
make build-linux
```

can be used to build a Linux-compatible binary.

## 🛰 Remote Deployment Placeholder

While you don't yet have a live server for deployment, the project includes a placeholder `Makefile` target to simulate remote deployment:

```bash
make deploy-remote
```

This command will print out the typical steps you'd use with `scp`, `ssh`, or `rsync`, such as:

```bash
scp -i ~/.ssh/id_rsa bin/go-demo-linux user@remote:/usr/local/bin/go-demo
ssh -i ~/.ssh/id_rsa user@remote 'pkill go-demo || true && nohup /usr/local/bin/go-demo &'
rsync -avz -e 'ssh -i ~/.ssh/id_rsa' bin/go-demo-linux user@remote:/usr/local/bin/go-demo
```

Once a real server is available, this can be turned into an actual deployment step in your CI/CD pipeline.
