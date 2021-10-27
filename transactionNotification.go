package przelewy24api

type TransactionNotification struct {
	MerchantId   int    `json:"merchantId"`
	PosId        int    `json:"posId"`
	SessionId    string `json:"sessionId"`
	Amount       int    `json:"amount"`
	OriginAmount int    `json:"originAmount"`
	Currency     string `json:"currency"`
	OrderId      int    `json:"orderId"`
	MethodId     int    `json:"methodId"`
	Statement    string `json:"statement"`
	Sign         string `json:"sign"`
}
