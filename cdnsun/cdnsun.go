package cdnsun

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type (
	ApiClient struct {
		username string
		password string
	}

	Options struct {
		Method string
		Url    string
		Data   map[string]interface{}
	}
)

func New(username string, password string) (*ApiClient, error) {
	if len(strings.TrimSpace(username)) == 0 {
		return nil, errors.New("username empty")
	}
	if len(strings.TrimSpace(password)) == 0 {
		return nil, errors.New("password empty")
	}

	client := ApiClient{username, password}
	return &client, nil
}

func (client *ApiClient) Get(options *Options) ([]byte, error) {
	if options == nil {
		return nil, errors.New("options empty")
	}
	options.Method = "GET"
	return client.request(options)
}

func (client *ApiClient) Post(options *Options) ([]byte, error) {
	if options == nil {
		return nil, errors.New("options empty")
	}
	options.Method = "POST"
	return client.request(options)
}

func (client *ApiClient) Put(options *Options) ([]byte, error) {
	if options == nil {
		return nil, errors.New("options empty")
	}
	options.Method = "PUT"
	return client.request(options)
}

func (client *ApiClient) Delete(options *Options) ([]byte, error) {
	if options == nil {
		return nil, errors.New("options empty")
	}
	options.Method = "DELETE"
	return client.request(options)
}

func (client *ApiClient) request(options *Options) (result []byte, err error) {

	if options == nil {
		return result, errors.New("options empty")
	}

	if options.Url == "" {
		return result, errors.New("options.Url empty")
	}

	if options.Method == "" {
		return result, errors.New("options.Method empty")
	}

	httpClient := &http.Client{Timeout: 60 * time.Second}

	var request *http.Request
	path := "https://cdnsun.com/api/"

	switch options.Method {
	case "POST", "PUT", "DELETE":
		jsonData := []byte{}
		if options.Data != nil {
			jsonData, err = json.Marshal(options.Data)
			if err != nil {
				return result, errors.New("Invalid options.Data format")
			}
		}
		request, err = http.NewRequest(options.Method, path+options.Url, bytes.NewBuffer(jsonData))
		if err != nil {
			return result, err
		}
	case "GET":
		request, err = http.NewRequest(options.Method, path+options.Url, nil)
		if err != nil {
			return result, err
		}
		query := request.URL.Query()
		for key, value := range options.Data {
			query.Add(key, value.(string))
		}
		request.URL.RawQuery = query.Encode()
	default:
		return result, errors.New("Unsupported method: " + options.Method)
	}

	request.Header.Add("Accept", "application/json")
	request.Header.Add("Content-Type", "application/json")
	request.SetBasicAuth(client.username, client.password)

	response, err := httpClient.Do(request)
	defer response.Body.Close()

	return ioutil.ReadAll(response.Body)
}
