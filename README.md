# [ncalc](https://godoc.org/github.com/clarketm/ncalc)

Command line utility for *quick* number base conversions ( **ascii** / **binary** / **octal** / **decimal** / **hexadecimal** ).

![release-badge](https://img.shields.io/github/release/clarketm/ncalc.svg)
![circleci-badge](https://circleci.com/gh/clarketm/ncalc.svg?style=shield)

```shell

NAME:
    ncalc – number base converter.

SYNOPSIS:
    ncalc [ opts... ] [ number|character ]

OPTIONS:
    -h, --help                  print usage.
    -i, --input format          input format. see FORMATS. (default: decimal|ascii)
    -o, --output format         output format. see FORMATS. (default: all)
    -q, --quiet                 suppress printing of output format type(s)
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

#### Convert `ascii` to `decimal`
```bash
# Short form
$ ncalc -i a -o d 'w'

decimal: 119


# Long form
$ ncalc -i ascii -o decimal 'w'

decimal: 119


# Very long form
$ ncalc --input ascii --output decimal 'w'

decimal: 119


# Quite mode (-q|--quiet)
$ ncalc -- -i ascii --output decimal 'w'

119
```

#### Convert `decimal` to `binary`
```shell
$ ncalc -i decimal -o binary 170

binary: 10101010
```

#### Convert `binary` to `decimal`
```shell
$ ncalc -i b -o d 10101010

decimal: 170
```

#### Convert `ascii` to `all` formats
```shell
$ ncalc -i a 'G'

ascii: 'G'
binary: 1000111
octal: 107
decimal: 71
hexadecimal: 47
```

---

You can see the full reference documentation for the **ncalc** package at [godoc.org](https://godoc.org/github.com/clarketm/ncalc), or through go's standard documentation system:
```bash
$ godoc -http=:6060

# Open browser to: "http://localhost:6060/pkg/github.com/clarketm/ncalc"  to view godoc.
```
