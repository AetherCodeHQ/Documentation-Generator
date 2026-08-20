# Documentation Generator

![CI](https://github.com/Qyroxen/Documentation-Generator/actions/workflows/ci.yml/badge.svg)
![CodeQL](https://github.com/Qyroxen/Documentation-Generator/actions/workflows/codeql.yml/badge.svg)
![Go](https://img.shields.io/badge/Go-1.23+-00ADD8?style=flat&logo=go)
![License](https://img.shields.io/badge/License-MIT-yellow.svg)
![Stars](https://img.shields.io/github/stars/Qyroxen/Documentation-Generator?style=social)
![Issues](https://img.shields.io/github/issues/Qyroxen/Documentation-Generator)
![PRs](https://img.shields.io/github/issues-pr/Qyroxen/Documentation-Generator)

> A production-ready CLI tool built with Go

[![Star Badge](https://img.shields.io/github/stars/Qyroxen/Documentation-Generator?style=social)](https://github.com/Qyroxen/Documentation-Generator/stargazers)

## What is it?

Documentation Generator is a production-ready CLI tool built with Go. It provides powerful functionality with a beautiful terminal interface.

## Features

- Fast and efficient (written in Go)
- Beautiful CLI with colored output
- Comprehensive documentation
- GitHub Actions CI/CD
- CodeQL security analysis
- Dependabot for dependency updates
- MIT Licensed
- Fully offline - zero cloud dependency

## Quick Start

```bash
# Install
git clone https://github.com/Qyroxen/Documentation-Generator.git
cd Documentation-Generator
go build -o documentationgenerator .

# Run
./documentationgenerator --help
```

## CLI Usage

```bash
# Basic usage
./documentationgenerator

# With flags
./documentationgenerator --verbose --output json

# Get help
./documentationgenerator --help
```

## Examples

```bash
# Example 1
./documentationgenerator example1

# Example 2
./documentationgenerator example2 --flag value
```

## Development

```bash
# Run tests
go test ./...

# Build
go build -o documentationgenerator .

# Lint
golangci-lint run

# Security scan
codeql analyze
```

## Contributing

Contributions are welcome! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for details.

## Security

For security vulnerabilities, please see [SECURITY.md](SECURITY.md).

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

<p align="center">
  <a href="https://github.com/Qyroxen/Documentation-Generator/stargazers">
    <img src="https://img.shields.io/github/stars/Qyroxen/Documentation-Generator?style=social" alt="Star this repo">
  </a>
  <a href="https://github.com/Qyroxen/Documentation-Generator/forks">
    <img src="https://img.shields.io/github/forks/Qyroxen/Documentation-Generator?style=social" alt="Fork this repo">
  </a>
  <a href="https://github.com/Qyroxen/Documentation-Generator/issues">
    <img src="https://img.shields.io/github/issues/Qyroxen/Documentation-Generator" alt="Issues">
  </a>
  <a href="https://github.com/Qyroxen/Documentation-Generator/pulls">
    <img src="https://img.shields.io/github/issues-pr/Qyroxen/Documentation-Generator" alt="Pull Requests">
  </a>
</p>
