package screensnaps

import (
	"os"
	"testing"
)

var userID = ""
var apiKey = ""

func init() {
	userID = os.Getenv("USER_ID")
	apiKey = os.Getenv("API_KEY")
}

func TestGetStatus(t *testing.T) {
	screensnaps := NewScreensnapsClient(userID, apiKey)

	_, exc, err := screensnaps.GetStatus()

	if err != nil {
		t.Fatal(err)
	}

	if exc != nil {
		t.Fatal(exc)
	}
}

func TestGetScreenshots(t *testing.T) {
	screensnaps := NewScreensnapsClient(userID, apiKey)

	_, exc, err := screensnaps.GetScreenshots()

	if err != nil {
		t.Fatal(err)
	}

	if exc != nil {
		t.Fatal(exc)
	}
}

func TestCreateScreenshotFromURL(t *testing.T) {
	screensnaps := NewScreensnapsClient(userID, apiKey)

	targetURL := "https://golang.org"
	_, exc, err := screensnaps.CreateScreenshotFromURL(targetURL)

	if err != nil {
		t.Fatal(err)
	}

	if exc != nil {
		t.Fatal(exc)
	}
}

func CreateScreenshotFromHTML(t *testing.T) {
	screensnaps := NewScreensnapsClient(userID, apiKey)

	targetHTML := "<html><body><p>I love Screensnaps!</p></body></html>"
	_, exc, err := screensnaps.CreateScreenshotFromHTML(targetHTML)

	if err != nil {
		t.Fatal(err)
	}

	if exc != nil {
		t.Fatal(exc)
	}
}
