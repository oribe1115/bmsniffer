# bmsniffer

## Installation

```sh
$ go get -u github.com/oribe1115/bmsniffer/cmd/bmsniffer
```

## Usage

```sh
$ go vet -vettool=`which bmsniffer` ./...
```

## Flags

### bmsniffer.loc
Baseline for LOC
Default: `0`

```sh
$ go vet -vettool=`which bmsniffer` -bmsniffer.loc=10 ./...
```

### bmsniffer.maxnesting
Baseline for MAXNESTING
Default: `0`

```sh
$ go vet -vettool=`which bmsniffer` -bmsniffer.maxnesting=3 ./...
```

### bmsniffer.nov
Baseline for NOV
Default: `0`

```sh
$ go vet -vettool=`which bmsniffer` -bmsniffer.nov=10 ./...
```

### bmsniffer.cyclo
Baseline for CYCLO
Default: `0`

```sh
$ go vet -vettool=`which bmsniffer` -bmsniffer.cyclo=3 ./...
```

### bmsniffer.test
If `bmsniffer.test` is `true`, output includes test functions
Default: `false`

```sh
$ go vet -vettool=`which bmsniffer` -bmsniffer.test=true ./...
```
