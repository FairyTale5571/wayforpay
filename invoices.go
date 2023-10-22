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
	"time"
)

type CreateInvoiceRequest struct {
	TransactionType         string        `json:"transactionType"`
	MerchantAccount         string        `json:"merchantAccount"`
	MerchantTransactionType string        `json:"merchantTransactionType,omitempty"`
	MerchantAuthType        SignatureMode `json:"merchantAuthType,omitempty"`
	MerchantDomainName      string        `json:"merchantDomainName"`
	MerchantSignature       string        `json:"merchantSignature"`
	ApiVersion              string        `json:"apiVersion"`
	Language                string        `json:"language,omitempty"`
	NotifyMethod            string        `json:"notifyMethod,omitempty"`
	ServiceUrl              string        `json:"serviceUrl,omitempty"`
	OrderReference          string        `json:"orderReference"`
	OrderDate               int64         `json:"orderDate"`
	Amount                  string        `json:"amount"`
	Currency                string        `json:"currency"`
	AlternativeAmount       string        `json:"alternativeAmount,omitempty"`
	AlternativeCurrency     string        `json:"alternativeCurrency,omitempty"`
	OrderTimeout            string        `json:"orderTimeout,omitempty"`
	HoldTimeout             string        `json:"holdTimeout,omitempty"`
	ProductName             []string      `json:"productName"`
	ProductPrice            []string      `json:"productPrice"`
	ProductCount            []string      `json:"productCount"`
	PaymentSystems          string        `json:"paymentSystems,omitempty"`
	ClientFirstName         string        `json:"clientFirstName,omitempty"`
	ClientLastName          string        `json:"clientLastName,omitempty"`
	ClientEmail             string        `json:"clientEmail,omitempty"`
	ClientPhone             string        `json:"clientPhone,omitempty"`
}

// NewCreateInvoiceRequest returns a new CreateInvoiceRequest.
func (w *WayForPay) NewCreateInvoiceRequest() *CreateInvoiceRequest {
	return &CreateInvoiceRequest{
		TransactionType:  "CREATE_INVOICE",
		ApiVersion:       "1",
		Language:         "EN",
		NotifyMethod:     "all",
		MerchantAccount:  w.merchantLogin,
		MerchantAuthType: SignatureModeSimple,
	}
}

func (c *CreateInvoiceRequest) params() (Params, error) {
	return Params{}, nil
}

func (c *CreateInvoiceRequest) method() string {
	return "/pay"
}

