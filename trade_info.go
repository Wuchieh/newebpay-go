package ebpay

import (
	"encoding/json"
	"fmt"
	"github.com/duke-git/lancet/v2/cryptor"
	"net/url"
	"strings"
	"time"
)

type RespondType string

const (
	RespondTypeJSON   RespondType = "JSON"
	RespondTypeString RespondType = "String"
)

type LangType string

const (
	LangTypeEN LangType = "en"
	LangTypeTW LangType = "zh-tw"
	LangTypeJP LangType = "jp"
)

type InstFlag int

const (
	InstFlagAll InstFlag = 1
	InstFlag3   InstFlag = 3
	InstFlag6   InstFlag = 6
	InstFlag8   InstFlag = 8
	InstFlag12  InstFlag = 12
	InstFlag18  InstFlag = 18
	InstFlag24  InstFlag = 24
	InstFlag30  InstFlag = 30
)

type BankType string

const (
	// BankTypeBOT 台灣銀行
	BankTypeBOT BankType = "BOT"
	// BankTypeHNCB 華南銀行
	BankTypeHNCB BankType = "HNCB"
)

type LgsType string

const (
	// LgsTypeB2C 大宗寄倉
	LgsTypeB2C LgsType = "B2C"

	// LgsTypeC2C 店到店
	LgsTypeC2C LgsType = "C2C"
)

