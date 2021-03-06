package main

import (
	"fmt"
	"mpos/mod"

	"github.com/AYJiaYou/restest"
)

func main() {
	client := restest.NewClient()
	client.SetCheckServerCert(false)

	ca := restest.NewCase()
	ca.SetHost("https://localhost:8083")
	ca.SetMethod("POST")
	ca.SetPath("orders")
	ca.SetToken(
		"Authorization",
		"'TSign ' + $SerialNumber + ':' + urlsafe_base64(hmac_sha1($ReqPath + '\n' + $ReqBody, $SecretKey))",
	)
	ca.SetVariable("SerialNumber", "53a6d334c3b84e45937f2c1e1d6364f2")
	ca.SetVariable("SecretKey", "7d6387d1c355467885d47f7d92cd0af5")

	order := mod.OrderInfo{
		ClientID:   1,
		CustomerID: 1,
		CardNo:     123,
		BankID:     1,
	}
	ca.SetBodyJSON(&order)

	res, err := client.RunCase(ca)
	if err != nil {
		fmt.Println("RunCase Error:", err)
	}
	str, err := res.GetString()
	fmt.Println(err, str)
}
