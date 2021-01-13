# go-screensnaps

![Build](https://github.com/Team-GC/go-screensnaps/workflows/Go/badge.svg)

A simple NodeJS library to interact with the APIs for screenshot generation on screensnaps.io.

## Installation

To install go-screensnaps: 

`go get github.com/Team-GC/go-screensnaps`

## Documentation

TODO: link godoc when public

## Tests

Add your keys as environmental variables:

`USER_ID=abc-123` and `API_KEY=def-456`

Then run:

`go test`

## Example Usage

```
func main() {
	// create a new client with your userID and apiKey provided in your account dashboard
	userID := "abc-123"
	apiKey := "def-456"
	screensnaps := screensnaps.NewScreensnapsClient(userID, apiKey)

	// create a new screenshot from a url
	targetURL := "https://golang.org"
	screenshot, exception, err := screensnaps.CreateScreenshotFromURL(targetURL)

	if err != nil {
		println("error: " + err.Error())
	}

	if exception != nil {
		println("exception: " + exception.Status)
	}

	println(screenshot.ImageURL)

	// create a new screenshot from HTML
	targetHTML := "<html><body><p>I love Screensnaps!</p></body></html>"
	screenshot, exception, err = screensnaps.CreateScreenshotFromHTML(targetHTML)

	if err != nil {
		println("error: " + err.Error())
	}

	if exception != nil {
		println("exception: " + exception.Status)
	}

	println(screenshot.ImageURL)

	// get the last 15 screenshots generated with your account
	screenshots, exception, err := screensnaps.GetScreenshots()

	if err != nil {
		println("error: " + err.Error())
	}

	if exception != nil {
		println("exception: " + exception.Status)
	}

	println(screenshots.Items[0].ImageURL)

	// get the status of the API
	status, exception, err := screensnaps.GetStatus()

	if err != nil {
		println("error: " + err.Error())
	}

	if exception != nil {
		println("exception: " + exception.Status)
	}

	println(status.Status)
}
```