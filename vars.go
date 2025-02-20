package ebpay

const (
	DateOnly = "20060102"
	TimeOnly = "150405"
	DateTime = "20060102150405"
	Version  = "2.2"
)

// 銀行代碼
const (
	BankEsun      = "Esun"      // 玉山銀行
	BankTaishin   = "Taishin"   // 台新銀行
	BankCTBC      = "CTBC"      // 中國信託銀行
	BankNCCC      = "NCCC"      // 聯合信用卡中心
	BankCathayBK  = "CathayBK"  // 國泰世華銀行
	BankCitibank  = "Citibank"  // 花旗銀行
	BankUBOT      = "UBOT"      // 聯邦銀行
	BankSKBank    = "SKBank"    // 新光銀行
	BankFubon     = "Fubon"     // 富邦銀行
	BankFirstBank = "FirstBank" // 第一銀行
	BankLINEBank  = "LINEBank"  // 連線商業銀行
	BankSinoPac   = "SinoPac"   // 永豐銀行
	BankBOT       = "BOT"       // 台灣銀行
	BankHNCB      = "HNCB"      // 華南銀行
)

// 交易類別
const (
	PaymentMethodCREDIT     = "CREDIT"     // 台灣發卡機構核發之信用卡
	PaymentMethodFOREIGN    = "FOREIGN"    // 國外發卡機構核發之卡
	PaymentMethodUNIONPAY   = "UNIONPAY"   // 銀聯卡
	PaymentMethodAPPLEPAY   = "APPLEPAY"   // ApplePay
	PaymentMethodGOOGLEPAY  = "GOOGLEPAY"  // GooglePay
	PaymentMethodSAMSUNGPAY = "SAMSUNGPAY" // SamsungPay
	PaymentMethodDCC        = "DCC"        // 動態貨幣轉換
)

// 繳費門市類別
const (
	StoreTypeSEVEN  = 1 // 統一超商
	StoreTypeFAMILY = 2 // 全家便利商店
	StoreTypeOK     = 3 // OK便利商店
	StoreTypeHILIFE = 4 // 萊爾富便利商店
)

// 繳費超商
const (
	PayStoreSEVEN  = "SEVEN"  // 統一超商
	PayStoreFAMILY = "FAMILY" // 全家便利商店
	PayStoreOK     = "OK"     // OK便利商店
	PayStoreHILIFE = "HILIFE" // 萊爾富便利商店
)

// 跨境通路類型
const (
	ChannelIDALIPAY    = "ALIPAY"    // 支付寶
	ChannelIDWECHATPAY = "WECHATPAY" // 微信支付
	ChannelIDACCLINK   = "ACCLINK"   // 約定連結帳戶
	ChannelIDCREDIT    = "CREDIT"    // 信用卡
	ChannelIDCVS       = "CVS"       // 超商代碼
	ChannelIDP2GEACC   = "P2GEACC"   // 簡單付電子帳戶轉帳
	ChannelIDVACC      = "VACC"      // ATM轉帳
	ChannelIDWEBATM    = "WEBATM"    // WebATM轉帳
)

// 加密貨幣代號
const (
	CryptoCurrencyBTC  = "BTC"  // 比特幣 BTC
	CryptoCurrencyETH  = "ETH"  // 以太幣 ETH
	CryptoCurrencyUSDT = "USDT" // 泰達幣 USDT
)

// 付款方式
const (
	PaymentTypeESUNWALLET = "ESUNWALLET" // 玉山 Wallet
	PaymentTypeLINEPAY    = "LINEPAY"    // Line Pay
	PaymentTypeTAIWANPAY  = "TAIWANPAY"  // 台灣 Pay
	PaymentTypeEZPAY      = "EZPAY"      // ezPay 電子錢包
	PaymentTypeEZPALIPAY  = "EZPALIPAY"  // Alipay 支付寶
	PaymentTypeEZPWECHAT  = "EZPWECHAT"  // WeChat 微信
)
