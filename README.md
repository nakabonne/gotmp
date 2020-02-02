# gotmp

Sets up a temporary go project in module mode.

## Installation

```
go get -u github.com/nakabonne/gotmp
```

## Usage

To run gotmp execute:

```
gotmp
```
Running gotmp creates a temporary directory to be removed after work and opens up main.go in `$EDITOR`. It creates the following files:

```
.
├── go.mod
└── main.go
```

