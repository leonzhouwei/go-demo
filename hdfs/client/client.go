package client

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/leonzhouwei/go-demo/hdfs/common/httputil"

	qerrors "github.com/qiniu/errors"
	"github.com/qiniu/log.v1"
)

var _ Client = &HTTPClient{}

type Client interface {
	GetHomeDir() (string, error)
	LsDir(path string) (string, error)
	WriteAndClose(path string, data io.Reader) error
}

type Config struct {
	Host      string
	Port      int
	APIPrefix string
	Username  string
}

type HTTPClient struct {
	conf Config
	http *http.Client
}

func NewHTTPClient(config Config) *HTTPClient {
	if config.Host == "" {
		config.Host = "127.0.0.1"
	}
	if config.Port == 0 {
		config.Port = 50070
	}
	if config.APIPrefix == "" {
		config.APIPrefix = "/webhdfs/v1"
	}

	ret := &HTTPClient{
		conf: config,
		http: httputil.NewClient(),
	}

	return ret
}

func (r *HTTPClient) NewURLString(path, op string) string {
	return fmt.Sprintf(
		"http://%s:%d%s%s?user.name=%s&op=%s",
		r.conf.Host,
		r.conf.Port,
		r.conf.APIPrefix,
		path,
		r.conf.Username,
		url.QueryEscape(op),
	)
}

func (r *HTTPClient) GetHomeDir() (string, error) {
	url := r.NewURLString("/", "GETHOMEDIRECTORY")
	resp, err := r.http.Get(url)
	if err != nil {
		return "", err
	}

	return httputil.ReadCloseResponseBodyAsString(resp)
}

func (r *HTTPClient) LsDir(path string) (string, error) {
	url := r.NewURLString(path, "LISTSTATUS")
	resp, err := r.http.Get(url)
	if err != nil {
		return "", err
	}

	return httputil.ReadCloseResponseBodyAsString(resp)
}

func (r *HTTPClient) WriteAndClose(path string, body io.Reader) error {
	url1 := r.NewURLString(path, "CREATE")
	// Step 1:
	// Submit a HTTP PUT request without automatically following
	// redirects and without sending the file data.
	req1, err := http.NewRequest(httputil.MethodPut, url1, nil)
	if err != nil {
		err = errors.New(qerrors.Info(err).LogMessage())
		return err
	}
	resp1, err := r.http.Do(req1)
	if err != nil {
		err = errors.New(qerrors.Info(err).LogMessage())
		return err
	}
	if resp1.StatusCode/100 != 3 {
		err := fmt.Errorf("oops: %v", resp1.StatusCode)
		err = errors.New(qerrors.Info(err).LogMessage())
		return err
	}

	// Step 2:
	// Submit another HTTP PUT request using the URL in
	// the Location header with the file data to be written.
	url2 := resp1.Header.Get("Location")
	log.Debug(url2)
	req2, err := http.NewRequest(httputil.MethodPut, url2, body)
	if err != nil {
		err = errors.New(qerrors.Info(err).LogMessage())
		return err
	}
	resp2, err2 := r.http.Do(req2)
	if err2 != nil {
		err2 = errors.New(qerrors.Info(err).LogMessage())
		return err2
	}
	log.Debug(resp2.StatusCode)
	if resp2.StatusCode != 201 {
		err := fmt.Errorf("%v", resp2.StatusCode)
		err = errors.New(qerrors.Info(err).LogMessage())
		return err
	}

	return nil
}
