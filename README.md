# English

**English** is a Go library for pretty-printing values in English.

## Install

    go get github.com/bfontaine/english

## Usage

* `english.Bool(...)` returns `"yes"` for `true` and `"no"` for `false`
* `english.OrdinalLiteral(...)` returns an ordinal literal, e.g. `"1st"` for
  `1`, `"123rd"` for `123`, etc.
* `english.Plural(word, count)` returns `word` either as is or with an `s`
  appended if `count` is different than `1`.

I only implemented the functions I needed; which is why there are only three of
them for now.
