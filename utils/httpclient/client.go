package httpclient

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/JingruiLea/ad_boost/common/logs"
	"io/ioutil"
	"net/http"
)

var CommonHeader = map[string]string{
	"Content-Type": "application/json",
	"Access-Token": "b7f1815041aa3835b1e8dcc4ede24e7a33cd103e",
}

type Client struct {
	httpClient *http.Client
}

func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{},
	}
}
func (c *Client) Get(ctx context.Context, url string, headers map[string]string, response interface{}, params ...map[string]interface{}) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		logs.CtxErrorf(ctx, "Get http.NewRequest error: %v", err)
		return err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	q := req.URL.Query()

	for _, param := range params {
		for k, v := range param {
			q.Add(k, convertToString(v))
		}
	}

	req.URL.RawQuery = q.Encode()

	logs.CtxDebugf(ctx, "Request: %s %s", req.Method, req.URL.String())

	resp, err := c.httpClient.Do(req)
	if err != nil {
		logs.CtxErrorf(ctx, "Get c.httpClient.Do error: %v", err)
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.CtxErrorf(ctx, "Get ioutil.ReadAll error: %v", err)
		return err
	}

	logs.CtxDebugf(ctx, "Response code: %d", resp.StatusCode)

	return json.Unmarshal(body, &response)
}

func (c *Client) Post(ctx context.Context, url string, headers map[string]string, request, response interface{}) error {
	jsonReq, err := json.Marshal(request)
	if err != nil {
		logs.CtxErrorf(ctx, "Post json.Marshal error: %v", err)
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonReq))
	if err != nil {
		logs.CtxErrorf(ctx, "Post http.NewRequest error: %v", err)
		return err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}
	logs.CtxDebugf(ctx, "Request: %s %s", req.Method, req.URL.String())
	logs.CtxDebugf(ctx, "Request body: %s", string(jsonReq))

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.CtxErrorf(ctx, "Post ioutil.ReadAll error: %v", err)
		return err
	}

	logs.CtxDebugf(ctx, "Response code: %d", resp.StatusCode)

	return json.Unmarshal(body, &response)
}

func convertToString(v interface{}) string {
	switch value := v.(type) {
	case string:
		return value
	default:
		jsonValue, _ := json.Marshal(value)
		return string(jsonValue)
	}
}
