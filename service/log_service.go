package service

import (
	"encoding/json"
	"os"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

var (
	history      []map[string]interface{} // Slice to store log entries
	historyMutex sync.Mutex               // Mutex for thread safety
	logFilePath  = "assets/history.json"  // Path to the log file
)

// Initialize logs by loading existing log data from the file
func init() {
	loadExistingLogs()
}

// AddLog adds a new log entry to the history array and logs it to stdout for debugging
func AddLog(logFields logrus.Fields, level string, message string) {
	historyMutex.Lock()
	defer historyMutex.Unlock()

	// Create a structured log entry
	logEntry := logFields
	logEntry["msg"] = message
	logEntry["level"] = level
	logEntry["time"] = time.Now().Format(time.RFC3339) // Use ISO 8601 format

	// Add the log entry to the history
	history = append(history, logEntry)

	// Log the entry to stdout with the specified log level
	logger := logrus.WithFields(logFields)
	switch level {
	case "info":
		logger.Info(message)
	case "warn":
		logger.Warn(message)
	case "error":
		logger.Error(message)
	default:
		logger.Info(message) // Default to "info" if an invalid level is passed
	}
}


// SaveLog writes the history array to the JSON file
func SaveLog() {
	historyMutex.Lock()
	defer historyMutex.Unlock()

	file, err := os.Create(logFilePath) // Overwrites the existing file
	if err != nil {
		logrus.Fatalf("Failed to create log file: %v", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // Pretty-print JSON
	if err := encoder.Encode(history); err != nil {
		logrus.Fatalf("Failed to write log file: %v", err)
	}
}

// GetLog retrieves the history array
func GetLog() []map[string]interface{} {
	historyMutex.Lock()
	defer historyMutex.Unlock()
	return history
}

// loadExistingLogs loads logs from the JSON file into the history array
func loadExistingLogs() {
	// Check if the file exists
	if _, err := os.Stat(logFilePath); os.IsNotExist(err) {
		return // No file to load
	}

	// Open the existing log file
	file, err := os.Open(logFilePath)
	if err != nil {
		logrus.Fatalf("Failed to open existing log file: %v", err)
	}
	defer file.Close()

	// Decode the existing logs into the history array
	if err := json.NewDecoder(file).Decode(&history); err != nil {
		logrus.Warnf("Failed to load existing logs: %v", err)
	}
}