type TradeInfo struct {
	// RespondType 回傳格式 String(6)
	RespondType RespondType `validate:"required,response_type" json:"RespondType,omitempty"`

	// TimeStamp 時間戳記 String(50)
	//	須確實帶入自 Unix 紀元到當前時間的秒數 以避免交易失敗。(容許誤差值 120 秒)
	TimeStamp time.Time `validate:"required" json:"TimeStamp"`

	// LangType 語系 String(5)
	//	英文版 = en
	//	繁體中文版 = zh-tw
	//	日文版 = jp
	LangType LangType `validate:"lang_type" json:"LangType,omitempty"`

	// MerchantOrderNo 商店訂單編號 String(30)
	MerchantOrderNo string `validate:"required" json:"MerchantOrderNo"`

	// Amt 訂單金額 Int(10)
	Amt uint `validate:"required" json:"Amt"`

	// ItemDesc 商品資訊 String(50)
	ItemDesc string `validate:"required" json:"ItemDesc"`

	// TradeLimit 交易有效時間 Int(3)
	//	秒數下限為 60 秒，小於 60 秒以 60 秒計算
	//	秒數上限為 900 秒，大於 900 秒以 900 秒計算
	TradeLimit int `json:"TradeLimit,omitempty"`

	// ExpireDate 繳費有效期限 String(10)
	//	僅適用於非即時支付 格式為 20060102(ebpay.DateOnly)
	ExpireDate time.Time `json:"ExpireDate,omitempty"`

	// ExpireTime 繳費截止時間 String(6)
	//	僅適用於超商代碼繳費 格式為 150405
	ExpireTime string `json:"ExpireTime,omitempty"`

	// ReturnURL 支付完成 返回商店網址 String(200)
	//	交易完成後，以 Form Post 方式導回商店頁面
	//	只接受 80 與 443 Port
	ReturnURL string `json:"ReturnURL,omitempty"`

	// NotifyURL 支付通知網址 String(200)
	//	webhook 網址
	//	只接受 80 與 443 Port
	NotifyURL string `json:"NotifyURL,omitempty"`

	// CustomerURL 商店取號網址 String(200)
	CustomerURL string `json:"CustomerURL,omitempty"`

	// ClientBackURL 返回商店網址 String(200)
	//	此參數若為空值時，則無返回鈕
	ClientBackURL string `json:"ClientBackURL,omitempty"`

	// Email 付款人電子信箱 String(50)
	Email string `json:"Email,omitempty"`

	// EmailModify 付款人電子信箱 是否開放修改 Int(1)
	EmailModify bool `json:"EmailModify,omitempty"`

	// OrderComment 商店備註 String(300)
	OrderComment string `json:"OrderComment,omitempty"`

	// CREDIT 信用卡 一次付清啟用 Int(1)
	CREDIT bool `json:"CREDIT,omitempty"`

	// APPLEPAY Apple Pay 啟用 Int(1)
	APPLEPAY bool `json:"APPLEPAY,omitempty"`

	// ANDROIDPAY Google Pay 啟用 Int(1)
	ANDROIDPAY bool `json:"ANDROIDPAY,omitempty"`

	// SAMSUNGPAY Samsung Pay 啟用 Int(1)
	SAMSUNGPAY bool `json:"SAMSUNGPAY,omitempty"`

	// LINEPAY LINE Pay 啟用 Int(1)
	LINEPAY bool `json:"LINEPAY,omitempty"`

	// ImageUrl LINE Pay 產品圖檔連結網址 String(200)
	ImageUrl string `json:"ImageUrl,omitempty"`

	// InstFlag 信用卡 分期付款啟用 String(18)
	//	此欄位值=1 時，即代表開啟所有分期期別，且不可帶入其他期別參數
	//	此欄位值=０或無值時，即代表不開啟分期
	//	接受參數 3 6 8 12 18 24 30
	//	多參數使用 "," 拼接
	InstFlag []InstFlag `json:"InstFlag,omitempty"`

	// CreditRed 信用卡 紅利啟用 Int(1)
	CreditRed bool `json:"CreditRed,omitempty"`

	// UNIONPAY 信用卡 銀聯卡啟用 Int(1)
	UNIONPAY bool `json:"UNIONPAY,omitempty"`

	// CREDITAE 信用卡 美國運通卡啟用 Int(1)
	CREDITAE bool `json:"CREDITAE,omitempty"`

	// WEBATM WEBATM 啟用 Int(1)
	WEBATM bool `json:"WEBATM,omitempty"`

	// VACC ATM 轉帳啟用 Int(1)
	VACC bool `json:"VACC,omitempty"`

	// BankType 金融機構 String(26)
	//	BOT=台灣銀行
	//	HNCB=華南銀行
	//	若未帶值，則預設值為支援所有指定銀行
	//	多參數使用 "," 拼接
	BankType []BankType `json:"BankType,omitempty"`

	// CVS 超商代碼繳費 啟用 Int(1)
	CVS bool `json:"CVS,omitempty"`

	// BARCODE 超商條碼繳費 啟用 Int(1)
	BARCODE bool `json:"BARCODE,omitempty"`

	// ESUNWALLET 玉山 Wallet Int(1)
	ESUNWALLET bool `json:"ESUNWALLET,omitempty"`

	// TAIWANPAY 台灣 Pay Int(1)
	TAIWANPAY bool `json:"TAIWANPAY,omitempty"`

	// BITOPAY BitoPay Int(1)
	BITOPAY bool `json:"BITOPAY,omitempty"`

	// CVSCOM 物流啟用 Int(1)
	CVSCOM bool `json:"CVSCOM,omitempty"`

	// EZPAY 簡單付電子錢包 Int(1)
	EZPAY bool `json:"EZPAY,omitempty"`

	// EZPWECHAT 簡單付微信支付 Int(1)
	EZPWECHAT bool `json:"EZPWECHAT,omitempty"`

	// EZPALIPAY 簡單付支付寶
	EZPALIPAY bool `json:"EZPALIPAY,omitempty"`

	// LgsType 物流型態 String(3)
	//	B2C=大宗寄倉(目前僅支援 7-ELEVEN)
	//	C2C=店到店(支援 7-ELEVEN、全家、萊爾富、OK mart)
	LgsType LgsType `json:"LgsType,omitempty"`

	// OrderDetail 訂單細項 array
	OrderDetail []OrderDetail `json:"OrderDetail,omitempty"`
}

