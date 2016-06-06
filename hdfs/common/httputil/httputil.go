package httputil

import (
	"io/ioutil"
	"net/http"
)

const (
	MethodPut = "PUT"
)

func NewClient() *http.Client {
	ret := &http.Client{
		Transport: &http.Transport{
			DisableKeepAlives: true,
		},
	}

	return ret
}

func CloseResponse(resp *http.Response) error {
	if resp == nil {
		return nil
	}
	if resp.Body == nil {
		return nil
	}

	return resp.Body.Close()
}

func ReadCloseResponseBody(resp *http.Response) ([]byte, error) {
	defer CloseResponse(resp)

	body := resp.Body
	return ioutil.ReadAll(body)
}

func ReadCloseResponseBodyAsString(resp *http.Response) (string, error) {
	bytes, err := ReadCloseResponseBody(resp)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
