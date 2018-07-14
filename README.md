# ship

[![Travis CI](https://travis-ci.org/jessfraz/ship.svg?branch=master)](https://travis-ci.org/jessfraz/ship)

Command line tool to track packages using the 
[AfterShip API](https://docs.aftership.com/api/4/overview).

## Installation

#### Binaries

- **darwin** [386](https://github.com/jessfraz/ship/releases/download/v0.0.0/ship-darwin-386) / [amd64](https://github.com/jessfraz/ship/releases/download/v0.0.0/ship-darwin-amd64)
- **freebsd** [386](https://github.com/jessfraz/ship/releases/download/v0.0.0/ship-freebsd-386) / [amd64](https://github.com/jessfraz/ship/releases/download/v0.0.0/ship-freebsd-amd64)
- **linux** [386](https://github.com/jessfraz/ship/releases/download/v0.0.0/ship-linux-386) / [amd64](https://github.com/jessfraz/ship/releases/download/v0.0.0/ship-linux-amd64) / [arm](https://github.com/jessfraz/ship/releases/download/v0.0.0/ship-linux-arm) / [arm64](https://github.com/jessfraz/ship/releases/download/v0.0.0/ship-linux-arm64)
- **solaris** [amd64](https://github.com/jessfraz/ship/releases/download/v0.0.0/ship-solaris-amd64)
- **windows** [386](https://github.com/jessfraz/ship/releases/download/v0.0.0/ship-windows-386) / [amd64](https://github.com/jessfraz/ship/releases/download/v0.0.0/ship-windows-amd64)

#### Via Go

```bash
$ go get github.com/jessfraz/ship
```

#### Running with Docker

```console
$ docker run --rm -it \
    -v /etc/localtime:/etc/localtime:ro \
    --name ship \
    -e "AFTERSHIP_API_KEY=your_api_key" \
    r.j3ss.co/ship
```

## Usage

```console
$ ship -h
     _     _
 ___| |__ (_)_ __
/ __| '_ \| | '_ \
\__ \ | | | | |_) |
|___/_| |_|_| .__/
            |_|

 Command line tool for tracking packages using the AfterShip API.
 Version: v0.0.0
 Build: 

  -apikey string
        AfterShip API Key (or env var AFTERSHIP_API_KEY) 
  -d    run in debug mode
  -v    print version and exit (shorthand)
  -version
        print version and exit
```
