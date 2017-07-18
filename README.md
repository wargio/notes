# notes
[![GoDoc](https://godoc.org/github.com/prologic/notes?status.svg)](https://godoc.org/github.com/prologic/notes)
[![Go Report Card](https://goreportcard.com/badge/github.com/prologic/notes)](https://goreportcard.com/report/github.com/prologic/notes)

notes is a self-hosted note taking web app that lets you keep track of your
notes and search them in a easy and minimal way.

## Installation

### Source

```#!bash
$ go install github.com/prologic/notes/...
```

## Usage

Run notes:

```#!bash
$ notes
```

Then visit: http://localhost:8000/

## Configuration

By default notes stores notes in `./data` in the current directory
and metadata in `notes.db` in the current directory. This can
be configured with the `-data` and `-dbpath` options.

## License

MIT
