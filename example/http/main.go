package main

import (
	"encoding/json"
	"fmt"
	"github.com/Wuchieh/newebpay-go"
	"log"
	"net/http"
	"os"
	"strconv"
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

func logJson(data any) {
	if data, err := json.MarshalIndent(data, "", "  "); err == nil {
		log.Println(string(data))
	}
}

func respJson(w http.ResponseWriter, data any) {
	marshal, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}

	logJson(data)

	w.Header().Set("Content-Type", "application/json")
	w.Write(marshal)
}

func main() {
	conf := getConfig()

	m := ebpay.NewMerchant(conf.HashKey, conf.HashIV, conf.MerchantID)
	fmt.Printf("Merchant:%+v\n", m)

	srv := http.NewServeMux()
	srv.HandleFunc("GET /mpg", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")

		mpgRequest, err := m.NewMPGRequest(ebpay.TradeInfo{
			TimeStamp:       time.Now(),
			MerchantOrderNo: strconv.Itoa(int(time.Now().Unix())),
			Amt:             100,
			ItemDesc:        "ItemDesc",
		})

		if err != nil {
			log.Println(err)
			w.WriteHeader(500)
			fmt.Fprint(w, err.Error())
			return
		}

		respJson(w, mpgRequest)
	})

	http.ListenAndServe(":8080", srv)
}
