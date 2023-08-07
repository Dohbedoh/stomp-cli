# stomp CLI

A simple CLI to help troubleshoot connectivity to a STOMP server.

References:

* [STOMP Messaging Protocol](https://stomp.github.io/)
* [go-stomp](https://github.com/go-stomp/stomp)

## Build Instructions

Commonly:

```
CGO_ENABLED=0 go build -o build/stomp-cli
```

Prepend with `GOOS=[linux|darwin|windows]` and `GOARCH=[amd64|386]` to build for a different platform.

## Usage Instructions

```bash
% stomp-cli check connect [address] [flags]

Check connection to a STOMP server and returns the version

Usage:
  stomp check connect [flags]

Aliases:
  connect, conn

Flags:
  -c, --connectTimeout int   Connection timeout in seconds (optional) (default 10)
  -h, --help                 help for connect
  -p, --password string      Password (optional)
  -k, --skipTls              Skip TLS (optional)
  -t, --transport string     Transport type (default|tls). Default to 'tls' (default "tls")
  -u, --username string      Username (optional)
```

### Example

```bash
% stomp-cli check connect my.server:61617 --transport=tls --skipTls=true --username my-user --password my-password
{"level":"info","time":1691409749,"message":"1.2"}
```