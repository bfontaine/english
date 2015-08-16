# English

**English** is a Go library for pretty-printing values in English.

## Install

    go get github.com/bfontaine/english

## Usage

* `english.Bool(...)` returns `"yes"` for `true` and `"no"` for `false`
* `englishOrdinalLiteral(...)` returns an ordinal literal, e.g. `"1st"` for
  `1`, `"123rd"` for `123`, etc.
