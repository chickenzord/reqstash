package reqstash

import (
	"io/ioutil"
	"net/http"
	"time"
)

type Request struct {
	Timestamp time.Time
	Method    string
	Host      string
	URL       string
	Header    map[string][]string
	Error     error
	Body      string
}

func NewRequest(r *http.Request) Request {
	req := Request{
		Timestamp: time.Now(),
		Method:    r.Method,
		Host:      r.Host,
		URL:       r.URL.String(),
		Header:    r.Header,
	}

	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		req.Error = err
	} else {
		req.Body = string(bytes)
	}

	return req
}
