package wayforpay

import (
	"crypto/hmac"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type RefundRequest struct {
	TransactionType   string `json:"transactionType"`
	MerchantAccount   string `json:"merchantAccount"`
	OrderReference    string `json:"orderReference"`
	Amount            int    `json:"amount"`
	Currency          string `json:"currency"`
	Comment           string `json:"comment"`
	MerchantSignature string `json:"merchantSignature"`
	ApiVersion        int    `json:"apiVersion"`
}

func (w *WayForPay) CreateRefund(request *RefundRequest) (*RefundResponse, error) {

	respBody := request.body(w.merchantSecret)
	if err := request.validate(); err != nil {
		return nil, err
	}
	params, err := request.params()
	if err != nil {
		return nil, err
	}
	var cir RefundResponse
	if err := w.makeRequest(fmt.Sprintf(APIEndpoint, request.method()), respBody, &cir, params); err != nil {
		return nil, err
	}
	return &cir, nil
}

func (r *RefundRequest) validate() error {
	return nil
}

func (r *RefundRequest) params() (Params, error) {
	return Params{}, nil
}

func (r *RefundRequest) method() string {
	return ""
}

func (w *WayForPay) NewRefundRequest() *RefundRequest {
	return &RefundRequest{
		TransactionType: "REFUND",
		ApiVersion:      1,
		MerchantAccount: w.merchantLogin,
	}
}

func (r *RefundRequest) body(secret string) io.Reader {
	data := []string{
		r.MerchantAccount,
		r.OrderReference,
		strconv.FormatInt(int64(r.Amount), 10),
		r.Currency,
	}

	message := strings.Join(data, ";")
	h := hmac.New(md5.New, []byte(secret))
	h.Write([]byte(message))
	r.MerchantSignature = hex.EncodeToString(h.Sum(nil))

	body, err := json.Marshal(r)
	if err != nil {
		return nil
	}

	return strings.NewReader(string(body))
}

func (r *RefundRequest) SetMerchantAccount(merchantAccount string) *RefundRequest {
	r.MerchantAccount = merchantAccount
	return r
}

func (r *RefundRequest) SetOrderReference(orderReference string) *RefundRequest {
	r.OrderReference = orderReference
	return r
}

func (r *RefundRequest) SetAmount(amount int) *RefundRequest {
	r.Amount = amount
	return r
}

func (r *RefundRequest) SetCurrency(currency string) *RefundRequest {
	r.Currency = currency
	return r
}

func (r *RefundRequest) SetComment(comment string) *RefundRequest {
	r.Comment = comment
	return r
}

func (r *RefundRequest) SetMerchantSignature(merchantSignature string) *RefundRequest {
	r.MerchantSignature = merchantSignature
	return r
}

func (r *RefundRequest) SetApiVersion(apiVersion int) *RefundRequest {
	r.ApiVersion = apiVersion
	return r
}

type RefundResponse struct {
	OrderReference    string `json:"orderReference"`
	TransactionStatus string `json:"transactionStatus"`
	ReasonCode        int    `json:"reasonCode"`
	Reason            string `json:"reason"`
	MerchantAccount   string `json:"merchantAccount"`
}

func (c *RefundResponse) Error() error {
	if c.ReasonCode != 1100 {
		return fmt.Errorf("%d: %s", c.ReasonCode, c.Reason)
	}
	return nil
}

func (c *RefundResponse) GetReasonCode() int {
	return c.ReasonCode
}

func (c *RefundResponse) GetReason() string {
	return c.Reason
}
