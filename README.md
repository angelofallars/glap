# glap

Glap is an incomplete rewrite of grep in Go. It can find lines that match a pattern from standard input or from mulptile files. Currently supports a few options.

Available options:

- `--h`, `--help`                - show help message and exit

-  `-i`, `--ignore-case`         - ignore case when finding matches

-  `-n`, `--line-number`         - print line number before matching lines

-  `-c`, `--count`               - only display the count of matching lines

-  `-v`, `--invert-match`        - display non-matching lines instead

-  `-m`, `--max-count [N]`           - stop after N selected lines

### License

MIT
