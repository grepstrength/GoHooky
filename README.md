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
- Windows operating system (for current implementation) for the target host

## Installation

1. Clone the repository.
2. Navigate to the project directory.
3. Replace `<ENTERWEBHOOKURL>` with your actual Discord webhook's URL. (Line 55)
3. Run `go build -ldflags "-s -w" main.go` to compile.

## Usage

The application provides two main functions:

- `runCommand(cmd string)`: Executes a Windows command and returns its output
- `sendToDiscord(webhookURL string, message string)`: Sends a message to a Discord channel via webhook

*Note: If you want to make this run on a Linux host, change `exec.Command("cmd", "/C", cmd)` to `exec.Command("bash", "-c", cmd)` on Line 19.*

## Examples

![1_iV_kbfBYECMoMNtsOoh8yw](https://github.com/user-attachments/assets/3bd9e55f-3a91-447a-a1e7-757033c531e8)

![1_JDYFWDr0CtQtBTn5sP23UQ](https://github.com/user-attachments/assets/4bec2c85-2bd9-4ade-896e-459ab3e65eb7)
