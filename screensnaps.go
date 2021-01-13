package screensnaps

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	baseURL    = "https://api.screensnaps.io"
	apiVersion = "v1"
)

var defaultClient = &http.Client{
	Timeout: time.Second * 30,
}

// NewScreensnapsClient creates a new client with a default HTTP Client using the provided userID and apiKey
func NewScreensnapsClient(userID, apiKey string) *Screensnaps {
	return NewScreensnapsClientCustomHTTP(userID, apiKey, nil)
}

// NewScreensnapsClientCustomHTTP creates a new custom client using the provided userID, apiKey, and HTTP Client
func NewScreensnapsClientCustomHTTP(userID, apiKey string, HTTPClient *http.Client) *Screensnaps {
	if HTTPClient == nil {
		HTTPClient = defaultClient
	}

	config := &Configuration{
		baseURL:    baseURL,
		apiVersion: apiVersion,
	}

	return &Screensnaps{
		userID:     userID,
		apiKey:     apiKey,
		config:     config,
		HTTPClient: HTTPClient,
	}
}

// CreateScreenshot creates a new screenshot with the provided targetURL or targetHTML and returns a SnapResponse if successful
func (screensnaps *Screensnaps) CreateScreenshot(targetURL string) (snapResponse *SnapResponse, exception *Exception, err error) {
	apiURL := screensnaps.config.baseURL + "/" + screensnaps.config.apiVersion + "/screenshot"

	values := map[string]string{"url": targetURL}

	res, err := screensnaps.post(values, apiURL)
	if err != nil {
		return snapResponse, exception, err
	}
	defer res.Body.Close()

	responseBody, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return snapResponse, exception, err
	}

	if res.StatusCode != http.StatusOK {
		exception = new(Exception)
		err = json.Unmarshal(responseBody, exception)

		return snapResponse, exception, err
	}

	snapResponse = new(SnapResponse)
	err = json.Unmarshal(responseBody, snapResponse)
	return snapResponse, exception, err
}

// GetScreenshots returns a SnapsResponse if successful, containing a list of the last 15 screenshots previously generated
func (screensnaps *Screensnaps) GetScreenshots() (snapsResponse *SnapsResponse, exception *Exception, err error) {
	apiURL := screensnaps.config.baseURL + "/" + screensnaps.config.apiVersion + "/screenshots"

	res, err := screensnaps.get(apiURL)
	if err != nil {
		return snapsResponse, exception, err
	}
	defer res.Body.Close()

	responseBody, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return snapsResponse, exception, err
	}

	if res.StatusCode != http.StatusOK {
		exception = new(Exception)
		err = json.Unmarshal(responseBody, exception)

		return snapsResponse, exception, err
	}

	snapsResponse = new(SnapsResponse)
	err = json.Unmarshal(responseBody, snapsResponse)
	return snapsResponse, exception, err
}

// GetStatus returns a SnapStatusResponse if successful, containing an API status
func (screensnaps *Screensnaps) GetStatus() (statusResponse *SnapStatusResponse, exception *Exception, err error) {
	apiURL := screensnaps.config.baseURL + "/" + screensnaps.config.apiVersion + "/status"

	res, err := screensnaps.get(apiURL)
	if err != nil {
		return statusResponse, exception, err
	}
	defer res.Body.Close()

	responseBody, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return statusResponse, exception, err
	}

	if res.StatusCode != http.StatusOK {
		exception = new(Exception)
		err = json.Unmarshal(responseBody, exception)

		return statusResponse, exception, err
	}

	statusResponse = new(SnapStatusResponse)
	err = json.Unmarshal(responseBody, statusResponse)
	return statusResponse, exception, err
}

func (screensnaps *Screensnaps) get(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", screensnaps.userID)

	q := req.URL.Query()
	q.Add("api_key", screensnaps.apiKey)
	req.URL.RawQuery = q.Encode()

	return screensnaps.do(req)
}

func (screensnaps *Screensnaps) post(values map[string]string, url string) (*http.Response, error) {
	jsonValue, _ := json.Marshal(values)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", screensnaps.userID)
	req.Header.Add("Content-Type", "application/json")

	q := req.URL.Query()
	q.Add("api_key", screensnaps.apiKey)
	req.URL.RawQuery = q.Encode()

	return screensnaps.do(req)
}

func (screensnaps *Screensnaps) do(req *http.Request) (*http.Response, error) {
	client := screensnaps.HTTPClient
	if client == nil {
		client = defaultClient
	}

	return client.Do(req)
}
