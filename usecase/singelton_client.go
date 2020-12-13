package usecase

import (
	"http-client/gohttp"
	"time"
)
var (
	httpClient = getHttpClient()
)
func getHttpClient() gohttp.Client  {
	client := gohttp.NewBuilder().
		SetConnectionTimeout(2 * time.Second).
		Build()
	return client
}
