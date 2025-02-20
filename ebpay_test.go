package ebpay

import (
	"encoding/json"
	"fmt"
	"github.com/duke-git/lancet/v2/cryptor"
	aesgo "github.com/wuchieh/aes-go"
	"os"
	"testing"
	"time"
)

type Config struct {
	HashKey    string `json:"hash_key"`
	HashIV     string `json:"hash_iv"`
	MerchantID string `json:"merchant_id"`
}

func getConfig() *Config {
	var conf *Config

	file, err := os.ReadFile("config.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(file, &conf)
	if err != nil {
		panic(err)
	}

	return conf
}

func TestPHP(t *testing.T) {
	data1 := "123"
	key := "qwewqeewqwqeqwasdfzxczxcxzaadsd4"
	iv := "asdasdasdasdasda"

	aes := aesgo.AESOptions{
		Mode:    aesgo.CBC,
		Padding: aesgo.PKCS7Padding,
		Output:  aesgo.Hex,
		Key:     []byte(key),
		IV:      []byte(iv),
	}

	edata1, err := aes.Encryption(data1)
	if err != nil {
		t.Fatal(err)
	}

	hashs := fmt.Sprintf("HashKey=%s&%s&HashIV=%s", key, edata1, iv)
	t.Log(hashs)
	t.Log(edata1)
	t.Log(cryptor.Sha256(hashs))
}

func TestMPG(t *testing.T) {
	conf := getConfig()
	m := NewMerchant(conf.HashKey, conf.HashIV, conf.MerchantID)
	info := TradeInfo{
		TimeStamp:       time.Now(),
		MerchantOrderNo: n2s(time.Now().Unix()),
		Amt:             100,
		ItemDesc:        "ItemDesc",
	}

	request, err := m.NewMPGRequest(info)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", request)
}
