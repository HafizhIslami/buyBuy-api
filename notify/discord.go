package notify

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
)

func DiscordNotification() {
    // JSON payload to send to the webhook
    payload := map[string]string{
        "content": "Hello, Discord!",
    }

    // Convert payload to JSON
    payloadBytes, err := json.Marshal(payload)
    if err != nil {
        panic(err)
    }

    // Send POST request to the webhook URL
    _, err = http.Post(os.Getenv("WEBHOOK_NOTIFICATION_URL"), "application/json", bytes.NewBuffer(payloadBytes))
    if err != nil {
        panic(err)
    }
}