func (c *CreateInvoiceRequest) body(secret string) io.Reader {
	data := []string{
		c.MerchantAccount,
		c.MerchantDomainName,
		c.OrderReference,
		strconv.FormatInt(c.OrderDate, 10),
		c.Amount,
		c.Currency,
	}

	data = append(data, c.ProductName...)
	data = append(data, c.ProductCount...)
	data = append(data, c.ProductPrice...)

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

// SetMerchantAccount sets the merchant account.
func (c *CreateInvoiceRequest) SetMerchantAccount(merchantAccount string) *CreateInvoiceRequest {
	c.MerchantAccount = merchantAccount
	return c
}

// SetMerchantTransactionType sets the merchant transaction type.
func (c *CreateInvoiceRequest) SetMerchantTransactionType(merchantTransactionType string) *CreateInvoiceRequest {
	c.MerchantTransactionType = merchantTransactionType
	return c
}

// SetMerchantAuthType sets the merchant auth type.
func (c *CreateInvoiceRequest) SetMerchantAuthType(merchantAuthType SignatureMode) *CreateInvoiceRequest {
	c.MerchantAuthType = merchantAuthType
	return c
}

// SetMerchantDomainName sets the merchant domain name.
func (c *CreateInvoiceRequest) SetMerchantDomainName(merchantDomainName string) *CreateInvoiceRequest {
	c.MerchantDomainName = merchantDomainName
	return c
}

// SetMerchantSignature sets the merchant signature.
func (c *CreateInvoiceRequest) SetMerchantSignature(merchantSignature string) *CreateInvoiceRequest {
	c.MerchantSignature = merchantSignature
	return c
}

// SetApiVersion sets the api version. Default: 1
func (c *CreateInvoiceRequest) SetApiVersion(apiVersion string) *CreateInvoiceRequest {
	c.ApiVersion = apiVersion
	return c
}

// SetLanguage sets the language. Default: EN
// Possible values: RU, UA, EN
func (c *CreateInvoiceRequest) SetLanguage(language string) *CreateInvoiceRequest {
	c.Language = language
	return c
}

// SetNotifyMethod sets the notify method.
// Possible values: - sms, email, bot, all
func (c *CreateInvoiceRequest) SetNotifyMethod(notifyMethod string) *CreateInvoiceRequest {
	c.NotifyMethod = notifyMethod
	return c
}

// SetServiceUrl sets the service url.
func (c *CreateInvoiceRequest) SetServiceUrl(serviceUrl string) *CreateInvoiceRequest {
	c.ServiceUrl = serviceUrl
	return c
}

func (c *CreateInvoiceRequest) SetOrderReference(orderReference string) *CreateInvoiceRequest {
	c.OrderReference = orderReference
	return c
}

func (c *CreateInvoiceRequest) SetOrderDate(orderDate time.Time) *CreateInvoiceRequest {
	c.OrderDate = orderDate.Unix()
	return c
}

func (c *CreateInvoiceRequest) SetAmount(amount string) *CreateInvoiceRequest {
	c.Amount = amount
	return c
}

func (c *CreateInvoiceRequest) SetCurrency(currency string) *CreateInvoiceRequest {
	c.Currency = currency
	return c
}

func (c *CreateInvoiceRequest) SetAlternativeAmount(alternativeAmount string) *CreateInvoiceRequest {
	c.AlternativeAmount = alternativeAmount
	return c
}

func (c *CreateInvoiceRequest) SetAlternativeCurrency(alternativeCurrency string) *CreateInvoiceRequest {
	c.AlternativeCurrency = alternativeCurrency
	return c
}

func (c *CreateInvoiceRequest) SetOrderTimeout(orderTimeout string) *CreateInvoiceRequest {
	c.OrderTimeout = orderTimeout
	return c
}

func (c *CreateInvoiceRequest) SetHoldTimeout(holdTimeout string) *CreateInvoiceRequest {
	c.HoldTimeout = holdTimeout
	return c
}

func (c *CreateInvoiceRequest) AddProduct(productName, productPrice, productCount string) *CreateInvoiceRequest {
	c.ProductName = append(c.ProductName, productName)
	c.ProductPrice = append(c.ProductPrice, productPrice)
	c.ProductCount = append(c.ProductCount, productCount)
	return c
}

func (c *CreateInvoiceRequest) SetPaymentSystems(paymentSystems ...string) *CreateInvoiceRequest {
	// split payment systems by semicolon
	c.PaymentSystems = strings.Join(paymentSystems, ";")
	return c
}

func (c *CreateInvoiceRequest) SetClientFirstName(clientFirstName string) *CreateInvoiceRequest {
	c.ClientFirstName = clientFirstName
	return c
}

func (c *CreateInvoiceRequest) SetClientLastName(clientLastName string) *CreateInvoiceRequest {
	c.ClientLastName = clientLastName
	return c
}

func (c *CreateInvoiceRequest) SetClientEmail(clientEmail string) *CreateInvoiceRequest {
	c.ClientEmail = clientEmail
	return c
}

func (c *CreateInvoiceRequest) SetClientPhone(clientPhone string) *CreateInvoiceRequest {
	c.ClientPhone = clientPhone
	return c
}

func (c *CreateInvoiceRequest) validate() error {
	if c.TransactionType == "" {
		return ErrTransactionTypeRequired
	}
	if c.MerchantAccount == "" {
		return ErrMerchantAccountRequired
	}
	if c.MerchantDomainName == "" {
		return ErrMerchantDomainNameRequired
	}
	if c.MerchantSignature == "" {
		return ErrMerchantSignatureRequired
	}
	if c.ApiVersion == "" {
		return ErrApiVersionRequired
	}
	if c.OrderReference == "" {
		return ErrOrderReferenceRequired
	}
	if c.OrderDate == 0 {
		return ErrOrderDateRequired
	}
	if c.Amount == "" {
		return ErrAmountRequired
	}
	if c.Currency == "" {
		return ErrCurrencyRequired
	}
	if len(c.ProductName) == 0 {
		return ErrProductNameRequired
	}
	if len(c.ProductPrice) == 0 {
		return ErrProductPriceRequired
	}
	if len(c.ProductCount) == 0 {
		return ErrProductCountRequired
	}
	return nil
}

type CreateInvoiceResponse struct {
	Reason     string `json:"reason"`
	ReasonCode int    `json:"reasonCode"`
	InvoiceURL string `json:"invoiceUrl"`
	QRCode     string `json:"qrCode"`
}

func (c *CreateInvoiceResponse) Error() error {
	if c.ReasonCode != 1100 {
		return fmt.Errorf("%d: %s", c.ReasonCode, c.Reason)
	}
	return nil
}

func (c *CreateInvoiceResponse) GetReasonCode() int {
	return c.ReasonCode
}

func (c *CreateInvoiceResponse) GetReason() string {
	return c.Reason
}

func (w *WayForPay) CreateInvoice(request *CreateInvoiceRequest) (*CreateInvoiceResponse, error) {

	respBody := request.body(w.merchantSecret)
	if err := request.validate(); err != nil {
		return nil, err
	}
	params, err := request.params()
	if err != nil {
		return nil, err
	}
	var cir CreateInvoiceResponse
	if err := w.makeRequest(fmt.Sprintf(APIEndpoint, request.method()), respBody, &cir, params); err != nil {
		return nil, err
	}
	return &cir, nil
}

type RemoveInvoiceRequest struct {
	TransactionType   string `json:"transactionType"`
	ApiVersion        string `json:"apiVersion"`
	MerchantAccount   string `json:"merchantAccount"`
	OrderReference    string `json:"orderReference"`
	MerchantSignature string `json:"merchantSignature"`
}

// NewRemoveInvoiceRequest returns a new RemoveInvoiceRequest.
func (w *WayForPay) NewRemoveInvoiceRequest() *RemoveInvoiceRequest {
	return &RemoveInvoiceRequest{
		TransactionType: "REMOVE_INVOICE",
		ApiVersion:      "1",
	}
}

// SetMerchantAccount sets the transaction type.
func (r *RemoveInvoiceRequest) SetMerchantAccount(merchantAccount string) *RemoveInvoiceRequest {
	r.MerchantAccount = merchantAccount
	return r
}

// SetMerchantSignature sets the merchant signature
// OPTIONAL: Only if you want set your own signature
func (r *RemoveInvoiceRequest) SetMerchantSignature(merchantSignature string) *RemoveInvoiceRequest {
	r.MerchantSignature = merchantSignature
	return r
}

// SetOrderReference sets the order reference.
func (r *RemoveInvoiceRequest) SetOrderReference(orderReference string) *RemoveInvoiceRequest {
	r.OrderReference = orderReference
	return r
}

func (r *RemoveInvoiceRequest) params() (Params, error) {
	return Params{}, nil
}

func (r *RemoveInvoiceRequest) method() string {
	return "/pay"
}

func (r *RemoveInvoiceRequest) body(secret string) io.Reader {
	data := []string{
		r.MerchantAccount,
		r.OrderReference,
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

func (r *RemoveInvoiceRequest) validate() error {
	if r.TransactionType == "" {
		return ErrTransactionTypeRequired
	}
	if r.MerchantAccount == "" {
		return ErrMerchantAccountRequired
	}
	if r.OrderReference == "" {
		return ErrOrderReferenceRequired
	}
	return nil
}

func (w *WayForPay) RemoveInvoice(request *RemoveInvoiceRequest) (*RemoveInvoiceResponse, error) {

	respBody := request.body(w.merchantSecret)
	if err := request.validate(); err != nil {
		return nil, err
	}
	params, err := request.params()
	if err != nil {
		return nil, err
	}
	var rir RemoveInvoiceResponse
	if err := w.makeRequest(fmt.Sprintf(APIEndpoint, request.method()), respBody, &rir, params); err != nil {
		return nil, err
	}
	return &rir, nil
}

type RemoveInvoiceResponse struct {
	Reason     string `json:"reason"`
	ReasonCode int    `json:"reasonCode"`
}

func (r *RemoveInvoiceResponse) Error() error {
	if r.ReasonCode != 1100 {
		return fmt.Errorf("%d: %s", r.ReasonCode, r.Reason)
	}
	return nil
}

func (r *RemoveInvoiceResponse) GetReasonCode() int {
	return r.ReasonCode
}

func (r *RemoveInvoiceResponse) GetReason() string {
	return r.Reason
}
