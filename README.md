# ship

[![Travis CI](https://img.shields.io/travis/jessfraz/ship.svg?style=for-the-badge)](https://travis-ci.org/jessfraz/ship)
[![GoDoc](https://img.shields.io/badge/godoc-reference-5272B4.svg?style=for-the-badge)](https://godoc.org/github.com/jessfraz/ship)
[![Github All Releases](https://img.shields.io/github/downloads/jessfraz/ship/total.svg?style=for-the-badge)](https://github.com/jessfraz/ship/releases)

## Installation

#### Binaries

For installation instructions from binaries please visit the [Releases Page](https://github.com/jessfraz/ship/releases).

- **darwin** [386](https://github.com/jessfraz/ship/releases/download/v0.0.4/ship-darwin-386) / [amd64](https://github.com/jessfraz/ship/releases/download/v0.0.4/ship-darwin-amd64)
- **freebsd** [386](https://github.com/jessfraz/ship/releases/download/v0.0.4/ship-freebsd-386) / [amd64](https://github.com/jessfraz/ship/releases/download/v0.0.4/ship-freebsd-amd64)
- **linux** [386](https://github.com/jessfraz/ship/releases/download/v0.0.4/ship-linux-386) / [amd64](https://github.com/jessfraz/ship/releases/download/v0.0.4/ship-linux-amd64) / [arm](https://github.com/jessfraz/ship/releases/download/v0.0.4/ship-linux-arm) / [arm64](https://github.com/jessfraz/ship/releases/download/v0.0.4/ship-linux-arm64)
- **solaris** [amd64](https://github.com/jessfraz/ship/releases/download/v0.0.4/ship-solaris-amd64)
- **windows** [386](https://github.com/jessfraz/ship/releases/download/v0.0.4/ship-windows-386) / [amd64](https://github.com/jessfraz/ship/releases/download/v0.0.4/ship-windows-amd64)

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
$ ship help
Usage: ship <command>

Commands:

  create   Create a shipment.
  get      Get details for a shipment.
  ls       List shipments.
  rm       Delete a shipment.
  version  Show the version information.
```

### Create a Shipment

```console
$ ship create -h
Usage: ship create [OPTIONS] TRACKING_NUMBER

Create a shipment.

Flags:

  -apikey  AfterShip API Key (or env var AFTERSHIP_API_KEY)
  -d       enable debug logging (default: false)
```

### Get a Shipment

```console
$ ship get -h
Usage: ship get [OPTIONS] TRACKING_NUMBER

Get details for a shipment.

Flags:

  -apikey  AfterShip API Key (or env var AFTERSHIP_API_KEY)
  -d       enable debug logging (default: false)
```

### List Shipments

```console
$ ship ls -h
Usage: ship ls 

List shipments.

Flags:

  -apikey  AfterShip API Key (or env var AFTERSHIP_API_KEY)
  -d       enable debug logging (default: false)
```

### Delete a Shipment

```console
$ ship rm -h
Usage: ship rm [OPTIONS] TRACKING_NUMBER

Delete a shipment.

Flags:

  -apikey  AfterShip API Key (or env var AFTERSHIP_API_KEY)
  -d       enable debug logging (default: false)
```
