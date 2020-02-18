package gomarta

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

const (
	ProdBaseURL  = "http://developer.itsmarta.com"
	TrainPath    = "/RealtimeTrain/RestServiceNextTrain/GetRealtimeArrivals"
	AllBusPath   = "/BRDRestService/RestBusRealTimeService/GetAllBus"
	BusRoutePath = "/BRDRestService/RestBusRealTimeService/GetBusByRoute/"
)

type HTTPClient interface {
	Get(url string) (resp *http.Response, err error)
}

type HTTPError struct {
	ResponseData []byte
	StatusCode   int
}

func (e *HTTPError) Error() string {
	return "bad status code:" + strconv.Itoa(e.StatusCode)
}

func NewDefaultClient(apiKey string) *Client {
	return &Client{
		ApiKey: apiKey,
		Host:   ProdBaseURL,
		HTTP:   &http.Client{},
	}
}

type Client struct {
	ApiKey string
	Host   string
	HTTP   HTTPClient
}

func (c *Client) GetTrains() (TrainAPIResponse, error) {
	u, err := url.Parse(c.Host + TrainPath + "?apiKey=" + c.ApiKey)
	if err != nil {
		return nil, err
	}

	buff, err := c.get(u)
	if err != nil {
		return nil, err
	}

	var r TrainAPIResponse
	err = (&r).UnmarshalJSON(buff)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (c *Client) GetAllBuses() (BusAPIResponse, error) {
	u, err := url.Parse(c.Host + AllBusPath + "?apiKey=" + c.ApiKey)
	if err != nil {
		return nil, err
	}

	buff, err := c.get(u)
	if err != nil {
		return nil, err
	}

	var r BusAPIResponse
	err = (&r).UnmarshalJSON(buff)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (c *Client) GetBusByRoute(route string) (BusAPIResponse, error) {
	u, err := url.Parse(c.Host + BusRoutePath + route + "?apiKey=" + c.ApiKey)
	if err != nil {
		return nil, err
	}

	buff, err := c.get(u)
	if err != nil {
		return nil, err
	}

	var r BusAPIResponse
	err = (&r).UnmarshalJSON(buff)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (c *Client) get(u *url.URL) ([]byte, error) {
	resp, err := c.HTTP.Get(u.String())
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	buff, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode < 200 && resp.StatusCode >= 300 {
		return nil, &HTTPError{
			ResponseData: buff,
			StatusCode:   resp.StatusCode,
		}
	}

	return buff, err
}
