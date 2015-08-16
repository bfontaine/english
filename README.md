# English

**English** is a Go library for pretty-printing values in English.

## Install

    go get github.com/bfontaine/english

## Usage

* `english.Bool(...)` returns `"yes"` for `true` and `"no"` for `false`
* `english.OrdinalLiteral(...)` returns an ordinal literal, e.g. `"1st"` for
  `1`, `"123rd"` for `123`, etc.

I only implemented the function I needed; which is why there are only two of
them for now.
