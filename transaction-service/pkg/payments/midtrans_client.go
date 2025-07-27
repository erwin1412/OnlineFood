package payments

import (
	"fmt"

	"github.com/veritrans/go-midtrans"
)

type MidtransClient struct {
	Client *midtrans.Client
}

func NewMidtransClient(serverKey string, isProduction bool) *MidtransClient {
	client := midtrans.NewClient()
	client.ServerKey = serverKey
	if isProduction {
		client.APIEnvType = midtrans.Production
	} else {
		client.APIEnvType = midtrans.Sandbox
	}

	return &MidtransClient{
		Client: &client,
	}
}

func (m *MidtransClient) CreateSnapToken(orderID string, amount int64, customerName, customerEmail string) (string, error) {
	// 🚩 PAKAI POINTER
	snapGateway := midtrans.SnapGateway{
		Client: *m.Client,
	}

	req := &midtrans.SnapReq{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  orderID,
			GrossAmt: amount,
		},
		CustomerDetail: &midtrans.CustDetail{
			FName: customerName,
			Email: customerEmail,
		},
	}

	snapResp, err := snapGateway.GetToken(req)
	if err != nil {
		return "", err
	}

	// redirect url print
	// Construct redirect URL based on environment
	var redirectURL string
	if m.Client.APIEnvType == midtrans.Production {
		redirectURL = fmt.Sprintf("https://app.midtrans.com/snap/v2/vtweb/%s", snapResp.Token)
	} else {
		redirectURL = fmt.Sprintf("https://app.sandbox.midtrans.com/snap/v2/vtweb/%s", snapResp.Token)
	}
	fmt.Println("Redirect URL:", redirectURL)

	return snapResp.Token, nil
}
