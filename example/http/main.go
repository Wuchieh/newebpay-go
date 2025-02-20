package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/wuchieh/newebpay-go"
	"io"
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
	ReturnURL  string `json:"return_url"`
	NotifyURL  string `json:"notify_url"`
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

func cors(w http.ResponseWriter, r *http.Request) (abort bool) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	if r.Method == http.MethodOptions {
		log.Println("OPTIONS")
		w.WriteHeader(http.StatusNoContent)
		return true
	}
	return false
}

func saveRequest(r *http.Request, fileName string) error {
	bodyByte, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	headerByte, err := json.MarshalIndent(r.Header, "", "  ")
	if err != nil {
		return err
	}

	n := []byte("\n")

	return os.WriteFile(fileName, bytes.Join([][]byte{headerByte, bodyByte}, n), 0666)
}

func main() {
	conf := getConfig()

	m := ebpay.NewMerchant(conf.HashKey, conf.HashIV, conf.MerchantID)
	fmt.Printf("Merchant:%+v\n", m)

	srv := http.NewServeMux()
	srv.HandleFunc("GET /mpg", func(w http.ResponseWriter, r *http.Request) {
		if cors(w, r) {
			return
		}

		mpgRequest, err := m.NewMPGRequest(ebpay.TradeInfo{
			TimeStamp:       time.Now(),
			MerchantOrderNo: strconv.Itoa(int(time.Now().Unix())),
			Amt:             100,
			ItemDesc:        "ItemDesc",
			ReturnURL:       conf.ReturnURL,
			NotifyURL:       conf.NotifyURL,
		})

		if err != nil {
			log.Println(err)
			w.WriteHeader(500)
			fmt.Fprint(w, err.Error())
			return
		}

		respJson(w, mpgRequest)
	})

	srv.HandleFunc("/return", func(w http.ResponseWriter, r *http.Request) {
		if cors(w, r) {
			return
		}

		if err := saveRequest(r, fmt.Sprintf("return_%d.txt", time.Now().Unix())); err != nil {
			log.Println(err)
			return
		}

		w.WriteHeader(http.StatusOK)
	})

	srv.HandleFunc("/notify", func(w http.ResponseWriter, r *http.Request) {
		if cors(w, r) {
			return
		}

		bodyByte, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
			w.WriteHeader(500)
			fmt.Fprint(w, err.Error())
			return
		}

		if err := os.WriteFile(fmt.Sprintf("notify_%d.txt", time.Now().Unix()), bodyByte, 0666); err != nil {
			log.Println(err)
		}

		data, err := m.ParseReturnData(bodyByte)
		if err != nil {
			log.Println(err)
			return
		}

		log.Printf("%+v\n", data)

		info, err := data.GetTradeInfo(m)
		if err != nil {
			log.Println(err)
			return
		}

		fmt.Printf("%#v\n", info)

		w.WriteHeader(http.StatusOK)
	})

	http.ListenAndServe(":8080", srv)
}
