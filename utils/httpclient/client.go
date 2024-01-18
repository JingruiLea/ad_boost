package httpclient

import (
	"bytes"
	"context"
	"fmt"
	"github.com/JingruiLea/ad_boost/common/logs"
	"github.com/JingruiLea/ad_boost/logic/auth"
	"github.com/JingruiLea/ad_boost/model/ttypes"
	jsoniter "github.com/json-iterator/go"
	"io/ioutil"
	"net/http"
	"reflect"
)

type Client struct {
	httpClient *http.Client
}

func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{},
	}
}

type AdResp struct {
	ttypes.BaseResp
	Data jsoniter.RawMessage `json:"data"`
}

func (c *Client) AdRequest(ctx context.Context, accountID int64, url string, request, response interface{}, method string, onRedo bool, params ...map[string]interface{}) error {
	accessToken, err := auth.GetAccessToken(ctx, accountID)
	if err != nil {
		logs.CtxErrorf(ctx, "AdRequest auth.GetAccessToken error: %v", err)
		return err
	}

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Access-Token"] = accessToken

	var adResp AdResp

	if method == "POST" {
		err = c.Post(ctx, url, headers, request, &adResp)
	} else if method == "GET" {
		err = c.Get(ctx, url, headers, &adResp, params...)
	}

	if err != nil {
		return err
	}

	if adResp.Code > 40100 && adResp.Code < 40200 { //Access-Token过期
		if onRedo {
			return fmt.Errorf("AdRequest Access-Token过期, 重试失败, RequestID:%s", adResp.RequestID)
		}
		logs.CtxInfof(ctx, "AdRequest Access-Token过期, 尝试刷新.")
		auth.RefreshTokenByAccountID(ctx, accountID)
		return c.AdRequest(ctx, accountID, url, request, response, method, true, params...)
	}

	if adResp.Code != 0 {
		return fmt.Errorf("AdRequest adResp.Code %d, Message: %s, RequestID:%s", adResp.Code, adResp.Message, adResp.RequestID)
	}

	decoder := jsoniter.NewDecoder(bytes.NewReader(adResp.Data))
	decoder.UseNumber()
	return decoder.Decode(&response)
}

func (c *Client) AdPost(ctx context.Context, accountID int64, url string, request, response interface{}) error {
	return c.AdRequest(ctx, accountID, url, request, response, "POST", false)
}

func (c *Client) AdGet(ctx context.Context, accountID int64, url string, response interface{}, params ...map[string]interface{}) error {
	return c.AdRequest(ctx, accountID, url, nil, response, "GET", false, params...)
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

	logs.CtxDebugf(ctx, "Response code: %d, body: %s", resp.StatusCode, string(body))
	decoder := jsoniter.NewDecoder(bytes.NewReader(body))
	decoder.UseNumber()
	return decoder.Decode(&response)
}

func (c *Client) Post(ctx context.Context, url string, headers map[string]string, request, response interface{}) error {
	jsonReq, err := jsoniter.Marshal(request)
	if err != nil {
		logs.CtxErrorf(ctx, "Post decoder.Marshal error: %v", err)
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

	decoder := jsoniter.NewDecoder(bytes.NewReader(body))
	decoder.UseNumber()
	return decoder.Decode(&response)
}

func convertToString(v interface{}) string {
	vt := reflect.TypeOf(v)
	switch vt.Kind() {
	case reflect.String:
		return reflect.ValueOf(v).String()
	default:
		jsonValue, _ := jsoniter.Marshal(v)
		return string(jsonValue)
	}
}
