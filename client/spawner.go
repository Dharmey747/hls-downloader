package client

import (
	"log"

	tls_client "github.com/bogdanfinn/tls-client"
)

func spawn(proxy string) tls_client.HttpClient {
	options := []tls_client.HttpClientOption{
		tls_client.WithTimeoutSeconds(30),
	}

	if proxy != "" {
		options = append(options, tls_client.WithProxyUrl("http://"+proxy))
	}

	client, err := tls_client.NewHttpClient(tls_client.NewNoopLogger(), options...)
	if err != nil {
		log.Println(err)
		return nil
	}

	return client
}
