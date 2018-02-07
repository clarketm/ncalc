# [ncalc](https://godoc.org/github.com/clarketm/ncalc)

Command line utility for *quick* number base conversions ( **ascii** / **binary** / **octal** / **decimal** / **hexadecimal** ).

![release-badge](https://img.shields.io/github/release/clarketm/ncalc.svg)
![circleci-badge](https://circleci.com/gh/clarketm/ncalc.svg?style=shield&circle-token=51853e44a4aff2fef83b0b89407ed15288bd641c)

```shell

NAME:
    ncalc – number base converter.

SYNOPSIS:
    ncalc [ opts... ] [ number|character ]

OPTIONS:
    -h, --help                  print usage.
    -i, --input format          input format. see FORMATS. (default: decimal|ascii)
    -o, --output format         output format. see FORMATS. (default: all)
    -v, --version               print version number.

FORMATS:
    (a)scii                     character
    (b)inary                    base 2
    (o)ctal                     base 8
    (d)ecimal                   base 10
    (h)exadecimal               base 16

EXAMPLES:
    ncalc "6"                               # output `decimal` number `6` in `all` formats
    ncalc "G"                               # output `ascii` character `G` in `all` formats
    ncalc -i a "f"                          # output `ascii` character `f` in `all` formats
    ncalc -i decimal -o ascii "15"          # output `decimal` number `15` as `ascii`
    ncalc --input h --output o "ff"         # output `hexadecimal` number `ff` as `octal`

```
## Installation

#### Golang
```shell
$ go get -u github.com/clarketm/ncalc
```

#### Install script
```shell
$ git clone "https://github.com/clarketm/ncalc.git"
$ cd ncalc && sudo sh install.sh
```

#### Source (Mac/Linux)
```shell
# List of builds: https://github.com/clarketm/ncalc/releases/

$ BUILD=darwin_amd64.tar.gz     # Mac (64 bit)
# BUILD=linux_amd64.tar.gz      # Linux (64 bit)

$ BIN_DIR=/usr/local/bin        # `bin` install directory
$ mkdir -p $BIN_DIR

$ curl -L https://github.com/clarketm/ncalc/releases/download/v1.1.2/$BUILD | tar xz -C $BIN_DIR        # install
```

#### Source (Windows)
* https://github.com/clarketm/ncalc/releases/download/v1.1.2/windows_amd64.zip


## Usage

You can see the full reference documentation for the **ncalc** package at [godoc.org](https://godoc.org/github.com/clarketm/ncalc), or through go's standard documentation system:
```bash
$ godoc -http=:6060

# Open browser to: "http://localhost:6060/pkg/github.com/clarketm/ncalc"  to view godoc.
```
