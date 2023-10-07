package wayforpay

import "errors"

var (
	ErrMerchantLoginRequired      = errors.New("merchant login is required")
	ErrMerchantSecretRequired     = errors.New("merchant secret is required")
	ErrSecretCodeRequired         = errors.New("secret code is required")
	ErrTransactionTypeRequired    = errors.New("transactionType is required")
	ErrMerchantAccountRequired    = errors.New("merchantAccount is required")
	ErrMerchantDomainNameRequired = errors.New("merchantDomainName is required")
	ErrMerchantSignatureRequired  = errors.New("merchantSignature is required")
	ErrApiVersionRequired         = errors.New("apiVersion is required")
	ErrOrderReferenceRequired     = errors.New("orderReference is required")
	ErrOrderDateRequired          = errors.New("orderDate is required")
	ErrAmountRequired             = errors.New("amount is required")
	ErrCurrencyRequired           = errors.New("currency is required")
	ErrProductNameRequired        = errors.New("productName is required")
	ErrProductPriceRequired       = errors.New("productPrice is required")
	ErrProductCountRequired       = errors.New("productCount is required")
)
