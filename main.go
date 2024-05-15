package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/prometheus/alertmanager/template"
)

type AlertManagerNotificationObject struct {
	Version           string                 `json:"version"`

	GroupKey          string                 `json:"groupKey"`
	Status            string                 `json:"status"`
	Receiver          string                 `json:"receiver"`
	GroupLabels       map[string]string      `json:"groupLabels"`
	CommonLabels      map[string]string      `json:"commonLabels"`
	CommonAnnotations map[string]string      `json:"commonAnnotations"`
	ExternalURL       string                 `json:"externalURL"`
	Alerts            []template.Alert       `json:"alerts"`
}

var (
	logFileMutex sync.Mutex
)

func main() {

	// Start the log file rotation.
	startLogFileRotation()

	// Create a new HTTP server.
	http.HandleFunc("/logger", handleAlert)
	http.HandleFunc("/", handleAllOtherRequests)

	// Start the HTTP server.
	log.Println("Starting Alert Manager Logger on port 9095")
	if err := http.ListenAndServe(":9095", nil); err != nil {
		log.Fatal(err)
	}
}

func handleAlert(w http.ResponseWriter, r *http.Request) {
	// Parse the alert notification body.
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var alertManagerMessage AlertManagerNotificationObject
	if err := json.Unmarshal(body, &alertManagerMessage); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Log the alert notification.
	log.Printf("Alert value: %v", alertManagerMessage)
	writeAlertLog(alertManagerMessage)
	log.Println("Alert received")
	log.Println(string(body))
	// Return a success response.
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Notification received"))
}
func handleAllOtherRequests(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusForbidden)
	w.Write([]byte("Forbidden"))
}
// rotateLogFile rotates the log file every day, but only if the file is not currently being written to.
func rotateLogFile() {
	logFileMutex.Lock()
	defer logFileMutex.Unlock()
	// Get the current date.
	now := time.Now()
	// Create a new log file name with the current date.
	newLogFileName := filepath.Join("/var/log/amlogger", "alerts-"+now.Format("2006-01-02")+".log")
	// Open the new log file.
	f, err := os.OpenFile(newLogFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	// Close the old log file.
	log.SetOutput(f)
}
