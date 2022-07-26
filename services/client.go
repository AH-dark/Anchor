package services

import (
	"net/http"
	"time"
)

var client *http.Client

func init() {
	client = http.DefaultClient
	client.Timeout = 5 * time.Second // 5秒超时时间
}
