# egoproxy

tcp proxy tool written in golang

## How to install
``` bash
go install github.com/umaumax/egoproxy@latest
```

## How to run
``` bash
egoproxy
```

## help
``` bash
Usage of egoproxy:
  -ieproxy
    	detect the proxy settings on Windows platform (default true)
  -p string
    	HTTP proxy service address (e.g., ":1080") (default ":1080")
  -verbose
    	should every proxy request be logged to stdout (default true)
```

`-ieproxy`: windows only (WARN: unverified, TODO: verify this option)
