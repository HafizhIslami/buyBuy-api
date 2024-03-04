package notify

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func DiscordNotif(state string, message string) {
	switch state {
	case "error":
		filePath := GetRelativePath()
		// JSON payload to send to the webhook
		payload := map[string]string{
			"content": fmt.Sprintf("**Error executing file:** %s\n\n%s", filePath, message),
		}

		// Convert payload to JSON
		payloadBytes, err := json.Marshal(payload)
		if err != nil {
			log.Fatal(err)
		}

		// Send POST request to the webhook URL
		_, err = http.Post(os.Getenv("WEBHOOK_NOTIFICATION_URL"), "application/json", bytes.NewBuffer(payloadBytes))
		if err != nil {
			log.Fatal(err)
		}
    default:
	}
}
