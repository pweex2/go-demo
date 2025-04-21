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

## 📦 Releasing a Version

This project supports automatic GitHub Releases. When you push a tag like `v1.0.x`, GitHub Actions will:

1. Build the binary
2. Create a GitHub Release
3. Upload the `bin/go-demo` binary to the Release assets

### Example:

```bash
git tag v1.0.x
git push origin v1.0.x
```

This will trigger the release workflow automatically.

## ✅ Linting

This project uses [golangci-lint](https://golangci-lint.run/) for static code analysis and style checks.

You can run linting locally:

```bash
make lint
```

Or let GitHub Actions automatically run it on every commit/pull request:

- Checks naming, unused code, errors, formatting, etc.
- Fails CI if issues are detected
