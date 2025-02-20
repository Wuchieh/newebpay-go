/* 4.2 MPG 交易[NPA-F01] */

package ebpay

type TradeSha struct {
}

type MPGRequest struct {
	MerchantID string

	// TradeInfo
	//	交易資料
	//	AES 加密
	TradeInfo string

	// TradeSha
	//	交易資料
	//	SHA256 加密
	TradeSha    string
	Version     string
	EncryptType int
}
