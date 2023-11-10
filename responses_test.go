package wayforpay_test

import (
	"net/http"
	"testing"

	wfp "github.com/fairytale5571/wayforpay"
	"github.com/stretchr/testify/require"
)

func TestWayForPay_NewResponse(t *testing.T) {
	cases := []struct {
		name string
		on   func() *wfp.Response
		want *wfp.Response
	}{
		{
			name: "success",
			on: func() *wfp.Response {
				client := &http.Client{
					Timeout: 10,
				}
				wfpClient, err := wfp.NewClient(client, merchantLogin, merchantSecret)
				if err != nil {
					t.Fatalf("NewClient() error = %v", err)
				}
				return wfpClient.NewResponse("AAA", "accept", 123456789)
			},
			want: &wfp.Response{OrderReference: "AAA", Status: "accept", Time: 123456789, Signature: "7685833061e09025c2d7afab01404f6e"},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.on()
			require.Equal(t, got, tt.want)
		})
	}
}
