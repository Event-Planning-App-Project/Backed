package midtrans

import (
	"bytes"
	"encoding/json"
	"event/config"
	"fmt"
	"net/http"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type SnapMidtrans struct {
	s snap.Client
}

func InitMidtrans() *SnapMidtrans {
	config := config.InitConfig()
	TokenMidtrans := config.TokenMidtrans
	s := snap.Client{}
	s.New(TokenMidtrans, midtrans.Sandbox)
	return &SnapMidtrans{
		s: s,
	}
}

func (s *SnapMidtrans) CreateTransaction(OrderID string, GrossAmt int64) map[string]interface{} {
	requestBody := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  OrderID,
			GrossAmt: GrossAmt,
		},
		// Callbacks: &snap.Callbacks{
		// 	Finish: ("http://54.179.30.163:8050/transaction/finish_payment"),
		// },
	}
	jsonReq, _ := json.Marshal(requestBody)
	buf := bytes.NewBuffer(jsonReq)
	type ResponseWithMap map[string]interface{}
	Resp := ResponseWithMap{}
	err := s.s.HttpClient.Call(http.MethodPost, "https://app.sandbox.midtrans.com/snap/v1/transactions", &s.s.ServerKey, s.s.Options, buf, &Resp)
	fmt.Println(err)
	return Resp
}
