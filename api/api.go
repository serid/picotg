package api

import (
	"bytes"
	"io"
	"net/http"
	"time"
)

type Api struct {
	token  string
	client *http.Client
}

func NewApi(token string) *Api {
	return &Api{token: token,
		client: &http.Client{Transport: &http.Transport{
			MaxIdleConns:       10,
			IdleConnTimeout:    30 * time.Second,
			DisableCompression: true,
		}}}
}

func (api *Api) Request(method string, body []byte) ([]byte, error) {
	bytesReader := bytes.NewReader(body)

	req, err := http.NewRequest("POST", "https://api.telegram.org/bot"+api.token+"/"+method, bytesReader)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")

	resp, err := api.client.Do(req)

	if err != nil {
		return nil, err
	}

	r, err := ReadBytes(resp.Body)

	if err != nil {
		return nil, err
	}

	err = resp.Body.Close()

	if err != nil {
		return nil, err
	}

	return r, nil
}

func ReadBytes(reader io.Reader) ([]byte, error) {
	buf := new(bytes.Buffer)

	_, err := buf.ReadFrom(reader)

	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
