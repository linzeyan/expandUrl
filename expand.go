package expandUrl

import (
	"net/http"
	"net/url"
)

/* Expand returns a real URL. */
func Expand(uri string) (string, error) {
	_, err := url.ParseRequestURI(uri)
	if err != nil {
		return "", err
	}

	var result = uri
	var client = &http.Client{
		Transport: &http.Transport{
			DisableKeepAlives: true,
		},
		CheckRedirect: func(req *http.Request, _ []*http.Request) error {
			result = req.URL.String()
			return nil
		},
	}

	req, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return "", err
	}

	resp, err := client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return "", err
	}
	if req.Host == "reurl.cc" && resp.Header.Get("Target") != "" {
		result = resp.Header.Get("Target")
	}

	return result, nil
}
