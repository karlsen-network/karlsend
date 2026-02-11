# karlsenctl

`karlsenctl` is an RPC client for `karlsend`.

## Requirements

Go 1.24 or later.

## Build from Source

* Install Go according to the installation instructions here:
  http://golang.org/doc/install

* Ensure Go was installed properly and is a supported version:

```bash
go version
```

* Run the following commands to obtain and install `karlsend`
  including all dependencies:

```bash
git clone https://github.com/karlsen-network/karlsend
cd karlsend/cmd/karlsenctl
go install .
```

* `karlsenctl` should now be installed in `$(go env GOPATH)/bin`. If
  you did not already add the bin directory to your system path
  during Go installation, you are encouraged to do so now.

## Usage

The full `karlsenctl` configuration options can be seen with:

```bash
karlsenctl --help
```

But the minimum configuration needed to run it is:

```bash
karlsenctl <REQUEST_JSON>
```

For example:

```
karlsenctl '{"getBlockDagInfoRequest":{}}'
```

For a list of all available requests check out the [RPC documentation](infrastructure/network/netadapter/server/grpcserver/protowire/rpc.md)
