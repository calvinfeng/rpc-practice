# gRPC (Google Remote Procedure Calls)

## Use Docker

If you use docker, then you can skip all the steps below.

    docker run -v $(pwd):/defs namely/protoc-all:1.17_0 -d protos -l go -o protogens/go/

## Install

### gRPC-Go

    go get -u google.golang.org/grpc

### Protocol Buffers

I prefer to use pre-built binaries. Just navigate to protobuf [release page][1] and download a binary
for your OS. I am using Linux so I downloaded `protoc-3.6.1-linux-x86_64.zip`. Then place the binary
into your `/usr/local/bin`.

    sudo unzip -o protoc-3.6.1-linux-x86_64.zip -d /usr/local bin/protoc

### Go Support

We also need Go support for protobuf. We can simply download the source code and compile the Go code.

```bash
GIT_TAG="v1.2.0"
go get -d -u github.com/golang/protobuf/protoc-gen-go
git -C "$(go env GOPATH)"/src/github.com/golang/protobuf checkout $GIT_TAG
go install github.com/golang/protobuf/protoc-gen-go
```

### Linter

This is a useful tool for linting protobuf code.

    go get github.com/ckaznocha/protoc-gen-lint

### Script

```bash
#!/bin/bash

# Colors Utilities
Color_Off='\033[0m'       # Text Reset
Black='\033[0;30m'        # Black
Red='\033[0;31m'          # Red
Green='\033[0;32m'        # Green
Yellow='\033[0;33m'       # Yellow
Blue='\033[0;34m'         # Blue
Purple='\033[0;35m'       # Purple
Cyan='\033[0;36m'         # Cyan
White='\033[0;37m'        # White

function info() {
    echo -e "[${Blue}INFO${Color_Off}]: $@"
}

function ok() {
    echo -e "[${Green}OKAY${Color_Off}]: $@"
}

function warning() {
    echo -e "[${Yellow}WARN${Color_Off}]: $@"
}

function error() {
    echo -e "[${Red}ERRO${Color_Off}]: $@"
}
 
function install_protoc-gen-lint() {
    ok "installing protoc-gen-lint"
    go get github.com/ckaznocha/protoc-gen-lint
}

function check() {
    info "checking prerequisites..."

    which protoc > /dev/null
    if [ $? -ne 0 ]; then
        error "protoc is missing, please install them at https://github.com/google/protobuf"
        exit 1
    else
        ok "protoc is installed properly"
    fi 

    which protoc-gen-lint > /dev/null
    if [ $? -ne 0 ]; then
        error "protoc-gen-lint is missing, please install them at https://github.com/ckaznocha/protoc-gen-lint"
        exit 1
    else
        ok "protoc-gen-lint is installed properly"
    fi

    which protoc-gen-go > /dev/null
    if [ $? -ne 0 ]; then
        error "protoc-gen-go is missing, please install v1.2.0 at https://github.com/golang/protobuf"
        warn "please ensure v1.2.0 is installed instead of master"
        exit 1
    else
        ok "protoc-gen-go is installed properly"
    fi
}

function lint() {
    local PROTOS="./protos"

    info "linting protobuf definitions"
    protoc \
    --plugin=protoc-gen-lint=${GOPATH}/bin/protoc-gen-lint \
    -I ${PROTOS} \
    --lint_out=sort_imports:. \
    ${PROTOS}/*/**.proto
}

function compile() {
    local OUTPUT="cprotos" 
    local PROTOS="./protos"
    local GO_OUT=$OUTPUT/golang
    local PY_OUT=$OUTPUT/python

    # Output is compiled protos
    rm -rf "$OUTPUT"
    mkdir -p "$OUTPUT"/{golang,python}

    info "generating protobuf Go outputs."
    for pathname in $(cd ${PROTOS} && find . -name '*.proto'); do
        dirname=$(dirname $pathname)

        rm -rf ${GO_OUT}/${pathname} 2>/dev/null
        mkdir -p ${GO_OUT}/${dirname}

        protoc \
            -I ${PROTOS} \
            --plugin=protoc-gen-go=${GOPATH}/bin/protoc-gen-go \
            --go_out=plugins=grpc:${GO_OUT} \
            ${PROTOS}/${pathname}

    done
}

function main() {
    if [[ "$GOPATH" == "" ]]; then
        warning "Required env var GOPATH is not set; aborting with error."
        go help gopath
        exit -1
    fi

    check
    lint
    compile
}

main $@
```

[1]: https://github.com/protocolbuffers/protobuf/releases