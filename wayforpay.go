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

func (w *WayForPay) makeRequest(endpoint string, body io.Reader, response Responder, params Params) error {
	method := fmt.Sprintf(APIEndpoint, endpoint)
	rawUrl, err := url.Parse(method)
	if err != nil {
		return err
	}
	rawUrl.RawQuery = buildParams(params).Encode()
	method = rawUrl.String()

	req, err := http.NewRequest(http.MethodPost, method, body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	res, err := w.client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	respBody, err := io.ReadAll(res.Body)

	if err := json.Unmarshal(respBody, &response); err != nil {
		return err
	}
	if response.GetReasonCode() != 1100 {
		return fmt.Errorf("api error: code: %v, reason: %v", response.GetReasonCode(), response.GetReason())
	}
	return nil
}
