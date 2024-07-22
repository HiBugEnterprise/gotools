package jeepay_go_sdk

import "net/http"

// Configuration stores the configuration of the API client
type Configuration struct {
	Host string `json:"host,omitempty"`
	// 请求协议
	Scheme        string            `json:"scheme,omitempty"`
	DefaultHeader map[string]string `json:"defaultHeader,omitempty"`
	UserAgent     string            `json:"userAgent,omitempty"`
	Debug         bool              `json:"debug,omitempty"`
	// jeepay中的
	AppId      string `json:"appId,omitempty"`
	AppSecret  string `json:"appSecret,omitempty"`
	SignType   string `json:"signType,omitempty"`
	HTTPClient *http.Client
}

// NewConfiguration returns a new Configuration object
func NewConfiguration() *Configuration {
	cfg := &Configuration{
		DefaultHeader: make(map[string]string),
		UserAgent:     "OpenAPI-Generator/2.1.0/go",
		Debug:         false,
	}
	return cfg
}

// AddDefaultHeader adds a new HTTP header to the default header in the request
func (c *Configuration) AddDefaultHeader(key string, value string) {
	c.DefaultHeader[key] = value
}

func (c *Configuration) GetLocalBasePath() (string, error) {
	if c.Host == "" {
		return "", reportError("Host is empty")
	}
	return c.Host, nil
}
