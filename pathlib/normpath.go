package pathlib

import "strings"

func expandDoubleDot(path string, idx int) []rune {
	newPathPrefix := strings.TrimRightFunc(path[0:idx], func(r rune) bool { return r != '/' })
	return append([]rune(newPathPrefix), []rune(path[idx+4:])...)
}

func Normpath(path string) string {
	path = strings.TrimPrefix(path, "./")
	newPath := []rune{}

	if strings.HasPrefix(path, "/../") {
		return "."
	}

	if strings.HasPrefix(path, "/.") {
		return "/"
	}

	for i := 0; i < len(path); i += 1 {
		newPath = append([]rune(newPath), rune(path[i]))

		if i+1 < len(path) && path[i:i+2] == "//" {
			i += 1
		} else if i+2 < len(path) && path[i:i+3] == "/./" {
			i += 2
		} else if i+4 < len(path) && path[i:i+4] == "/../" {
			newPath = expandDoubleDot(path, i)
			i += len(path) - i + 4
		}
	}

	if len(newPath) > 1 && newPath[len(newPath)-1] == '/' {
		newPath = newPath[0 : len(newPath)-1]
	}

	if len(newPath) > 2 {
		newPath = []rune(strings.TrimSuffix(string(newPath), "/."))
	}

	if strings.HasSuffix(string(newPath), "/..") {
		newPath = expandDoubleDot(string(path), strings.Index(string(newPath), "/.."))
	}

	if string(newPath) != path {
		return Normpath(string(newPath))
	}

	if string(newPath) == "" {
		return "."
	}

	return string(newPath)
}
