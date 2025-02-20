package ebpay

import (
	"errors"
	"github.com/wuchieh/aes-go"
)

type Merchant struct {
	key string
	iv  string
	mid string

	aes *aesgo.AESOptions
}

func (m Merchant) getAes() aesgo.AESOptions {
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

func (m Merchant) NewMPGRequest(info TradeInfo) (*MPGRequest, error) {
	aes, err := info.aes(&m)
	if err != nil {
		return nil, errors.Join(err, ErrMPGAesEncryption)
	}

	sha := info.sha(&m, aes)

	return &MPGRequest{
		MerchantID:  m.mid,
		TradeInfo:   aes,
		TradeSha:    sha,
		Version:     Version,
		EncryptType: 0,
	}, nil
}

func NewMerchant(HashKey string, HashIV string, MerchantID string) *Merchant {
	return &Merchant{key: HashKey, iv: HashIV, mid: MerchantID}
}
