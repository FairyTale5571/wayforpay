package wayforpay

import (
	"crypto/hmac"
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"strings"
)

type Response struct {
	OrderReference string `json:"orderReference"`
	Status         string `json:"status"`
	Time           int64  `json:"time"`
	Signature      string `json:"signature"`
}

func (r *Response) sign(secret string) {
	data := []string{
		r.OrderReference,
		r.Status,
		strconv.FormatInt(r.Time, 10),
	}
	message := strings.Join(data, ";")
	h := hmac.New(md5.New, []byte(secret))
	h.Write([]byte(message))
	r.Signature = hex.EncodeToString(h.Sum(nil))
}

func (w *WayForPay) NewResponse(orderReference, status string, time int64) *Response {
	resp := &Response{
		OrderReference: orderReference,
		Status:         status,
		Time:           time,
	}
	resp.sign(w.merchantSecret)
	return resp
}
