# wildcard

A simple golang customizable [wildcard](https://en.wikipedia.org/wiki/Wildcard_character) matcher. Golang has pretty well built regex functionalities, but it does not have basic wildcard matcher that works as nicely. Therefore this package serves the need to check whether a string matches a pattern in the rule of wildcard.

To keep simplicity, the matcher supports only two rules:
- `"?"` for a single char.
- `"*"` for any number (including zero) of chars.

Charset like `"[A-Za-z]"` or SQL style wild cards like `%` are not supported.

## usage

To import the package, `go get` the module.

```
go get -u github.com/vodkaslime/wildcard@main
```



To match pattern, use a matcher.

```
package main

import (
	"github.com/vodkaslime/wildcard"
)

func main() {
	matcher := wildcard.NewMatcher()
	p := "a?c"
	s := "abc"
	m, _ := matcher.Match(p, s)
	println(m)
}

```

The default wildcard chars are `"?"` for single chars and `"*"` for multiple chars. To customize this rule, tune the `S` field and `M` field accordingly.

For example to use `"."` as single char wildcard symbol:

```
package main

import (
	"github.com/vodkaslime/wildcard"
)

func main() {
	matcher := wildcard.NewMatcher()
	matcher.S = '.'
	p := "a.c"
	s := "abc"
	m, _ := matcher.Match(p, s)
	println(m)
}

```
