# gcrt

crt is a simple in cli tool to check certificate information of certain domains against `https://gcrt.sh/`.

## usage
```
gcrt is a tool to query the Certificate Transparency Logs
                   it does so by querying https://crt.sh
                   Complete documentation is available at https://github.com/jhinds/gcrt

Usage:
  gcrt [flags]

Flags:
      --between string   The dates to run the query for in the format start-date:end-date.  The dates should have the format YYYY-MM-DD
      --count string     Don't return the results just the count
      --days string      How many days back to query
      --domain string    Domain to find certificates for. % is a wildcard
  -h, --help             help for gcrt
```

## to build
`go build -o bin/gcrt`

## to download
`go get -u github.com/jhinds/gcrt`