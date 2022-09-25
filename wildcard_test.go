package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	p string
	s string
	m bool
}

func TestWildCardMatch(t *testing.T) {

	testCases := []testCase{
		{"", "", true},
		{"*", "", true},
		{".", "", false},
		{"", "a", false},
		{"abc", "abc", true},
		{"abc", "ac", false},
		{"abc", "abd", false},
		{"a.c", "abc", true},
		{"a*c", "abc", true},
		{"a*c", "abcbc", true},
		{"a*c", "abcbd", false},
		{"a*b..c", "ajcbjcklbjic", true},
		{"a*b..c", "ajcbjcklbjimc", false},
		{"a*b*c", "ajkembbcldkcedc", true},
	}

	for _, tc := range testCases {
		m, err := WildCardMatch(tc.p, tc.s)
		assert.Equal(t, m, tc.m)
		assert.Nil(t, err)
	}
}
