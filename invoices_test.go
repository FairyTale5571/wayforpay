package wayforpay_test

import (
	wfp "github.com/fairytale5571/wayforpay"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
	"time"
)

const (
	merchantLogin  = "test_merch_n1"
	merchantSecret = "flk3409refn54t54t*FNJRET"
)

func TestWayForPay_NewCreateInvoiceRequest(t *testing.T) {
	cases := []struct {
		name string
		want *wfp.CreateInvoiceRequest
	}{
		{
			name: "success",
			want: &wfp.CreateInvoiceRequest{
				TransactionType: "CREATE_INVOICE", MerchantAccount: merchantLogin, MerchantTransactionType: "", MerchantAuthType: wfp.SignatureModeSimple, MerchantDomainName: "", MerchantSignature: "", ApiVersion: "1", Language: "EN", NotifyMethod: "all", ServiceUrl: "", OrderReference: "", OrderDate: 0, Amount: "", Currency: "", AlternativeAmount: "", AlternativeCurrency: "", OrderTimeout: "", HoldTimeout: "", ProductName: []string(nil), ProductPrice: []string(nil), ProductCount: []string(nil), PaymentSystems: "", ClientFirstName: "", ClientLastName: "", ClientEmail: "", ClientPhone: "",
			},
		},
	}
	client := &http.Client{
		Timeout: 10,
	}
	wfpClient, err := wfp.NewClient(client, merchantLogin, merchantSecret)
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := wfpClient.NewCreateInvoiceRequest()
			require.NotNil(t, got)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestWayForPay_SendInvoice(t *testing.T) {

	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	wfpClient, err := wfp.NewClient(client, merchantLogin, merchantSecret)
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}
	cases := []struct {
		name        string
		request     *wfp.CreateInvoiceRequest
		want        *wfp.APIResponse
		expectedErr bool
	}{
		{
			name: "success",
			request: wfpClient.NewCreateInvoiceRequest().
				SetMerchantDomainName("test.com").
				SetOrderDate(time.Now()).
				SetAmount("100").
				SetCurrency("UAH").
				SetOrderReference(uuid.New().String()).
				AddProduct("test", "100", "1"),
		},
		{
			name: "error currency",
			request: wfpClient.NewCreateInvoiceRequest().
				SetMerchantDomainName("test.com").
				SetOrderDate(time.Now()).
				SetAmount("100").
				SetCurrency("UAH_NOT_RUB").
				SetOrderReference(uuid.New().String()).
				AddProduct("test", "100", "1"),
			expectedErr: true,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := wfpClient.CreateInvoice(tt.request)
			if tt.expectedErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			require.NotNil(t, resp)
		})
	}
}
