package notify

import (
	"runtime"
	"strings"
)

// This error handler will send notification to discord 
func ErrorHandler(err error) error {
	DiscordNotif("", err.Error())
	
	return err
}

func GetRelativePath() string {
	_, filename, _, ok := runtime.Caller(2)
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