func setFormData(form url.Values, key string, data any) {
	switch val := data.(type) {
	case int:
		if val > 0 {
			form.Set(key, n2s(val))
		}
	case uint:
		if val > 0 {
			form.Set(key, n2s(val))
		}
	case string:
		if val != "" {
			form.Set(key, val)
		}
	case bool:
		if val {
			form.Set(key, "1")
		}
	case time.Time:
		switch key {
		case "TimeStamp":
			form.Set(key, n2s(val.Unix()))
		case "ExpireDate":
			if !val.IsZero() {
				form.Set(key, val.Format(DateOnly))
			}
		}
	case RespondType:
		switch val {
		case RespondTypeString:
			form.Set(key, string(val))
		default:
			form.Set(key, string(RespondTypeJSON))
		}
	case LgsType:
		if val != "" {
			form.Set(key, string(val))
		}
	case LangType:
		if val != "" {
			form.Set(key, string(val))
		}
	case []InstFlag:
		if val != nil && len(val) > 0 {
			instFlag := ""

			for i, flag := range val {
				if flag == InstFlagAll {
					instFlag = n2s(int(flag))
					break
				}

				if i > 0 {
					instFlag += "," + n2s(int(flag))
				} else {
					instFlag = n2s(int(flag))
				}
			}

			if instFlag != "" {
				form.Set(key, instFlag)
			}
		}
	case []BankType:
		if val != nil && len(val) > 0 {
			bankType := ""
			for i, bt := range val {
				if i > 0 {
					bankType += "," + string(bt)
				} else {
					bankType = string(bt)
				}
			}
			if bankType != "" {
				form.Set(key, bankType)
			}
		}
	case []OrderDetail:
		if val != nil && len(val) > 0 {
			marshal, err := json.Marshal(val)
			if err != nil {
				return
			}
			form.Set(key, string(marshal))
		}
	}
}

func (t TradeInfo) aes(m *Merchant) (string, error) {
	form := url.Values{}

	{
		setFormData(form, "MerchantID", m.mid)
		setFormData(form, "RespondType", t.RespondType)
		setFormData(form, "TimeStamp", t.TimeStamp)
		setFormData(form, "Version", Version)
		setFormData(form, "LangType", t.LangType)
		setFormData(form, "MerchantOrderNo", t.MerchantOrderNo)
		setFormData(form, "Amt", t.Amt)
		setFormData(form, "ItemDesc", t.ItemDesc)
		setFormData(form, "TradeLimit", t.TradeLimit)
		setFormData(form, "ExpireDate", t.ExpireDate)
		setFormData(form, "ExpireTime", t.ExpireTime)
		setFormData(form, "ReturnURL", t.ReturnURL)
		setFormData(form, "NotifyURL", t.NotifyURL)
		setFormData(form, "CustomerURL", t.CustomerURL)
		setFormData(form, "ClientBackURL", t.ClientBackURL)
		setFormData(form, "Email", t.Email)
		setFormData(form, "EmailModify", t.EmailModify)
		setFormData(form, "OrderComment", t.OrderComment)
		setFormData(form, "CREDIT", t.CREDIT)
		setFormData(form, "APPLEPAY", t.APPLEPAY)
		setFormData(form, "ANDROIDPAY", t.ANDROIDPAY)
		setFormData(form, "SAMSUNGPAY", t.SAMSUNGPAY)
		setFormData(form, "LINEPAY", t.LINEPAY)
		setFormData(form, "ImageUrl", t.ImageUrl)
		setFormData(form, "InstFlag", t.InstFlag)
		setFormData(form, "CreditRed", t.CreditRed)
		setFormData(form, "UNIONPAY", t.UNIONPAY)
		setFormData(form, "CREDITAE", t.CREDITAE)
		setFormData(form, "WEBATM", t.WEBATM)
		setFormData(form, "VACC", t.VACC)
		setFormData(form, "BankType", t.BankType)
		setFormData(form, "CVS", t.CVS)
		setFormData(form, "BARCODE", t.BARCODE)
		setFormData(form, "ESUNWALLET", t.ESUNWALLET)
		setFormData(form, "TAIWANPAY", t.TAIWANPAY)
		setFormData(form, "BITOPAY", t.BITOPAY)
		setFormData(form, "CVSCOM", t.CVSCOM)
		setFormData(form, "EZPAY", t.EZPAY)
		setFormData(form, "EZPWECHAT", t.EZPWECHAT)
		setFormData(form, "EZPALIPAY", t.EZPALIPAY)
		setFormData(form, "LgsType", t.LgsType)
		setFormData(form, "OrderDetail", t.OrderDetail)
	}

	fmt.Println(form)

	aes := m.getAes()

	encryption, err := aes.Encryption(form.Encode())
	if err != nil {
		return "", err
	}

	return encryption, nil
}

