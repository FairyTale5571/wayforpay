package wayforpay

func IsSuccessHttpCode(code int) bool {
	return code >= 200 && code < 300
}
