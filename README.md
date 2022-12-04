# Advent of Code 2022 🕯️🕯️🕯️🕯️ 
[![Go Unit Tests](https://github.com/capthiron/advent-of-code-2022/actions/workflows/go.yml/badge.svg)](https://github.com/capthiron/advent-of-code-2022/actions/workflows/go.yml)

My personal, sloppy take on the [Advent of Code 2022](https://adventofcode.com/) puzzles 🫠

## Run and have fun 🦌

Clone the repository if you haven't done so
*via https*
```bash
git clone https://github.com/capthiron/advent-of-code-2022.git
```
*via ssh*
```bash
git clone git@github.com:capthiron/advent-of-code-2022.git
```

Run solution of days 01-25 by executing
```bash
go run cmd/DAY/main.go --fileName "path/to/input/file"
```

or with encrypted input file

```bash
go run cmd/DAY/main.go --fileName "path/to/input/file" --key "KEY"
```

### How to bootstrap a new day 🎁

```bash
go run cmd/bootstrap/main.go --day [01-25]
```

### How to encrypt input files 🔒

```bash
go run cmd/crypto/main.go --key "KEY" --fileName "path/to/input/file.txt" > path/to/encrypted/input/file.txt
```