func (t TradeInfo) sha(m *Merchant, aes string) string {
	hashs := fmt.Sprintf("HashKey=%s&%s&HashIV=%s", m.key, aes, m.iv)
	return strings.ToUpper(cryptor.Sha256(hashs))
}

// tradeInfo 交易資料
//type tradeInfo struct {
//	// MerchantID 商店代號 String(15)
//	MerchantID string `validate:"required" json:"MerchantID"`
//
//	// RespondType 回傳格式 String(6)
//	RespondType RespondType `validate:"required,response_type" json:"RespondType,omitempty"`
//
//	// TimeStamp 時間戳記 String(50)
//	//	須確實帶入自 Unix 紀元到當前時間的秒數 以避免交易失敗。(容許誤差值 120 秒)
//	TimeStamp string `validate:"required" json:"TimeStamp"`
//
//	// Version 串接程式版本 String(5)
//	//	2.0
//	Version string `validate:"required" json:"Version"`
//
//	// LangType 語系 String(5)
//	//	英文版 = en
//	//	繁體中文版 = zh-tw
//	//	日文版 = jp
//	LangType string `validate:"lang_type" json:"LangType,omitempty"`
//
//	// MerchantOrderNo 商店訂單編號 String(30)
//	MerchantOrderNo string `validate:"required" json:"MerchantOrderNo"`
//
//	// Amt 訂單金額 Int(10)
//	Amt int `validate:"required" json:"Amt"`
//
//	// ItemDesc 商品資訊 String(50)
//	ItemDesc string `validate:"required" json:"ItemDesc"`
//
//	// TradeLimit 交易有效時間 Int(3)
//	//	秒數下限為 60 秒，小於 60 秒以 60 秒計算
//	//	秒數上限為 900 秒，大於 900 秒以 900 秒計算
//	TradeLimit int `json:"TradeLimit,omitempty"`
//
//	// ExpireDate 繳費有效期限 String(10)
//	//	僅適用於非即時支付 格式為 20060102(ebpay.DateOnly)
//	ExpireDate string `json:"ExpireDate,omitempty"`
//
//	// ExpireTime 繳費截止時間 String(6)
//	//	僅適用於超商代碼繳費 格式為 150405(ebpay.TimeOnly)
//	ExpireTime string `json:"ExpireTime,omitempty"`
//
//	// ReturnURL 支付完成 返回商店網址 String(200)
//	//	交易完成後，以 Form Post 方式導回商店頁面
//	//	只接受 80 與 443 Port
//	ReturnURL string `json:"ReturnURL,omitempty"`
//
//	// NotifyURL 支付通知網址 String(200)
//	//	webhook 網址
//	//	只接受 80 與 443 Port
//	NotifyURL string `json:"NotifyURL,omitempty"`
//
//	// CustomerURL 商店取號網址 String(200)
//	CustomerURL string `json:"CustomerURL,omitempty"`
//
//	// ClientBackURL 返回商店網址 String(200)
//	//	此參數若為空值時，則無返回鈕
//	ClientBackURL string `json:"ClientBackURL,omitempty"`
//
//	// Email 付款人電子信箱 String(50)
//	Email string `json:"Email,omitempty"`
//
//	// EmailModify 付款人電子信箱 是否開放修改 Int(1)
//	EmailModify int `json:"EmailModify,omitempty"`
//
//	// OrderComment 商店備註 String(300)
//	OrderComment string `json:"OrderComment,omitempty"`
//
//	// CREDIT 信用卡 一次付清啟用 Int(1)
//	CREDIT int `json:"CREDIT,omitempty"`
//
//	// APPLEPAY Apple Pay 啟用 Int(1)
//	APPLEPAY int `json:"APPLEPAY,omitempty"`
//
//	// ANDROIDPAY Google Pay 啟用 Int(1)
//	ANDROIDPAY int `json:"ANDROIDPAY,omitempty"`
//
//	// SAMSUNGPAY Samsung Pay 啟用 Int(1)
//	SAMSUNGPAY int `json:"SAMSUNGPAY,omitempty"`
//
//	// LINEPAY LINE Pay 啟用 Int(1)
//	LINEPAY int `json:"LINEPAY,omitempty"`
//
//	// ImageUrl LINE Pay 產品圖檔連結網址 String(200)
//	ImageUrl string `json:"ImageUrl,omitempty"`
//
//	// InstFlag 信用卡 分期付款啟用 String(18)
//	//	此欄位值=1 時，即代表開啟所有分期期別，且不可帶入其他期別參數
//	//	此欄位值=０或無值時，即代表不開啟分期
//	//	接受參數 3 6 8 12 18 24 30
//	//	多參數使用 "," 拼接
//	InstFlag string `json:"InstFlag,omitempty"`
//
//	// CreditRed 信用卡 紅利啟用 Int(1)
//	CreditRed int `json:"CreditRed,omitempty"`
//
//	// UNIONPAY 信用卡 銀聯卡啟用 Int(1)
//	UNIONPAY int `json:"UNIONPAY,omitempty"`
//
//	// CREDITAE 信用卡 美國運通卡啟用 Int(1)
//	CREDITAE int `json:"CREDITAE,omitempty"`
//
//	// WEBATM WEBATM 啟用 Int(1)
//	WEBATM int `json:"WEBATM,omitempty"`
//
//	// VACC ATM 轉帳啟用 Int(1)
//	VACC int `json:"VACC,omitempty"`
//
//	// BankType 金融機構 String(26)
//	//	BOT=台灣銀行
//	//	HNCB=華南銀行
//	//	若未帶值，則預設值為支援所有指定銀行
//	//	多參數使用 "," 拼接
//	BankType string `json:"BankType,omitempty"`
//
//	// CVS 超商代碼繳費 啟用 Int(1)
//	CVS int `json:"CVS,omitempty"`
//
//	// BARCODE 超商條碼繳費 啟用 Int(1)
//	BARCODE int `json:"BARCODE,omitempty"`
//
//	// ESUNWALLET 玉山 Wallet Int(1)
//	ESUNWALLET int `json:"ESUNWALLET,omitempty"`
//
//	// TAIWANPAY 台灣 Pay Int(1)
//	TAIWANPAY int `json:"TAIWANPAY,omitempty"`
//
//	// BITOPAY BitoPay Int(1)
//	BITOPAY int `json:"BITOPAY,omitempty"`
//
//	// CVSCOM 物流啟用 Int(1)
//	CVSCOM int `json:"CVSCOM,omitempty"`
//
//	// EZPAY 簡單付電子錢包 Int(1)
//	EZPAY int `json:"EZPAY,omitempty"`
//
//	// EZPWECHAT 簡單付微信支付 Int(1)
//	EZPWECHAT int `json:"EZPWECHAT,omitempty"`
//
//	// EZPALIPAY 簡單付支付寶
//	EZPALIPAY int `json:"EZPALIPAY,omitempty"`
//
//	// LgsType 物流型態 String(3)
//	//	B2C=大宗寄倉(目前僅支援 7-ELEVEN)
//	//	C2C=店到店(支援 7-ELEVEN、全家、萊爾富、OK mart)
//	LgsType string `json:"LgsType,omitempty"`
//
//	// OrderDetail 訂單細項 array
//	OrderDetail []OrderDetail `json:"OrderDetail,omitempty"`
//}

// number to string
func n2s[T int | int64 | uint](i T) string {
	return fmt.Sprintf("%d", i)
}
