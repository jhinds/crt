---
title: "Home"
---

gcrt is a simple in cli tool to check certificate information of certain domains against [`https://crt.sh/`](https://crt.sh/).

## usage
```
gcrt is a tool to query the Certificate Transparency Logs
     it does so by querying https://crt.sh
     Complete documentation is available at https://github.com/jhinds/gcrt
     Homepage https://gcrt.jhinds.dev

Usage:
  gcrt [flags]

Flags:
      --between string   The dates to run the query for in the format start-date:end-date.  The dates should have the format YYYY-MM-DD
  -c, --count string     Don't return the results just the count
      --days string      How many days back to query
  -d, --domain string    Domain to find certificates for. % is a wildcard
  -h, --help             help for gcrt
```

# examples

### see all certs for domain
```bash
> gcrt --domain 'hinds.io'
[
  ...
  {
    "issuer_ca_id": 50556,
    "issuer_name": "C=US, O=Google Trust Services, CN=GTS CA 1D2",
    "common_name": "www.st4dium.com",
    "name_value": "hinds.io",
    "id": 3485418628,
    "entry_timestamp": "2020-10-09T10:47:40.771Z",
    "not_before": "2020-10-05T22:19:35Z",
    "not_after": "2021-01-03T22:19:35Z",
    "serial_number": "7a2db2a7c785388f09000000003328b1"
  },
  ...
]
```

### see all certs for domain for past 60 days
```bash
> gcrt --domain 'hinds.io' --days 60
  ...
  {
    "issuer_ca_id": 50556,
    "issuer_name": "C=US, O=Google Trust Services, CN=GTS CA 1D2",
    "common_name": "www.st4dium.com",
    "name_value": "hinds.io",
    "id": 3485418628,
    "entry_timestamp": "2020-10-09T10:47:40.771Z",
    "not_before": "2020-10-05T22:19:35Z",
    "not_after": "2021-01-03T22:19:35Z",
    "serial_number": "7a2db2a7c785388f09000000003328b1"
  },
  ...
]
```

### print output to as text
```bash
> gcrt --domain 'hinds.io' --days 60 --output text
   COMMONNAME   	    SANS    	           CREATED            	            ISSUER            	            START            	             END             	EXPIRESIN
www.st4dium.com 	hinds.io    	2020-10-09 10:47:40.771 +0000 	C=US, O=Google Trust Services,	2021-01-03 22:19:35 +0000 UTC	2020-10-05 22:19:35 +0000 UTC	35 days
                	            	UTC                           	CN=GTS CA 1D2
thosta-group.com	www.hinds.io	2020-10-09 09:55:50.513 +0000 	C=US, O=Google Trust Services,	2021-01-04 03:53:14 +0000 UTC	2020-10-06 03:53:14 +0000 UTC	35 days
                	            	UTC                           	CN=GTS CA 1D2
thosta-group.com	www.hinds.io	2020-10-06 04:53:15.695 +0000 	C=US, O=Google Trust Services,	2021-01-04 03:53:14 +0000 UTC	2020-10-06 03:53:14 +0000 UTC	35 days
                	            	UTC                           	CN=GTS CA 1D2
www.st4dium.com 	hinds.io    	2020-10-05 23:19:37.326 +0000 	C=US, O=Google Trust Services,	2021-01-03 22:19:35 +0000 UTC	2020-10-05 22:19:35 +0000 UTC	35 days
                	            	UTC                           	CN=GTS CA 1D2
```