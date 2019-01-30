package main

import (
	"log"
	"net/url"

	"github.com/gorilla/websocket"
)

type JsonRPC2 struct {
	Version string `json:"jsonrpc"`
	Method string `json:"method"`
	Params interface{} `json:"params"`
	Result interface{} `json:"result,omitempty"`
	Id *int `json:"id,omitempty"`
}
type SubscribeParams struct {
	Channel string `json:"channel"`
}

func main() {
	//urlを生成
	u := url.URL{Scheme: "wss", Host: "ws.lightstream.bitflyer.com", Path: "/json-rpc"}
	log.Printf("connecting to %s", u.String())

	//websocketのコネクション先を設定
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()
	//json形式で書き出し
	if err := c.WriteJSON(&JsonRPC2{Version: "2.0", Method: "subscribe", Params: &SubscribeParams{"lightning_ticker_BTC_JPY"}}); err != nil {
		log.Fatal("subscribe:", err)
		return
	}
	//json形式で出力
	for {
		message := new(JsonRPC2)
		if err := c.ReadJSON(message); err != nil {
			log.Println("read:", err)
			return
		}

		if message.Method == "channelMessage" {
			log.Println(message.Params);
		}
	}
}