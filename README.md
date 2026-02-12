# Karlsend

[![Build Status](https://github.com/karlsen-network/karlsend/actions/workflows/tests.yaml/badge.svg)](https://github.com/karlsen-network/karlsend/actions/workflows/tests.yaml)
[![Code Coverage](https://codecov.io/gh/karlsen-network/karlsend/graph/badge.svg)](https://codecov.io/gh/karlsen-network/karlsend)
[![GitHub release](https://img.shields.io/github/v/release/karlsen-network/karlsend.svg)](https://github.com/karlsen-network/karlsend/releases)
[![GitHub license](https://img.shields.io/github/license/karlsen-network/karlsend.svg)](https://github.com/karlsen-network/karlsend/blob/master/LICENSE)
[![GitHub downloads](https://img.shields.io/github/downloads/karlsen-network/karlsend/total.svg)](https://github.com/karlsen-network/karlsend/releases)
[![Discord users](https://img.shields.io/discord/1169939685280337930.svg)](https://discord.gg/ZPZRvgMJDT)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/karlsen-network/karlsend/)

Karlsend is the reference full node Karlsen implementation written in
Go (golang). It is a [blockDAG](https://en.wikipedia.org/wiki/Directed_acyclic_graph)
as a proof-of-work cryptocurrency with instant confirmations and
sub-second block times with a work-in-progress (WIP) attempt to
implement smart contracts as a side-chain. It is based on
[the PHANTOM protocol](https://eprint.iacr.org/2018/104.pdf), a
generalization of Nakamoto consensus.

## Overview

Karlsen is a fork of [Kaspa](https://github.com/kaspanet/kaspad)
introducing a GPU-centric fork as a solution to the dominance of ASIC
mining farms, aiming to empower small-scale miners and enhance
decentralization. We focus on bridging the gap between blockchain
technology, decentralized finance and the real world of payment systems
and traditional finance.

With Kaspa, our approach is one of friendly cohabitation. We operate
under the same protocol, enjoy similar advantages, and face
unidentified issues. Any significant improvements that need to be made
in the primary codebase will be shared back.

## Small Scale Miners

The Karlsen Network team believes in decentralization and small-scale
miners. We will ensure long-term GPU-friendly mining.

### Hashing Function

We initially started with `kHeavyHash` and `blake3` modifications
on-top. This algorithm is called `KarlsenHashv1`.

`KarlsenHashv1` is currently used in [mainnet](https://github.com/karlsen-network/karlsend/releases/tag/v1.1.0)
and can be mined using the following miners maintained by the Karlsen
developers:

* Built-in CPU miner from `karlsend`
* Karlsen [GPU miner](https://github.com/karlsen-network/karlsen-miner)
  as reference implementation of `kHeavyHash` with `blake3`.

The following third-party miners are available and have added
`KarlsenHashv1`:

* [lolMiner](https://github.com/Lolliedieb/lolMiner-releases)
* [Team Red Miner](https://github.com/todxx/teamredminer)
* [SRBMiner](https://github.com/doktor83/SRBMiner-Multi)
* [BzMiner](https://github.com/bzminer/bzminer)
* [Rigel](https://github.com/rigelminer/rigel)
* [GMiner](https://github.com/develsoftware/GMinerRelease)

`KarlsenHashv2` will become active via hardfork at DAA score `26.962.009`.
It is based on [FishHash](https://github.com/iron-fish/fish-hash/blob/main/FishHash.pdf)
written from scratch in our Golang node implementation. It is FPGA/ASIC
resistent. It is the worlds first implementation of FishHash with Golang
in `mainnet` in a 1bps blockchain.

`KarlsenHashv2` is currently used in [mainnet](https://github.com/karlsen-network/karlsend/releases/tag/v2.1.0)
and can be mined using the following miners maintained by the Karlsen
developers:

* Built-in CPU miner from `karlsend`
* Karlsen [GPU miner](https://github.com/karlsen-network/karlsen-miner/releases/tag/v2.0.0)
  as bleeding edge and unoptimized reference implementation of
  `KarlsenHashv2`. Please follow the steps in the [README.md](https://github.com/karlsen-network/karlsen-miner/blob/main/README.md)
  to generate a DAG file.

The following third-party miners are available and have added
`KarlsenHashv2`:

* [SRBMiner](https://github.com/doktor83/SRBMiner-Multi)

### DAG Generation

To start mining using the built-in CPU miner it needs a pre-generated
DAG file. `KarlsenHashv2` miner uses a 4GB DAG for efficient mining.
It generates this DAG with 8 CPU threads and saves it as `hashes.dat`
for faster subsequent runs.

* First Run: Generates a 4GB DAG using 8 CPU threads. This may take
  time depending on your computer. Saves the DAG as `hashes.dat` for
  future use.
* Next Runs: Loads `hashes.dat` to skip DAG generation, speeding up
  startup.

If you need to regenerate the DAG, delete `hashes.dat` and run the
`karlsenminer` again.

## Smart Contracts

The Karlsen Network team is launching an R&D project to connect the
Karlsen blockchain with other blockchain ecosystems using a smart
contract layer based on the [Cosmos SDK](https://v1.cosmos.network/sdk).

This initiative aims to enhance interoperability, efficiency, and
innovation in the blockchain space. By leveraging the Cosmos SDK's
advanced features, we'll facilitate seamless transactions across
different networks, offering new opportunities for users and
developers.

### Cosmos Hub

[Cosmos](https://cosmos.network/) is a highly attractive ecosystem due
to its innovative approach to blockchain interoperability, scalability,
and usability. By enabling different blockchains to seamlessly
communicate and exchange value, Cosmos opens up vast opportunities for
businesses and developers to build and deploy decentralized
applications that can operate across multiple blockchain environments.
This interoperability fosters a more connected and efficient digital
economy, potentially driving adoption and usage across various sectors,
including finance, supply chain, and beyond.

By connecting to Cosmos, we will open the door of a web3 ecosystem
connected to a complete network of other blockchain project, making
Karlsen Network more competitive and adaptable in the rapidly evolving
landscape of decentralized technologies.

### Karlsen Sidechain

The creation of the Karlsen [sidechain](https://github.com/john-light/sidechains)
with fast transaction times, smart contract capabilities, and a dual
coin model, designed to integrate seamlessly with the Cosmos ecosystem
and utilize the Karlsen (KLS) across all interconnected platforms,
signifies a strategic advancement in Karlsen Network.

This sidechain will not only enable quick and efficient
inter-blockchain transactions but also support complex decentralized
applications through its smart contract functionality.

The ability to use Karlsen (KLS) across the entire ecosystem will
ensure a unified and streamlined user experience, promoting greater
adoption and utility within the Cosmos network. This sidechain aims
to enhance scalability, foster innovation, and provide a flexible and
user-centric blockchain solution that meets the diverse needs of
developers, users, and investors within the Cosmos ecosystem.

## Installation

### Install from Binaries

Pre-compiled binaries for Linux `x86_64` and `aarch64`, Windows `x64`
and macOS `x64` as universal binary can be downloaded at: [https://github.com/karlsen-network/karlsend/releases](https://github.com/karlsen-network/karlsend/releases)

### Build from Source

Go 1.24 or later is required. Install Go according to the installation
instructions at [http://golang.org/doc/install](http://golang.org/doc/install).
Ensure Go was installed properly and is a supported version:

```bash
go version
```

Run the following commands to obtain and install karlsend including
all dependencies:

```bash
git clone https://github.com/karlsen-network/karlsend
cd karlsend
go install . ./cmd/...
```

Karlsend (and utilities) should now be installed in
`$(go env GOPATH)/bin`. If you did not already add the `bin` directory
to your system path during Go installation, you are encouraged to do
so now.

### Getting Started

Karlsend has several configuration options available to tweak how it
runs, but all of the basic operations work with zero configuration.

```bash
karlsend
```

## Discord

Join our discord server using the following link: [https://discord.gg/ZPZRvgMJDT](https://discord.gg/ZPZRvgMJDT)

## Issue Tracker

The [integrated github issue tracker](https://github.com/karlsen-network/karlsend/issues)
is used for this project.

## Documentation

The [documentation](https://github.com/karlsen-network/docs) is a
work-in-progress.

## License

Karlsend is licensed under the copyfree [ISC License](https://choosealicense.com/licenses/isc/).
