package screensnaps

import "net/http"

// Configuration type
type Configuration struct {
	baseURL    string
	apiVersion string
}

// Screensnaps type
type Screensnaps struct {
	userID     string
	apiKey     string
	config     *Configuration
	HTTPClient *http.Client
}

// SnapParams type
type SnapParams struct {
	url               *string
	html              *string
	height            *int
	width             *int
	pageTarget        *string
	fullPage          *int
	deviceScaleFactor *int
	apiKey            *string
}

// SnapResponse type
type SnapResponse struct {
	Status   string `json:"status"`
	ImageURL string `json:"image_url"`
	Metrics  *ScreenshotMetric
}

// ScreenshotMetric type
type ScreenshotMetric struct {
	Launch     int `json:"launch"`
	Browser    int `json:"browser"`
	Age        int `json:"age"`
	Screenshot int `json:"screenshot"`
	Target     int `json:"target"`
	Upload     int `json:"upload"`
	Total      int `json:"total"`
}

// SnapsParams type
type SnapsParams struct {
	offset *int
	limit  *int
}

// SnapsResponse type
type SnapsResponse struct {
	Status string `json:"status"`
	Items  []Screenshot
}

// Screenshot type
type Screenshot struct {
	ID       string `json:"id"`
	Mode     string `json:"mode"`
	ImageURL string `json:"image_url"`
}

// SnapStatusResponse type
type SnapStatusResponse struct {
	Status string `json:"status"`
}

// Exception type
type Exception struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
