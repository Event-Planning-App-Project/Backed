package midtrans

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"event/config"
	"event/delivery/view/transaction"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
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
		Callbacks: &snap.Callbacks{
			Finish: ("http://54.179.30.163:8050/transaction/finish_payment"),
		},
	}
	jsonReq, _ := json.Marshal(requestBody)
	buf := bytes.NewBuffer(jsonReq)
	type ResponseWithMap map[string]interface{}
	Resp := ResponseWithMap{}
	err := s.s.HttpClient.Call(http.MethodPost, "https://app.sandbox.midtrans.com/snap/v1/transactions", &s.s.ServerKey, s.s.Options, buf, &Resp)
	fmt.Println(err)
	return Resp
}

func (s *SnapMidtrans) FinishPayment(order string) transaction.ResponsePayment {
	url := fmt.Sprintf("https://api.sandbox.midtrans.com/v2/%s/status", order)
	method := "GET"

	payload := strings.NewReader("\n\n")

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err, "err")
	}

	key := s.s.ServerKey
	EncodeKey := base64.StdEncoding.EncodeToString([]byte(key))
	fmt.Println(EncodeKey)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", EncodeKey)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	var response transaction.ResponsePayment
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	json.Unmarshal(body, &response)
	return response
}

func PayOk() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Success Pay")
	}
}
