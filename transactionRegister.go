package przelewy24api

import (
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
)

//TODO: Document fields?
//CartItem is meant to hold information about purchases in TransactionRegister. They are required if user is paying via paypal,
type CartItem struct {
	SellerId       string `json:"sellerId"`
	SellerCategory string `json:"sellerCategory"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	Quantity       int    `json:"quantity"`
	Price          int    `json:"price"`
	Number         string `json:"number"`
}

//Shipping is meant to hold shipping info, which is additional info to TransactionRegister.
type Shipping struct {
	Type    int    `json:"type"`
	Address string `json:"address"`
	Zip     string `json:"zip"`
	City    string `json:"city"`
	Country string `json:"country"`
}

//Additional is nested part of TransactionRegister struct, defined by przelewy24api. Check api docs for reference https://developers.przelewy24.pl/index.php?pl#tag/Obsluga-transakcji-API/paths/~1api~1v1~1transaction~1register/post
type Additional struct {
	Shipping Shipping `json:"shipping"`
}

//TransactionRegister is meant to be sent to create new transaction in przelewy24.
type TransactionRegister struct {
	//MerchantId required
	MerchantId uint64 `json:"merchantId"`
	//PosId required
	PosId uint64 `json:"posId"`
	//SessionId required
	SessionId string `json:"sessionId"`
	//Amount required
	Amount uint64 `json:"amount"`
	//Currency required
	Currency string `json:"currency"`
	//Description required
	Description string `json:"description"`
	//Email required
	Email   string `json:"email"`
	Client  string `json:"client,omitempty"`
	Address string `json:"address,omitempty"`
	Zip     string `json:"zip,omitempty"`
	City    string `json:"city,omitempty"`
	//Country required
	Country string `json:"country"`
	Phone   string `json:"phone,omitempty"`
	//Language required
	Language string `json:"language"`
	Method   int    `json:"method,omitempty"`
	//URLReturn required
	URLReturn        string `json:"urlReturn"`
	URLStatus        string `json:"urlStatus,omitempty"`
	TimeLimit        int    `json:"timeLimit,omitempty"`
	Channel          int    `json:"channel,omitempty"`
	WaitForResult    bool   `json:"waitForResult,omitempty"`
	RegulationAccept bool   `json:"regulationAccept,omitempty"`
	Shipping         int    `json:"shipping,omitempty"`
	TransferLabel    string `json:"transferLabel,omitempty"`
	MobileLib        int    `json:"mobileLib,omitempty"`
	SDKVersion       string `json:"sdkVersion,omitempty"`
	//Sign required
	Sign        string `json:"sign"`
	Encoding    string `json:"encoding,omitempty"`
	MethodRefId string `json:"methodRefId,omitempty"`
	//Cart required if user is paying via paypal
	Cart       []CartItem  `json:"cart,omitempty"`
	Additional *Additional `json:"additional,omitempty"`
}

//sign signs transactionRegister as required by przelewy24 api. See api documentation for reference https://developers.przelewy24.pl/index.php?pl#tag/Obsluga-transakcji-API/paths/~1api~1v1~1transaction~1register/post
func (tr *TransactionRegister) sign(crc string) {
	type SigningStruct struct {
		SessionId  string `json:"sessionId"`
		MerchantId uint64 `json:"merchantId"`
		Amount     uint64 `json:"amount"`
		Currency   string `json:"currency"`
		Crc        string `json:"crc"`
	}
	signingStruct := SigningStruct{
		SessionId:  tr.SessionId,
		MerchantId: tr.MerchantId,
		Amount:     tr.Amount,
		Currency:   tr.Currency,
		Crc:        crc,
	}
	b, _ := json.Marshal(signingStruct)
	sum := sha512.Sum384(b)
	tr.Sign = hex.EncodeToString(sum[:])
}
