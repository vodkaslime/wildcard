package wildcard

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	p string
	s string
	m bool
}

func TestMatch(t *testing.T) {

	testCases1 := []testCase{
		{"", "", true},
		{"*", "", true},
		{"?", "", false},
		{"", "a", false},
		{"abc", "abc", true},
		{"abc", "ac", false},
		{"abc", "abd", false},
		{"a?c", "abc", true},
		{"a*c", "abc", true},
		{"a*c", "abcbc", true},
		{"a*c", "abcbd", false},
		{"a*b??c", "ajcbjcklbjic", true},
		{"a*b??c", "ajcbjcklbjimc", false},
		{"a*b*c", "ajkembbcldkcedc", true},
	}

	m1 := NewMatcher()
	for _, tc := range testCases1 {
		m, err := m1.Match(tc.p, tc.s)
		if !assert.Equal(t, m, tc.m) {
			println(tc.p, tc.s, tc.m)
		}
		assert.Nil(t, err)
	}

	m2 := NewMatcher()
	m2.S = '.'

	testCases2 := []testCase{
		{"", "", true},
		{"*", "", true},
		{".", "", false},
		{"", "a", false},
		{"abc", "abc", true},
		{"abc", "ac", false},
		{"abc", "abd", false},
		{"a.c", "abc", true},
		{"a?c", "abc", false},
		{"a*c", "abc", true},
		{"a*c", "abcbc", true},
		{"a*c", "abcbd", false},
		{"a*b..c", "ajcbjcklbjic", true},
		{"a*b.?c", "ajcbjcklbjic", false},
		{"a*b..c", "ajcbjcklbjimc", false},
		{"a*b*c", "ajkembbcldkcedc", true},
	}

	for _, tc := range testCases2 {
		m, err := m2.Match(tc.p, tc.s)
		if !assert.Equal(t, m, tc.m) {
			println(tc.p, tc.s, tc.m)
		}
		assert.Nil(t, err)
	}
}
