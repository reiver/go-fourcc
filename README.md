# go-fourcc

Package **fourcc** is a Go implementation of **FOURCC** (**four character code**) (**4CC**) identifiers for a video codecs, compression formats, colors, and pixel format used in media files.


## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-fourcc

[![GoDoc](https://godoc.org/github.com/reiver/go-fourcc?status.svg)](https://godoc.org/github.com/reiver/go-fourcc)


## Introduction

FOURCC, also sometimes written as **FourCC** and **4CC**, are all short for “four character code”.

**FOURCC**, is a sequence of _4 8-bit bytes_ packed into an _unsigned 32-bit integer_ which, **uniquely identifies** a data format.

(A character in this context is a 1 byte (i.e., 8 bit) value. Thus a FOURCC always takes up exatly 32 bits (i.e., 4 bytes).)

These _data formats_ tend to be video codecs, compression formats, colors, and pixel format used in media files.


## Samples

Here are some sample FOURCC.

| Label | Hexadecimal (unsigned integer) |
|-------|--------------------------------|
| PAL8  | 0x384C4150                     |
| VP80  | 0x30385056                     |
| H264  | 0x34363248                     |
| MJPG  | 0x47504A4D                     |

The way you can understand how these _labels_ get turned in _unsigned integers_ is, if we consider the FOURCC “PAL8”, then:
```
ASCII ‘P’ = hexadecimal 0x50
ASCII ‘A’ = hexadecimal 0x41
ASCII ‘L’ = hexadecimal 0x4C
ASCII ‘8’ = hexadecimal 0x38
```
And then these 8-bit bytes as packed into an unsigned integer in reverse order.

I.e.,:
```go
var value uint32 = (uint32('P') | (uint32('A') << 8) | (uint32('L') << 16) | (uint32('8') << 24))
```


## Examples

Here is an example usage:

```go
var colorspace fourcc.Type = fourcc.FOURCC('P','A','L','8')

```

Here is a constant:
```go
const (
	VP8 = fourcc.Const("VP80")
)
```

A fourcc.Const needs to be turned into a fourcc.Type before using it. So:
```go
var datum fourcc.Type
var err error

datum, err = VP8.FOURCC()
```

Some uses of Scan:
```go
var colorspace fourcc.Type

err := colorspace.Scan("PAL8")
```

And:
```go
var p [4]byte = [4]byte{'P','A','L','8'}

var colorspace fourcc.Type

err := colorspace.Scan(p)
```

And:
```go
var p []byte = []byte{'P','A','L','8'}

var colorspace fourcc.Type

err := colorspace.Scan(p)
```


## See Also

* [FOURCC.org](https://www.fourcc.org/)
