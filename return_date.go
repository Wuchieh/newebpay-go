package ebpay

// CreditResult 信用卡支付回傳參數
//
//	一次付清、分期、紅利、DCC、Apple Pay、Google Pay、Samaung Pay、國民旅遊卡、銀聯、AE
type CreditResult struct {
	// AuthBank 收單金融機構
	AuthBank string `json:"AuthBank"`

	// CardBank 發卡金融機構
	CardBank string `json:"CardBank"`

	// RespondCode 金融機構回應碼
	RespondCode string `json:"RespondCode"`

	// Auth 授權碼
	Auth string `json:"Auth"`

	// Card6No 卡號前六碼
	Card6No int `json:"Card6No"`

	// Card4No 卡號末四碼
	Card4No int `json:"Card4No"`

	// Inst 分期-期別 信用卡分期交易期別。
	Inst int `json:"Inst"`

	// InstFirst 分期-首期金額 信用卡分期交易首期金額。
	InstFirst int `json:"InstFirst"`

	// InstEach 分期-每期金額 信用卡分期交易每期金額。
	InstEach int `json:"InstEach"`

	// ECI
	ECI string `json:"ECI"`

	// TokenUseStatus 信用卡快速結帳使用狀態
	TokenUseStatus int

	// RedAmt 紅利折抵後實際金額
	RedAmt int `json:"RedAmt"`

	// PaymentMethod 交易類別
	PaymentMethod string `json:"PaymentMethod"`

	// DCCAmt 外幣金額
	DCCAmt float64 `json:"DCC_Amt"`

	// DCCRate 匯率
	DCCRate float64 `json:"DCC_Rate"`

	// DCCMarkup 風險匯率
	DCCMarkup float64 `json:"DCC_Markup"`

	// DCCCurrency 幣別
	DCCCurrency string `json:"DCC_Currency"`

	// DCCCurrencyCode 幣別代碼
	DCCCurrencyCode int `json:"DCC_Currency_Code"`
}

// ATMResult WEBATM、ATM繳費回傳參數
type ATMResult struct {
	// PayBankCode 付款人金融機構代碼
	PayBankCode string `json:"PayBankCode"`

	// PayerAccount5Code 付款人金融機構帳號末五碼
	PayerAccount5Code string `json:"PayerAccount5Code"`
}

// ConvenienceStoreResult 超商代碼繳費回傳參數
type ConvenienceStoreResult struct {
	// CodeNo 繳費代碼 String(30) 繳費代碼。
	CodeNo string `json:"CodeNo"`

	// StoreType 繳費門市類別 Int(1)
	StoreType int `json:"StoreType"`

	// StoreID 繳費門市代號 String(10) 繳費門市代號 (全家回傳門市中文名稱)
	StoreID string `json:"StoreID"`
}

// ConvenienceStoreBarcodeResult 超商條碼繳費回傳參數
type ConvenienceStoreBarcodeResult struct {
	// Barcode1 第一段條碼 繳費條碼第一段條碼資料。
	Barcode1 string `json:"Barcode_1"`

	// Barcode2 第二段條碼 繳費條碼第二段條碼資料。
	Barcode2 string `json:"Barcode_2"`

	// Barcode3 第三段條碼 繳費條碼第三段條碼資料。
	Barcode3 string `json:"Barcode_3"`

	// RepayTimes 付款次數 條碼繳費付款次數
	RepayTimes int `json:"RepayTimes"`

	// PayStore 繳費超商 付款人至超商繳費，該收款超商的代碼，
	PayStore string `json:"PayStore"`
}

// CrossBorder 超商條碼繳費回傳參數
type CrossBorder struct {
	// ChannelID 跨境通路類型
	ChannelID string `json:"ChannelID"`

	// ChannelNo 跨境通路 交易序號
	ChannelNo string `json:"ChannelNo"`
}

// Logistics 超商物流回傳參數(尚未啟用)
type Logistics struct {
	// StoreCode 超商門市編號
	StoreCode string
	// StoreName 超商門市名稱
	StoreName string
	// StoreType 超商類別名稱
	StoreType string
	// StoreAddr 超商門市地址
	StoreAddr string
	// TradeType 取件交易方式
	TradeType int
	// CVSCOMName 取貨人
	CVSCOMName string
	// CVSCOMPhone 取貨人手機號碼
	CVSCOMPhone string
	// LgsNo 物流寄件單號
	LgsNo string
	// LgsType 物流型態
	LgsType string
}

