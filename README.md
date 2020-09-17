# ship

[![make-all](https://github.com/jessfraz/ship/workflows/make%20all/badge.svg)](https://github.com/jessfraz/ship/actions?query=workflow%3A%22make+all%22)
[![make-image](https://github.com/jessfraz/ship/workflows/make%20image/badge.svg)](https://github.com/jessfraz/ship/actions?query=workflow%3A%22make+image%22)
[![GoDoc](https://img.shields.io/badge/godoc-reference-5272B4.svg?style=for-the-badge)](https://godoc.org/github.com/jessfraz/ship)
[![Github All Releases](https://img.shields.io/github/downloads/jessfraz/ship/total.svg?style=for-the-badge)](https://github.com/jessfraz/ship/releases)

Command line tool to track packages using the [AfterShip API](https://docs.aftership.com/api/4/overview).

**Table of Contents**

<!-- toc -->

- [Installation](#installation)
    + [Binaries](#binaries)
    + [Via Go](#via-go)
    + [Running with Docker](#running-with-docker)
- [Usage](#usage)
  * [Create a Shipment](#create-a-shipment)
  * [Get a Shipment](#get-a-shipment)
  * [List Shipments](#list-shipments)
  * [Delete a Shipment](#delete-a-shipment)

<!-- tocstop -->

## Installation

#### Binaries

For installation instructions from binaries please visit the [Releases Page](https://github.com/jessfraz/ship/releases).

#### Via Go

```console
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
ship -  Command line tool to track packages using the AfterShip API.

Usage: ship <command>

Flags:

  -d        enable debug logging (default: false)
  --apikey  AfterShip API Key (or env var AFTERSHIP_API_KEY) (default: <none>)

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

  --apikey  AfterShip API Key (or env var AFTERSHIP_API_KEY) (default: <none>)
  -d        enable debug logging (default: false)
```

### Get a Shipment

```console
$ ship get -h
Usage: ship get [OPTIONS] TRACKING_NUMBER

Get details for a shipment.

Flags:

  --apikey  AfterShip API Key (or env var AFTERSHIP_API_KEY) (default: <none>)
  -d        enable debug logging (default: false)
```

### List Shipments

```console
$ ship ls -h
Usage: ship ls 

List shipments.

Flags:

  --apikey  AfterShip API Key (or env var AFTERSHIP_API_KEY) (default: <none>)
  -d        enable debug logging (default: false)
```

### Delete a Shipment

```console
$ ship rm -h
Usage: ship rm [OPTIONS] TRACKING_NUMBER

Delete a shipment.

Flags:

  --apikey  AfterShip API Key (or env var AFTERSHIP_API_KEY) (default: <none>)
  -d        enable debug logging (default: false)
```