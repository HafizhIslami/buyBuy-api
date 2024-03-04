package utils

import (
	"runtime"
	"strings"
)

func getRelativePath() string {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		return "unknown"
	}

	// Assuming your project directory is "buyBuy"
	projectDir := "buyBuy-api"
	idx := strings.Index(filename, projectDir)
	if idx == -1 {
		return filename // Couldn't find project directory, return full path
	}
	return filename[idx+len(projectDir)+1:] // +1 to also remove the "/"
}