# Emoji Fetcher
a simple tool for pulling the latest unicode emoji set

## Running
```
go run main.go
```
will create an updated `latest.json` inside `export/`

## Caveats
This pulls from the `unicode.org` ftp server, so dont use it a lot or they might shut that down, consider using this raw link instead.
```
https://raw.githubusercontent.com/Robindiddams/emoji-fetcher/master/export/latest.json?token=AB4ENWJH5DPFB2UPJTIFLXC5UDUEA
```
