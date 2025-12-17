# stackitStatus
A lightweight macOS menu bar application that monitors the status of STACKIT services.

## What it does
Menu Bar Status: Shows Stackit: ✅ directly in your menu bar for at-a-glance status.

Detailed View: Click the menu item to see a breakdown of every component (Portal, API, Network, etc.) and their specific
status.

Auto-Refresh: Updates automatically every 60 seconds.

## Installation
Go to the [Releases](https://github.com/RiRa12621/stackitStatus/releases) page.

Download the latest .dmg file.

Open the .dmg and drag stackitStatus.app to your Applications folder.

Run the app.

**Note**: On the first run, you might need to right-click the app and select "Open" to bypass macOS security checks for
unsigned applications.

## Disclaimer
⚠️ **Unofficial Project**: This tool is a community project and is not an official product of STACKIT.

⚠️ **Beta Quality**: This software is provided "as is" and is not fully production-ready. Use it at your own discretion.

## Development
To build the app locally:

```Bash
# Clone the repo
git clone https://github.com/rira12621/stackitStatus.git
cd stackitStatus

# Run directly
go run main.go


```