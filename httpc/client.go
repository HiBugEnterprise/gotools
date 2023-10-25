package httpc

import (
	"bytes"
	"github.com/HiBugEnterprise/gotools/errorx"
	"github.com/go-resty/resty/v2"
	"io"
	"k8s.io/apimachinery/pkg/util/sets"
	"net/http"
	"os"
	"time"
)

const (
	UserAgent      = "HiBug Client"
	TimeoutSeconds = 60
	MaxRetryCount  = 3
)

type Client struct {
	*resty.Client

	Host    string // Host is the fully qualified domain name of the system, or an IP Address. Port and protocol are required if necessary.
	BaseURI string // BaseURI is the base uri for every request, starting with a slash, for example: /api/v1
	//IgnoreCodes sets.Int // IgnoreCodes ignores some code to be returned as an error.
	IgnoreCodes sets.Set[int]
}

func Get(url string, rfs ...RequestFunc) (*resty.Response, error) {
	return New().Get(url, rfs...)
}

func Post(url string, rfs ...RequestFunc) (*resty.Response, error) {
	return New().Post(url, rfs...)
}

func Patch(url string, rfs ...RequestFunc) (*resty.Response, error) {
	return New().Patch(url, rfs...)
}

func Put(url string, rfs ...RequestFunc) (*resty.Response, error) {
	return New().Put(url, rfs...)
}

func Delete(url string, rfs ...RequestFunc) (*resty.Response, error) {
	return New().Delete(url, rfs...)
}

func Head(url string, rfs ...RequestFunc) (*resty.Response, error) {
	return New().Head(url, rfs...)
}

func Options(url string, rfs ...RequestFunc) (*resty.Response, error) {
	return New().Options(url, rfs...)
}

// Download retrieves content from the given url and write it to path.
func Download(url, path string, rfs ...RequestFunc) error {
	// download may take more time
	cl := New(UnsetTimeout())
	res, err := cl.Get(url, rfs...)
	if err != nil {
		return err
	}

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer func() {
		_ = f.Close()
	}()

	_, err = io.Copy(f, bytes.NewReader(res.Body()))

	return err
}

func New(cfs ...ClientFunc) *Client {
	r := resty.New()
	r.SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetHeader("User-Agent", UserAgent).
		SetTimeout(TimeoutSeconds * time.Second).
		SetRetryCount(MaxRetryCount)

	c := &Client{
		Client:      r,
		IgnoreCodes: sets.New[int](),
	}

	for _, cf := range cfs {
		cf(c)
	}
	return c
}

func (c *Client) Get(url string, rfs ...RequestFunc) (*resty.Response, error) {
	return c.Request(resty.MethodGet, url, rfs...)
}

func (c *Client) Post(url string, rfs ...RequestFunc) (*resty.Response, error) {
	return c.Request(resty.MethodPost, url, rfs...)
}

func (c *Client) Patch(url string, rfs ...RequestFunc) (*resty.Response, error) {
	return c.Request(resty.MethodPatch, url, rfs...)
}

func (c *Client) Put(url string, rfs ...RequestFunc) (*resty.Response, error) {
	return c.Request(resty.MethodPut, url, rfs...)
}

func (c *Client) Delete(url string, rfs ...RequestFunc) (*resty.Response, error) {
	return c.Request(resty.MethodDelete, url, rfs...)
}

func (c *Client) Head(url string, rfs ...RequestFunc) (*resty.Response, error) {
	return c.Request(resty.MethodHead, url, rfs...)
}

func (c *Client) Options(url string, rfs ...RequestFunc) (*resty.Response, error) {
	return c.Request(resty.MethodOptions, url, rfs...)
}

func (c *Client) Request(method, url string, rfs ...RequestFunc) (*resty.Response, error) {
	if c.BaseURI != "" {
		url = c.BaseURI + url
	}
	r := c.R()

	for _, rf := range rfs {
		rf(r)
	}

	return c.wrapError(r.Execute(method, url))
}

func (c *Client) wrapError(res *resty.Response, err error) (*resty.Response, error) {
	if err != nil {
		return res, err
	}
	code := res.StatusCode()
	if res.IsError() && !c.IgnoreCodes.Has(code) {
		return res, errorx.New(http.StatusText(code), code, res.String()).
			WithMetadata(errorx.Metadata{"err": res.Error()})
	}

	return res, nil
}
