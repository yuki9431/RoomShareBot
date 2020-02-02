package main

import (
	"flag"
	"log"
	"myLinebot/config"
	"net/http"
	"os"

	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/line/line-bot-sdk-go/linebot/httphandler"
	"github.com/yuki9431/logger"
)

const (
	logfile    = "/var/log/linebot.log"
	configFile = "config/config.json"
	mongoDial  = "mongodb://localhost/mongodb"
	mongoName  = "mongodb"
)

// APIIDs API等の設定
type APIIDs struct {
	ChannelSecret string `json:"channelSecret"`
	ChannelToken  string `json:"channelToken"`
	CertFile      string `json:"certFile"`
	KeyFile       string `json:"keyFile"`
}

func main() {
	// log出力設定
	file, err := os.OpenFile(logfile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	logger := logger.New(file)

	// 設定ファイル読み込み
	apiIDs := new(APIIDs)
	config := config.NewConfig(configFile)
	if err := config.Read(apiIDs); err != nil {
		logger.Fatal(err)
	}

	// Setup HTTP Server for receiving requests from LINE platform
	handler, err := httphandler.New(apiIDs.ChannelSecret, apiIDs.ChannelToken)
	if err != nil {
		logger.Fatal(err)
	}

	handler.HandleEvents(func(events []*linebot.Event, r *http.Request) {
		bot, err := handler.NewClient()
		if err != nil {
			logger.Fatal(err)
			return
		}

		// イベント処理
		for _, event := range events {

			logger.Write("start event : " + event.Type)

			// TODO 処理

			// 出費の記録
			// 2人で割った結果を返信

			logger.Write("end event")
		}
	})

	// 使用するポートを取得
	var addr = flag.String("addr", ":443", "アプリケーションのアドレス")
	flag.Parse()

	logger.Write("start server RoomShareBot port", *addr)

	http.Handle("/callback/RoomShareBot", handler)
	if err := http.ListenAndServeTLS(*addr, apiIDs.CertFile, apiIDs.KeyFile, nil); err != nil {
		logger.Fatal("ListenAndServe: ", err)
	}
}
