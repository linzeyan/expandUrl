package expandUrl

import (
	"net/http"
)

func Expand(url string) (string, error) {
	const ua = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36"
	var result string
	result = url

	var client = &http.Client{
		CheckRedirect: func(req *http.Request, _ []*http.Request) error {
			result = req.URL.String()
			return nil
		},
		Transport: &http.Transport{
			DisableKeepAlives: true,
		},
	}

	var req, err = http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("User-Agent", ua)

	resp, err := client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return "", err
	}

	return result, nil
}
