# Go Compress v1.0.0

A Go utility program that <strong style="text-decoration: underline"><em>compresses</em> and <em>decompresses</em> files</strong>.

## Usage

```sh
go run [-help] [-algorithm <alg>] [-decompress] [-print-algorithms] [-verbose] [-quiet] main.go <input-file1> <output-file1> [input-file2] [output-file2] ...
```

- The `-quiet` flag **silences all output** and _overrides_ the `-verbose` flag.
- The program does **not** throw an error when there aren't an _even number_ of input and output _files_. The program will loop over pairs of input and output files _until there is one left out_ (the odd one), ignoring that file. For example, <span style="text-decoration: underline">`in1.txt out1.txt in2.txt` will only compress `in1.txt` and `in2.txt`</span>.

## Supported Algorithms

0. <strong>Run-Length Encryption</strong> (`rle`): Replaces **continuous characters** of the same value with a character's value that is the count and then the actual character. For instance, `aaaaaaaaaabbbbbbbbbb` will turn into `<NEWLINE>a<NEWLINE>b` since there are ten of each and the value of `<NEWLINE>` is ten.
