package client

import (
	"encoding/json"
	"os"
	"reflect"
	"testing"

	"github.com/Netsocs-Team/driver.sdk_go/pkg/config"
)

func createTempFile(t *testing.T, content string) string {
	tmpfile, err := os.Create("events.json")
	if err != nil {
		t.Fatalf("Failed to create temp file: %s", err)
	}

	if _, err := tmpfile.Write([]byte(content)); err != nil {
		t.Fatalf("Failed to write to temp file: %s", err)
	}

	if err := tmpfile.Close(); err != nil {
		t.Fatalf("Failed to close temp file: %s", err)
	}

	return tmpfile.Name()
}

func removeTempFile(t *testing.T, filename string) {
	if err := os.Remove(filename); err != nil {
		t.Fatalf("Failed to remove temp file: %s", err)
	}
}

func TestLoadsEventsFromFile_Success(t *testing.T) {
	expectedEvents := config.GetEventsAvailableResponse{
		{
			Key: "key1",
			TranslationStrings: map[string]config.EventSchemaTranslationStrings{
				"en": {
					Name:        "Event 1",
					Description: "Description 1",
				},
			},
			Priority: 1,
		},
	}

	content, err := json.Marshal(expectedEvents)
	if err != nil {
		t.Fatalf("Failed to marshal JSON: %s", err)
	}

	filename := createTempFile(t, string(content))
	defer removeTempFile(t, filename)

	events, err := loadsEventsFromFile(filename)
	if err != nil {
		t.Fatalf("Expected no error, got %s", err)
	}

	if !reflect.DeepEqual(events, expectedEvents) {
		t.Errorf("Expected %v, got %v", expectedEvents, events)
	}
}

func TestLoadsEventsFromFile_FileNotFound(t *testing.T) {
	_, err := loadsEventsFromFile("nonexistentfile.json")
	if err == nil {
		t.Fatalf("Expected error, got nil")
	}
}

func TestLoadsEventsFromFile_InvalidJSON(t *testing.T) {
	filename := createTempFile(t, "{invalid json}")
	defer removeTempFile(t, filename)

	_, err := loadsEventsFromFile(filename)
	if err == nil {
		t.Fatalf("Expected error, got nil")
	}
}
