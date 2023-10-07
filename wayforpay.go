package wayforpay

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type WayForPay struct {
	client         *http.Client
	merchantLogin  string
	merchantSecret string
}

func NewClient(httpClient *http.Client, merchantLogin, merchantSecret string) (*WayForPay, error) {
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	if merchantLogin == "" {
		return nil, ErrMerchantLoginRequired
	}
	if merchantSecret == "" {
		return nil, ErrMerchantSecretRequired
	}
	return &WayForPay{
		client:         httpClient,
		merchantLogin:  merchantLogin,
		merchantSecret: merchantSecret,
	}, nil
}

func buildParams(in Params) url.Values {
	if in == nil {
		return url.Values{}
	}

	out := url.Values{}

	for key, value := range in {
		out.Set(key, value)
	}

	return out
}

func (w *WayForPay) makeRequest(endpoint string, body Payment, params Params) (*APIResponse, error) {
	method := fmt.Sprintf(APIEndpoint, endpoint)
	rawUrl, err := url.Parse(method)
	if err != nil {
		return nil, err
	}
	rawUrl.RawQuery = buildParams(params).Encode()
	method = rawUrl.String()

	req, err := http.NewRequest(http.MethodPost, method, body.body(w.merchantSecret))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	res, err := w.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	respBody, err := io.ReadAll(res.Body)

	var response APIResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, err
	}
	if response.ReasonCode != 1100 {
		return nil, fmt.Errorf("api error: code: %v, reason: %v", response.ReasonCode, response.Reason)
	}
	return &response, nil
}

func (w *WayForPay) Request(p Payment) (Responder, error) {
	params, err := p.params()
	if err != nil {
		return nil, err
	}
	return w.makeRequest(p.method(), p, params)
}
