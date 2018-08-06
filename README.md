# go-fourcc

Package **fourcc** is a Go implementation of **FOURCC** (**four character code**) (**4CC**) identifiers for a video codecs, compression formats, colors and pixel format used in media files. 

FOURCC (also sometimes called 4CC) is short for "four character code".

FOURCC is an identifier for a video codec, compression format, color or pixel format used in media files.

A character in this context is a 1 byte/8 bit value, thus a FOURCC always takes up exatly 32 bits/4 bytes in a file.


## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-fourcc

[![GoDoc](https://godoc.org/github.com/reiver/go-fourcc?status.svg)](https://godoc.org/github.com/reiver/go-fourcc)


## Example

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


And:
```go
var p []byte = []byte{'P','A','L','8'}

var colorspace fourcc.Type

err := colorspace.Scan(p)
