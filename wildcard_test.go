package wildcard

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type wildPatternTestCase struct {
	p string
	m bool
}

type matchTestCase struct {
	p string
	s string
	m bool
}

func TestIsWildPattern(t *testing.T) {
	testCases1 := []wildPatternTestCase{
		{"*", true},
		{"**", true},
		{"*?", true},
		{"?", true},
		{".", false},
		{"a", false},
		{"a?c", true},
	}

	m1 := NewMatcher()
	for _, tc := range testCases1 {
		b := m1.isWildPattern(tc.p)
		if !assert.Equal(t, b, tc.m) {
			println(tc.p, tc.m)
		}
	}

	testCases2 := []wildPatternTestCase{
		{"*", true},
		{"**", true},
		{"*.", true},
		{"?", false},
		{".", true},
		{"a", false},
		{"a.c", true},
	}

	m2 := NewMatcher()
	m2.S = '.'
	for _, tc := range testCases2 {
		b := m2.isWildPattern(tc.p)
		if !assert.Equal(t, b, tc.m) {
			println(tc.p, tc.m)
		}
	}

}

func TestMatch(t *testing.T) {

	testCases1 := []matchTestCase{
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

	testCases2 := []matchTestCase{
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
