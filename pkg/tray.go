package pkg

import (
	"fmt"
	"time"

	"github.com/caseymrm/menuet"
)

var latestStatusList []Component

func RunTray() error {

	// Define the Menu View ONCE
	menuet.App().Label = "com.stackit.status"

	menuet.App().Children = func() []menuet.MenuItem {
		// This function now just reads the global variable
		return buildMenu(latestStatusList)
	}

	// 3. Start the background updater
	go startStatusLoop("1s5n5g5wh9fr")

	// 4. Start the UI (Blocks Main)
	menuet.App().RunApplication()
	return nil
}

func startStatusLoop(id string) {
	for {
		err, statusList, overallStatus := GetStatus(id)

		if err != nil {
			// On error, show a warning in the bar
			menuet.App().SetMenuState(&menuet.MenuState{Title: "Stackit: ❓️"})
			time.Sleep(10 * time.Second)
			continue
		}

		// Update the global data
		latestStatusList = statusList

		// Calculate overall status for the Title Bar (Tick or Warning)
		overallEmoji := "✅"
		if overallStatus != "none" {
			overallEmoji = "⚠️"
		}

		// Update the Title Bar text
		menuet.App().SetMenuState(&menuet.MenuState{
			Title: fmt.Sprintf("Stackit: %s", overallEmoji),
		})

		time.Sleep(60 * time.Second)
	}
}

// buildMenu is a helper to keep the logic clean
func buildMenu(list []Component) []menuet.MenuItem {
	var items []menuet.MenuItem

	// Header
	items = append(items, menuet.MenuItem{
		Text:       "Stackit Status",
		FontWeight: menuet.WeightBold,
	})
	items = append(items, menuet.MenuItem{Type: menuet.Separator})

	// List Components
	for _, component := range list {
		emoji := "❓"
		switch component.Status {
		case "operational":
			emoji = "✅"
		case "degraded_performance":
			emoji = "⚠️"
		case "partial_outage":
			emoji = "🟠"
		case "major_outage":
			emoji = "🔥"
		case "maintenance":
			emoji = "🛠️"
		}

		items = append(items, menuet.MenuItem{
			Text: fmt.Sprintf("%s: %s", component.Name, emoji),
		})
	}

	return items
}
