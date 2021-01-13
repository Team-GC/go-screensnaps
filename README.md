# go-screensnaps

A simple NodeJS library to interact with the APIs for screenshot generation on screensnaps.io.

## Documentation

TODO: link godoc when public

## Example Usage

```
func main() {
	// create a new client with your userID and apiKey provided in your account dashboard
	userID := "abc-123"
	apiKey := "def-456"
	screensnaps := screensnaps.NewScreensnapsClient(userID, apiKey)

	// create a new screenshot from a url
	targetURL := "https://golang.org"
	screenshot, exception, err := screensnaps.CreateScreenshot(targetURL)

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