# crt

crt is a simple in cli tool to check certificate information of certain domains against `https://crt.sh/`.

## usage
```
crt is a tool to query the Certificate Transparency Logs
                   it does so by querying https://crt.sh
                   Complete documentation is available at https://github.com/jhinds/crt

Usage:
  crt [flags]

Flags:
      --between string   The dates to run the query for in the format start-date:end-date.  The dates should have the format YYYY-MM-DD
      --count string     Don't return the results just the count
      --days string      How many days back to query
      --domain string    Domain to find certificates for. % is a wildcard
  -h, --help             help for crt
```

## to build
`go build -o bin/crt`