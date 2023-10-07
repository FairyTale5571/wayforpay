package wayforpay

import (
	"fmt"
	"io"
)

type Params map[string]string

type Payment interface {
	params() (Params, error)
	method() string
	body(secret string) io.Reader
}

type APIResponse struct {
	ReasonCode int    `json:"reasonCode,omitempty"`
	Reason     string `json:"reason,omitempty"`
}

func (r *APIResponse) Error() error {
	if r.ReasonCode != 1100 {
		return fmt.Errorf("%d: %s", r.ReasonCode, r.Reason)
	}
	return nil
}

type Responder interface {
	Error() error
}

type SignatureMode string

const (
	SignatureModeSimple        SignatureMode = "SimpleSignature"
	SignatureModeTicket        SignatureMode = "Ticket"
	SignatureModePassword      SignatureMode = "Password"
	SignatureModeCheckString   SignatureMode = "CheckString"
	SignatureModeCheckStringPb SignatureMode = "CheckStringPb"
	SignatureModeLiqPay3       SignatureMode = "LiqPay3Siganture"
	SignatureModeEcwidEcheck   SignatureMode = "EcwidEcheckSiganture"
)
