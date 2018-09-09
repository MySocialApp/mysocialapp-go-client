package msa

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"bytes"
)

const (
	contentTypeJson = "application/json"
)

func NewConfiguration(appId string, clientConfiguration *ClientConfiguration) *Configuration {
	if clientConfiguration == nil {
		d := DefaultClientConfiguration
		clientConfiguration = &d
	}
	c := Configuration{
		ClientConfiguration: clientConfiguration,
		ApiEndpoint:         fmt.Sprintf("https://%s-api.mysocialapp.io/api/v1", appId),
		WebSocketEndpoint:   fmt.Sprintf("wss://%s-ws.mysocialapp.io/ws", appId),
	}
	c.Init()
	return &c
}

type Configuration struct {
	ClientConfiguration *ClientConfiguration
	ApiEndpoint         string
	WebSocketEndpoint   string
	HttpClient          *http.Client
}

func (c *Configuration) Init() {
	defaultRoundTripper := http.DefaultTransport
	defaultTransportPointer, ok := defaultRoundTripper.(*http.Transport)
	if !ok {
		panic(fmt.Sprintf("defaultRoundTripper not an *http.Transport"))
	}
	defaultTransport := *defaultTransportPointer // dereference it to get a copy of the struct that the pointer points to
	defaultTransport.MaxIdleConns = c.ClientConfiguration.ConnectionPool
	defaultTransport.MaxIdleConnsPerHost = c.ClientConfiguration.ConnectionPool
	defaultTransport.DisableKeepAlives = false

	c.HttpClient = &http.Client{
		Timeout:   c.ClientConfiguration.Timeout,
		Transport: &defaultTransport,
	}
}

func (c *Configuration) getUrlFromPath(path string) string {
	return c.ApiEndpoint + path
}

func (c *Configuration) addHeaders(request *http.Request) {
	if c.ClientConfiguration.Headers == nil {
		return
	}
	for key, value := range c.ClientConfiguration.Headers {
		request.Header.Set(key, value)
	}
}

func (c *Configuration) setContentType(request *http.Request, contentType string) {
	request.Header.Set("Content-Type", contentType)
}

func (c *Configuration) Get(path string, model ModelInterface) *RestError {
	r, err := http.NewRequest(http.MethodGet, c.getUrlFromPath(path), nil)
	if err != nil {
		return &RestError{Error: err}
	}
	c.setContentType(r, contentTypeJson)
	if restErr := c.doRequest(r, model); restErr != nil {
		return restErr
	}
	model.SetConfiguration(c)
	return nil
}

func (c *Configuration) Post(path string, model ModelInterface, body interface{}) *RestError {
	return c.push(http.MethodPost, path, model, body)
}

func (c *Configuration) Put(path string, model ModelInterface, body interface{}) *RestError {
	return c.push(http.MethodPut, path, model, body)
}

func (c *Configuration) push(method string, path string, model ModelInterface, body interface{}) *RestError {
	bodyJson, err := json.Marshal(body)
	if err != nil {
		return &RestError{Error: err}
	}
	r, err := http.NewRequest(method, c.getUrlFromPath(path), bytes.NewReader(bodyJson))
	if err != nil {
		return &RestError{Error: err}
	}
	c.setContentType(r, contentTypeJson)
	if restErr := c.doRequest(r, model); restErr != nil {
		return restErr
	}
	model.SetConfiguration(c)
	return nil
}

func (c *Configuration) doRequest(r *http.Request, model interface{}) *RestError {
	c.addHeaders(r)
	resp, err := c.HttpClient.Do(r)
	if err != nil {
		return &RestError{Error: err}
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &RestError{Error: err}
	}
	if resp.StatusCode >= 400 {
		var restErr RestError
		err = json.Unmarshal(body, &restErr)
		if err != nil {
			return &RestError{Error: err}
		}
		return &restErr
	}
	err = json.Unmarshal(body, &model)
	if err != nil {
		return &RestError{Error: err}
	}
	return nil
}
