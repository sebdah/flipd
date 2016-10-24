# flipd

Feature flipper service implemented in [Go](https://golang.org/).

## Development

### Prerequisites

- [Thrift](https://thrift.apache.org/)

### Dependencies

The dependencies are maintained via git submodules in the `vendor` folder. To
fetch all dependencies, run

    git submodule update --init --recursive

We are using [vendetta](https://github.com/dpw/vendetta) for maintaining the
dependencies. Use it like this:

 - `vendetta` - Vendorize your new dependencies
 - `vendetta -p` - Remove unused dependencies
 - `vendetta -u` - Upgrade dependencies

### Running tests

The tests can be executed by running:

    make clean test

## Releasing

1. Generate Thrift interfaces and objects, `make generate-thrift`

## Author

[Sebastian Dahlgren](http://sebastiandahlgren.se)
([@sebdah](https://github.com/sebdah))

## License

The MIT License (MIT)
Copyright (c) 2016 Sebastian Dahlgren

Permission is hereby granted, free of charge, to any person obtaining a copy of
this software and associated documentation files (the "Software"), to deal in
the Software without restriction, including without limitation the rights to
use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
the Software, and to permit persons to whom the Software is furnished to do so,
subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.-
