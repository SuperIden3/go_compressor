# Go Compress

A Go utility program that compresses files.

## Usage

```sh
go run [-help] [-algorithm <alg>] [-decompress] [-print-algorithms] main.go <input-file1> <output-file1> [input-file2] [output-file2] ...
```

## Supported Algorithms

1. Run-Length Encryption (`rle`): Replaces continuous characters of the same value with a character's value that is the count and then the actual character. For instance, `aaaaaaaaaabbbbbbbbbb` will turn into `<NEWLINE>a<NEWLINE>b` since there are ten of each and the value of `<NEWLINE>` is ten.