type Result struct {
	// MerchantID 商店代號
	MerchantID string `json:"MerchantID"`

	// Amt 交易金額
	Amt int `json:"Amt"`

	// TradeNo 藍新金流交易序號
	TradeNo string `json:"TradeNo"`

	// MerchantOrderNo 商店訂單編號 商店自訂訂單編號
	MerchantOrderNo string `json:"MerchantOrderNo"`

	// PaymentType
	PaymentType string `json:"PaymentType"`

	// RespondType
	RespondType string `json:"RespondType"`

	// PayTime 支付完成時間
	PayTime string `json:"PayTime"`

	// IP 付款人取號或交易時的IP
	IP string `json:"IP"`

	// EscrowBank 款項保管銀行
	EscrowBank string `json:"EscrowBank"`

	// Exp
	Exp string `json:"Exp"`

	/*信用卡支付回傳參數（一次付清、分期、紅利、DCC、Apple Pay、Google Pay、Samaung Pay、國民旅遊卡、銀聯、AE）*/
	// AuthBank 收單金融機構
	AuthBank string `json:"AuthBank"`

	// CardBank 發卡金融機構
	CardBank string `json:"CardBank"`

	// RespondCode 金融機構回應碼
	RespondCode string `json:"RespondCode"`

	// Auth 授權碼
	Auth string `json:"Auth"`

	// Card6No 卡號前六碼
	Card6No int `json:"Card6No"`

	// Card4No 卡號末四碼
	Card4No int `json:"Card4No"`

	// Inst 分期-期別 信用卡分期交易期別。
	Inst int `json:"Inst"`

	// InstFirst 分期-首期金額 信用卡分期交易首期金額。
	InstFirst int `json:"InstFirst"`

	// InstEach 分期-每期金額 信用卡分期交易每期金額。
	InstEach int `json:"InstEach"`

	// ECI
	ECI string `json:"ECI"`

	// TokenUseStatus 信用卡快速結帳使用狀態
	TokenUseStatus int

	// RedAmt 紅利折抵後實際金額
	RedAmt int `json:"RedAmt"`

	// PaymentMethod 交易類別
	PaymentMethod string `json:"PaymentMethod"`

	// DCCAmt 外幣金額
	DCCAmt float64 `json:"DCC_Amt"`

	// DCCRate 匯率
	DCCRate float64 `json:"DCC_Rate"`

	// DCCMarkup 風險匯率
	DCCMarkup float64 `json:"DCC_Markup"`

	// DCCCurrency 幣別
	DCCCurrency string `json:"DCC_Currency"`

	// DCCCurrencyCode 幣別代碼
	DCCCurrencyCode int `json:"DCC_Currency_Code"`

	/*WEBATM、ATM繳費回傳參數*/

	// PayBankCode 付款人金融機構代碼
	PayBankCode string `json:"PayBankCode"`

	// PayerAccount5Code 付款人金融機構帳號末五碼
	PayerAccount5Code string `json:"PayerAccount5Code"`

	/*超商代碼繳費回傳參數 */

	// CodeNo 繳費代碼 String(30) 繳費代碼。
	CodeNo string `json:"CodeNo"`

	// StoreType 繳費門市類別 Int(1)
	StoreType int `json:"StoreType"`

	// StoreID 繳費門市代號 String(10) 繳費門市代號 (全家回傳門市中文名稱)
	StoreID string `json:"StoreID"`

	/*超商條碼繳費回傳參數*/

	// Barcode1 第一段條碼 繳費條碼第一段條碼資料。
	Barcode1 string `json:"Barcode_1"`

	// Barcode2 第二段條碼 繳費條碼第二段條碼資料。
	Barcode2 string `json:"Barcode_2"`

	// Barcode3 第三段條碼 繳費條碼第三段條碼資料。
	Barcode3 string `json:"Barcode_3"`

	// RepayTimes 付款次數 條碼繳費付款次數
	RepayTimes int `json:"RepayTimes"`

	// PayStore 繳費超商 付款人至超商繳費，該收款超商的代碼，
	PayStore string `json:"PayStore"`

	/*跨境支付回傳參數(包含簡單付電子錢包、簡單付微信支付、簡單付支付寶)*/

	// ChannelID 跨境通路類型
	ChannelID string `json:"ChannelID"`

	// ChannelNo 跨境通路 交易序號
	ChannelNo string `json:"ChannelNo"`

	/*玉山Wallet回傳參數*/

	// RedDisAmt 紅利折抵金額
	RedDisAmt int `json:"RedDisAmt"`

	/*BitoPay 回傳參數*/

	// CryptoCurrency 加密貨幣代號
	CryptoCurrency string `json:"CryptoCurrency"`

	// CryptoAmount   加密貨幣數量
	CryptoAmount string `json:"CryptoAmount"`

	// CryptoRate     加密貨幣匯率
	CryptoRate string `json:"CryptoRate"`

	/*玉山Wallet/台灣Pay/BitoPay 都有*/

	// PayAmt 實際付款金額
	PayAmt int `json:"PayAmt"`
}

func (r *Result) CreditResult() CreditResult {
	return CreditResult{
		AuthBank:        r.AuthBank,
		CardBank:        r.CardBank,
		RespondCode:     r.RespondCode,
		Auth:            r.Auth,
		Card6No:         r.Card6No,
		Card4No:         r.Card4No,
		Inst:            r.Inst,
		InstFirst:       r.InstFirst,
		InstEach:        r.InstEach,
		ECI:             r.ECI,
		TokenUseStatus:  r.TokenUseStatus,
		RedAmt:          r.RedAmt,
		PaymentMethod:   r.PaymentMethod,
		DCCAmt:          r.DCCAmt,
		DCCRate:         r.DCCRate,
		DCCMarkup:       r.DCCMarkup,
		DCCCurrency:     r.DCCCurrency,
		DCCCurrencyCode: r.DCCCurrencyCode,
	}
}

func (r *Result) ATMResult() ATMResult {
	return ATMResult{
		PayBankCode:       r.PayBankCode,
		PayerAccount5Code: r.PayerAccount5Code,
	}
}

func (r *Result) ConvenienceStoreResult() ConvenienceStoreResult {
	return ConvenienceStoreResult{
		CodeNo:    r.CodeNo,
		StoreType: r.StoreType,
		StoreID:   r.StoreID,
	}
}

func (r *Result) ConvenienceStoreBarcodeResult() ConvenienceStoreBarcodeResult {
	return ConvenienceStoreBarcodeResult{
		Barcode1:   r.Barcode1,
		Barcode2:   r.Barcode2,
		Barcode3:   r.Barcode3,
		RepayTimes: r.RepayTimes,
		PayStore:   r.PayStore,
	}
}

func (r *Result) CrossBorder() CrossBorder {
	return CrossBorder{
		ChannelID: r.ChannelID,
		ChannelNo: r.ChannelNo,
	}
}

type ReturnDate struct {
	Status  string `json:"Status"`
	Message string `json:"Message"`
	Result  Result `json:"Result"`
}
