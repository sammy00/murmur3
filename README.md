# MurmurHash3

A Golang Implementation of MurmurHash3.

[![CircleCI](https://circleci.com/gh/sammyne/murmur3.svg?style=svg)](https://circleci.com/gh/sammyne/murmur3)
[![codecov](https://codecov.io/gh/sammyne/murmur3/branch/master/graph/badge.svg)](https://codecov.io/gh/sammyne/murmur3)
[![Go Report Card](https://goreportcard.com/badge/github.com/sammyne/murmur3)](https://goreportcard.com/report/github.com/sammyne/murmur3)
[![LICENSE](https://img.shields.io/badge/license-ISC-blue.svg)](LICENSE)

## Overview

Murmur3 is the 3rd version of MurmurHash by Austin Appleby. As described in wiki, it serves as a general hash-table lookup algorithm. **When it comes to need of a cryptographic hash function, this should be definitely out of your choice.**

## Requirement

- go-1.11

## Installation

```bash
go get -u -v github.com/sammyne/murmur3
```

## References

- [MurmurHash Wiki](https://en.wikipedia.org/wiki/MurmurHash)
- [MurmurHash3](https://github.com/aappleby/smhasher/wiki/MurmurHash3#bulk-speed-test-hashing-an-8-byte-aligned-256k-block)
