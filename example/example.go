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
