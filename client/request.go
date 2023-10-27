package client

import (
	"io"

	http "github.com/bogdanfinn/fhttp"
)

func Request(url string, proxy string, headers map[string]string) []byte {
	client := spawn(proxy)

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil
	}

	request.Header = http.Header{}
	for key, value := range headers {
		request.Header.Add(key, value)
	}

	resp, err := client.Do(request)
	if err != nil {
		return nil
	}

	defer resp.Body.Close()

	byteResponse, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil
	}

	if resp.StatusCode != 200 {
		return nil
	}

	return byteResponse
}
