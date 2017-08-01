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
### Installation
#### Source (Mac/Linux)
#####
```bash
$ BUILD=darwin_amd64.tar.gz     # Mac (64 bit)
$ BUILD=linux_amd64.tar.gz      # Linux (64 bit)

$ curl -O https://github.com/clarketm/ncalc/releases/download/v1.0.0/$BUILD
$ { BIN_DIR=/usr/local/bin; mkdir -p $BIN_DIR; tar -xzf build/darwin_amd64.tar.gz -C $BIN_DIR; }
```
#### Source (Windows)
* https://github.com/clarketm/ncalc/releases/download/v1.0.0/windows_amd64.zip

#### Golang
$ go get github.com/clarketm/ncalc
