# karlsenminer

`karlsenminer` is a CPU-based miner for `karlsend`.

## Requirements

Go 1.19 or later.

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
cd karlsend/cmd/karlsenminer
go install .
```

* `karlsenminer` should now be installed in `$(go env GOPATH)/bin`.
  If you did not already add the bin directory to your system path
  during Go installation, you are encouraged to do so now.
  
## Usage

The full `karlsenminer` configuration options can be seen with:

```bash
karlsenminer --help
```

But the minimum configuration needed to run it is:

```bash
karlsenminer --miningaddr=<YOUR_MINING_ADDRESS>
```
