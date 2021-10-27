package przelewy24api

import (
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
)

//TransactionVerify is meant to be sent after receiving transactionNotification to przelewy24api. It is compulsory to confirm transaction, if transactionVerify is not sent, transaction is not confirmed and money is not available to the seller.
type TransactionVerify struct {
	MerchantId uint64 `json:"merchantId"`
	PosId      uint64 `json:"posId"`
	SessionId  string `json:"sessionId"`
	Amount     uint64 `json:"amount"`
	Currency   string `json:"currency"`
	OrderId    int    `json:"orderId"`
	Sign       string `json:"sign"`
}

//sign signs transactionVerify as required by przelewy24 api. See api documentation for reference https://developers.przelewy24.pl/index.php?pl#tag/Obsluga-transakcji-API/paths/~1api~1v1~1transaction~1register/post
func (tv *TransactionVerify) sign(crc string) {
	type SigningStruct struct {
		SessionId string `json:"sessionId"`
		OrderId   int    `json:"orderId"`
		Amount    uint64 `json:"amount"`
		Currency  string `json:"currency"`
		Crc       string `json:"crc"`
	}
	signingStruct := SigningStruct{
		SessionId: tv.SessionId,
		OrderId:   tv.OrderId,
		Amount:    tv.Amount,
		Currency:  tv.Currency,
		Crc:       crc,
	}
	b, _ := json.Marshal(signingStruct)
	sum := sha512.Sum384(b)
	tv.Sign = hex.EncodeToString(sum[:])
}
