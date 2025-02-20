package ebpay

import (
	"errors"
	"github.com/gorilla/schema"
	"github.com/wuchieh/aes-go"
	"net/url"
)

var (
	decoder = schema.NewDecoder()
)

type Merchant struct {
	key string
	iv  string
	mid string

	aes *aesgo.AESOptions
}

func (m *Merchant) GetKey() string {
	return m.key
}

func (m *Merchant) GetIv() string {
	return m.iv
}

func (m *Merchant) GetMid() string {
	return m.mid
}

func (m *Merchant) GetAes() aesgo.AESOptions {
	return *m.aes
}

func NewMerchant(HashKey string, HashIV string, MerchantID string) *Merchant {
	return &Merchant{key: HashKey, iv: HashIV, mid: MerchantID}
}

func (m *Merchant) getAes() aesgo.AESOptions {
	if m.aes == nil {
		m.aes = &aesgo.AESOptions{
			Mode:    aesgo.CBC,
			Padding: aesgo.PKCS7Padding,
			Output:  aesgo.Hex,
			Key:     []byte(m.key),
			IV:      []byte(m.iv),
		}
	}

	return *m.aes
}

func (m *Merchant) NewMPGRequest(info TradeInfo) (*MPGRequest, error) {
	aes, err := info.aes(m)
	if err != nil {
		return nil, errors.Join(err, ErrMPGAesEncryption)
	}

	sha := info.sha(m, aes)

	return &MPGRequest{
		MerchantID:  m.mid,
		TradeInfo:   aes,
		TradeSha:    sha,
		Version:     Version,
		EncryptType: 0,
	}, nil
}

func (m *Merchant) ParseReturnData(data []byte) (*ReturnDate, error) {
	query, err := url.ParseQuery(string(data))
	if err != nil {
		return nil, err
	}

	var r returnDate
	err = decoder.Decode(&r, query)
	if err != nil {
		return nil, err
	}

	if !r.check(m) {
		return nil, ErrTradeShaCheckFail
	}

	rData := r.ReturnDate()

	return rData, nil
}
