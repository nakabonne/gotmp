# gotmp

Sets up a temporary go project in module mode.

## Installation

```
go get -u github.com/nakabonne/gotmp
```

## Usage

Running gotmp creates a temporary directory and opens up main.go in `$EDITOR`.

```
gotmp
```

It creates the following files underneath a temporary directory to be removed after work:

```
.
├── go.mod
└── main.go
```

