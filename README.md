# gRPC
## Install 
### Protocol Buffers
I prefer to use pre-built binaries. Just navigate to protobuf [release page][1] and download a 
binary for your OS. I am using Linux so I downloaded `protoc-3.6.1-linux-x86_64.zip`. Then place the
binary into your `/usr/local/bin`.
```bash
sudo unzip -o protoc-3.6.1-linux-x86_64.zip -d /usr/local bin/protoc
```

### Go Support.
We also need Go support for protobuf. We can simply download the source code and compile the Go code.
```
GIT_TAG="v1.2.0"
go get -d -u github.com/golang/protobuf/protoc-gen-go
git -C "$(go env GOPATH)"/src/github.com/golang/protobuf checkout $GIT_TAG
go install github.com/golang/protobuf/protoc-gen-go
```

### Linter
This is a useful tool for linting protobuf code.
```
go get github.com/ckaznocha/protoc-gen-lint
```

[1]: https://github.com/protocolbuffers/protobuf/releases