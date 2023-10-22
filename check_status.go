package wayforpay

import (
	"crypto/hmac"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"io"
	"strings"
)

type CheckStatus struct {
	TransactionType   string `json:"transactionType"`
	MerchantAccount   string `json:"merchantAccount"`
	OrderReference    string `json:"orderReference"`
	MerchantSignature string `json:"merchantSignature"`
	APIVersion        string `json:"apiVersion"`
}

func (w *WayForPay) NewCheckStatus(orderReference string) *CheckStatus {
	return &CheckStatus{
		TransactionType: "CHECK_STATUS",
		MerchantAccount: w.merchantLogin,
		OrderReference:  orderReference,
		APIVersion:      "1",
	}
}

func (c *CheckStatus) Validate() error {
	if c.TransactionType == "" {
		return ErrTransactionTypeRequired
	}
	if c.MerchantAccount == "" {
		return ErrMerchantAccountRequired
	}
	if c.OrderReference == "" {
		return ErrOrderReferenceRequired
	}
	if c.MerchantSignature == "" {
		return ErrMerchantSignatureRequired
	}
	if c.APIVersion == "" {
		return ErrApiVersionRequired
	}
	return nil
}

func (c *CheckStatus) body(secret string) io.Reader {
	data := []string{
		c.MerchantAccount,
		c.OrderReference,
	}

	message := strings.Join(data, ";")
	h := hmac.New(md5.New, []byte(secret))
	h.Write([]byte(message))
	c.MerchantSignature = hex.EncodeToString(h.Sum(nil))

	body, err := json.Marshal(c)
	if err != nil {
		return nil
	}

	return strings.NewReader(string(body))
}

type CheckStatusResponse struct {
	MerchantAccount   string  `json:"merchantAccount"`
	OrderReference    string  `json:"orderReference"`
	MerchantSignature string  `json:"merchantSignature"`
	Amount            string  `json:"amount"`
	Currency          string  `json:"currency"`
	AuthCode          string  `json:"authCode"`
	CreatedDate       int     `json:"createdDate"`
	ProcessingDate    int     `json:"processingDate"`
	CardPan           string  `json:"cardPan"`
	CardType          string  `json:"cardType"`
	IssuerBankCountry string  `json:"issuerBankCountry"`
	IssuerBankName    string  `json:"issuerBankName"`
	TransactionStatus string  `json:"transactionStatus"`
	Reason            string  `json:"reason"`
	ReasonCode        string  `json:"reasonCode"`
	SettlementDate    string  `json:"settlementDate"`
	SettlementAmount  float64 `json:"settlementAmount"`
	Fee               float64 `json:"fee"`
}
