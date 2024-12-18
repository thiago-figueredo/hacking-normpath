package normpath

import "testing"
import "github.com/stretchr/testify/assert"

type TestCase struct {
	input    string
	expected string
}

func TestReplacesDoubleSlashesWithOneSlash(t *testing.T) {
	testCases := []TestCase{
		{input: "//", expected: "/"},
		{input: "foo//bar", expected: "foo/bar"},
		{input: "foo/bar//baz", expected: "foo/bar/baz"},
	}

	for _, tc := range testCases {
		assert.Equal(t, tc.expected, normpath(tc.input))
	}
}

func TestRemovesSlashOnTheRight(t *testing.T) {
	testCases := []TestCase{
		{input: "//", expected: "/"},
		{input: "foo//", expected: "foo"},
		{input: "foo//bar/", expected: "foo/bar"},
		{input: "foo//bar///", expected: "foo/bar"},
		{input: "foo/bar///", expected: "foo/bar"},
		{input: "foo//bar////", expected: "foo/bar"},
		{input: "foo/bar//baz//", expected: "foo/bar/baz"},
		{input: "foo/bar//baz///", expected: "foo/bar/baz"},
	}

	for _, tc := range testCases {
		assert.Equal(t, tc.expected, normpath(tc.input))
	}
}

func TestIgnoresDotSlash(t *testing.T) {
	testCases := []TestCase{
		{input: "./", expected: "."},
		{input: "././", expected: "."},
		{input: "./foo./bar", expected: "foo./bar"},
		{input: "./foo/./bar", expected: "foo/bar"},
		{input: "./foo./bar/./", expected: "foo./bar"},
		{input: "./foo./bar/.//", expected: "foo./bar"},
		{input: "./foo./bar/.///", expected: "foo./bar"},
		{input: "./foo./bar/.////", expected: "foo./bar"},
		{input: "/./", expected: "/"},
		{input: "//.//", expected: "/"},
		{input: "foo/./", expected: "foo"},
		{input: "foo/.//bar.//", expected: "foo/bar."},
		{input: "foo//.bar///", expected: "foo/.bar"},
		{input: "foo/bar.baz.//", expected: "foo/bar.baz."},
		{input: ".foo./bar././baz./././", expected: ".foo./bar./baz."},
	}

	for _, tc := range testCases {
		assert.Equal(t, tc.expected, normpath(tc.input))
	}
}

func TestExpandDoubleDotSlash(t *testing.T) {
	testCases := []TestCase{
		{input: "/../", expected: "."},
		{input: "../", expected: ".."},
		{input: "./../", expected: ".."},
		{input: ".././", expected: ".."},
		{input: "./foo../bar", expected: "foo../bar"},
		{input: "./foo/../bar", expected: "bar"},
		{input: "./foo./bar/../../", expected: "."},
		{input: "./foo./bar/./../", expected: "foo."},
		{input: "./foo./bar/.../../...././", expected: "foo./bar/...."},
		{input: "//..//", expected: "/"},
		{input: "//..////..//", expected: "/"},
		{input: "//..//.//..//", expected: "/"},
		{input: "//..//..//..//", expected: "/"},
		{input: "foo/./", expected: "foo"},
		{input: "foo/.//bar.//", expected: "foo/bar."},
		{input: "foo//.bar///", expected: "foo/.bar"},
		{input: "foo/bar.baz.//", expected: "foo/bar.baz."},
		{input: ".foo./bar././baz./././", expected: ".foo./bar./baz."},
	}

	for _, tc := range testCases {
		assert.Equal(t, tc.expected, normpath(tc.input))
	}
}
