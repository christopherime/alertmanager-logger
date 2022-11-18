package main

import "time"

type AlertManagerNotificationObject struct {
	Receiver    string        `json:"receiver"`
	Status      string        `json:"status"`
	Alerts      []AlertObject `json:"alerts"`
	GroupLabels struct {
		Alertname string `json:"alertname"`
		Job       string `json:"job"`
	} `json:"groupLabels"`
	CommonLabels struct {
		Alertname string `json:"alertname"`
		Dc        string `json:"dc"`
		Instance  string `json:"instance"`
		Job       string `json:"job"`
	} `json:"commonLabels"`
	CommonAnnotations struct {
		Description string `json:"description"`
	} `json:"commonAnnotations"`
	ExternalURL string `json:"externalURL"`
	Version     string `json:"version"`
	GroupKey    string `json:"groupKey"`
}

type AlertObject struct {
	Status string `json:"status"`
	Labels struct {
		Alertname string `json:"alertname"`
		Dc        string `json:"dc"`
		Instance  string `json:"instance"`
		Job       string `json:"job"`
	} `json:"labels"`
	Annotations struct {
		Description string `json:"description"`
		Summary     string `json:"summary"`
	} `json:"annotations"`
	StartsAt     time.Time `json:"startsAt"`
	EndsAt       time.Time `json:"endsAt"`
	GeneratorURL string    `json:"generatorURL"`
}
