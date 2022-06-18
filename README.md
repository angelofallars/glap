# glap

Glap is an incomplete rewrite of grep in Go. It can find lines that match a pattern from standard input or from multiple files. It currently supports a few options.

Available options:

- `--h`, `--help`                - show help message and exit

-  `-i`, `--ignore-case`         - ignore case when finding matches

-  `-n`, `--line-number`         - print line number before matching lines

-  `-c`, `--count`               - only display the count of matching lines

-  `-v`, `--invert-match`        - display non-matching lines instead

-  `-m`, `--max-count [N]`           - stop after N selected lines

### Example usage

```bash
$ du -h --all --max-depth 1
4.5M    ./.git
4.0K    ./LICENSE
4.0K    ./go.sum
4.0K    ./utils
4.0K    ./go.mod
4.0K    ./.gitignore
8.0K    ./main.go
2.2M    ./glap
4.0K    ./README.md
6.7M    .

$ du -h --all --max-depth 1 | glap go
4.0K    ./go.sum
4.0K    ./go.mod
8.0K    ./main.go
```

### License

MIT
