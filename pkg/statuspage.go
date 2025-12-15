package pkg

import (
	"encoding/json"
	"net/http"
	"time"
)

type StatusPageResponse struct {
	Page struct {
		Id        string    `json:"id"`
		Name      string    `json:"name"`
		Url       string    `json:"url"`
		TimeZone  string    `json:"time_zone"`
		UpdatedAt time.Time `json:"updated_at"`
	} `json:"page"`
	Components []struct {
		Id                 string    `json:"id"`
		Name               string    `json:"name"`
		Status             string    `json:"status"`
		CreatedAt          time.Time `json:"created_at"`
		UpdatedAt          time.Time `json:"updated_at"`
		Position           int       `json:"position"`
		Description        *string   `json:"description"`
		Showcase           bool      `json:"showcase"`
		StartDate          *string   `json:"start_date"`
		GroupId            *string   `json:"group_id"`
		PageId             string    `json:"page_id"`
		Group              bool      `json:"group"`
		OnlyShowIfDegraded bool      `json:"only_show_if_degraded"`
		Components         []string  `json:"components,omitempty"`
	} `json:"components"`
	Incidents []struct {
		Id              string      `json:"id"`
		Name            string      `json:"name"`
		Status          string      `json:"status"`
		CreatedAt       time.Time   `json:"created_at"`
		UpdatedAt       time.Time   `json:"updated_at"`
		MonitoringAt    time.Time   `json:"monitoring_at"`
		ResolvedAt      interface{} `json:"resolved_at"`
		Impact          string      `json:"impact"`
		Shortlink       string      `json:"shortlink"`
		StartedAt       time.Time   `json:"started_at"`
		PageId          string      `json:"page_id"`
		IncidentUpdates []struct {
			Id                 string    `json:"id"`
			Status             string    `json:"status"`
			Body               string    `json:"body"`
			IncidentId         string    `json:"incident_id"`
			CreatedAt          time.Time `json:"created_at"`
			UpdatedAt          time.Time `json:"updated_at"`
			DisplayAt          time.Time `json:"display_at"`
			AffectedComponents []struct {
				Code      string `json:"code"`
				Name      string `json:"name"`
				OldStatus string `json:"old_status"`
				NewStatus string `json:"new_status"`
			} `json:"affected_components"`
			DeliverNotifications bool        `json:"deliver_notifications"`
			CustomTweet          interface{} `json:"custom_tweet"`
			TweetId              interface{} `json:"tweet_id"`
		} `json:"incident_updates"`
		Components []struct {
			Id                 string      `json:"id"`
			Name               string      `json:"name"`
			Status             string      `json:"status"`
			CreatedAt          time.Time   `json:"created_at"`
			UpdatedAt          time.Time   `json:"updated_at"`
			Position           int         `json:"position"`
			Description        interface{} `json:"description"`
			Showcase           bool        `json:"showcase"`
			StartDate          string      `json:"start_date"`
			GroupId            string      `json:"group_id"`
			PageId             string      `json:"page_id"`
			Group              bool        `json:"group"`
			OnlyShowIfDegraded bool        `json:"only_show_if_degraded"`
		} `json:"components"`
		ReminderIntervals interface{} `json:"reminder_intervals"`
	} `json:"incidents"`
	ScheduledMaintenances []struct {
		Id              string      `json:"id"`
		Name            string      `json:"name"`
		Status          string      `json:"status"`
		CreatedAt       time.Time   `json:"created_at"`
		UpdatedAt       time.Time   `json:"updated_at"`
		MonitoringAt    interface{} `json:"monitoring_at"`
		ResolvedAt      interface{} `json:"resolved_at"`
		Impact          string      `json:"impact"`
		Shortlink       string      `json:"shortlink"`
		StartedAt       time.Time   `json:"started_at"`
		PageId          string      `json:"page_id"`
		IncidentUpdates []struct {
			Id                 string    `json:"id"`
			Status             string    `json:"status"`
			Body               string    `json:"body"`
			IncidentId         string    `json:"incident_id"`
			CreatedAt          time.Time `json:"created_at"`
			UpdatedAt          time.Time `json:"updated_at"`
			DisplayAt          time.Time `json:"display_at"`
			AffectedComponents []struct {
				Code      string `json:"code"`
				Name      string `json:"name"`
				OldStatus string `json:"old_status"`
				NewStatus string `json:"new_status"`
			} `json:"affected_components"`
			DeliverNotifications bool        `json:"deliver_notifications"`
			CustomTweet          interface{} `json:"custom_tweet"`
			TweetId              interface{} `json:"tweet_id"`
		} `json:"incident_updates"`
		Components []struct {
			Id                 string      `json:"id"`
			Name               string      `json:"name"`
			Status             string      `json:"status"`
			CreatedAt          time.Time   `json:"created_at"`
			UpdatedAt          time.Time   `json:"updated_at"`
			Position           int         `json:"position"`
			Description        interface{} `json:"description"`
			Showcase           bool        `json:"showcase"`
			StartDate          string      `json:"start_date"`
			GroupId            string      `json:"group_id"`
			PageId             string      `json:"page_id"`
			Group              bool        `json:"group"`
			OnlyShowIfDegraded bool        `json:"only_show_if_degraded"`
		} `json:"components"`
		ScheduledFor   time.Time `json:"scheduled_for"`
		ScheduledUntil time.Time `json:"scheduled_until"`
	} `json:"scheduled_maintenances"`
	Status struct {
		Indicator   string `json:"indicator"`
		Description string `json:"description"`
	} `json:"status"`
}

type Component struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

func GetStatus(id string) (error, []Component, string) {
	var data StatusPageResponse
	var currentStatus []Component
	var overallStatus string
	resp, err := http.Get("https://" + id + ".statuspage.io/api/v2/summary.json")
	if err != nil {
		return err, currentStatus, overallStatus
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return err, currentStatus, overallStatus
	}

	for _, component := range data.Components {
		currentStatus = append(currentStatus, Component{
			Name:   component.Name,
			Status: component.Status,
		})
	}
	overallStatus = data.Status.Indicator

	return nil, currentStatus, overallStatus
}
