package streams

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"

	"github.com/ggrrrr/trade-executor-app/binance/models"
	"github.com/ggrrrr/trade-executor-app/controller"
	"github.com/ggrrrr/trade-executor-app/utils"
)

const (
	H_API_KEY    = "binance-api-key"
	H_API_SECRET = "binance-api-secret"
)

type wsMsg struct {
	msgType int
	payload []byte
}

var (
	url      string
	pushMsg  chan *wsMsg
	conn     *websocket.Conn
	run      = true
	requests = sync.WaitGroup{}
	done     chan struct{}
	// url: "wss://testnet.binance.vision/ws",
	// url: "wss://testnet.binance.vision/ws/btcusdt@trade",
	// url: "wss://testnet.binance.vision/ws/BTCUSDT@bookTicker",
	// url: "wss://testnet.binance.vision/ws/BTCUSDT@bookTicker",
	// url: "wss://testnet-dex.binance.org/api/ws/btcusdt@bookTicker",
	// }
)

func Config(shutdown chan struct{}, tr *controller.MarketOrder) error {
	var err error
	url = strings.TrimSpace(utils.GetString("binance", "ws.url"))
	if url == "" {
		logrus.Errorf("url is empty")
		return fmt.Errorf("url is empty")
	}

	// interrupt := make(chan os.Signal, 1)
	// signal.Notify(interrupt, os.Interrupt)
	logrus.Infof("url: %s", url)
	conn, _, err = websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		logrus.Errorf("connect %v %v", url, err)
		return fmt.Errorf("unable to connect: %s %v", url, err)
	}

	done = shutdown
	return nil
}

func Start(bookData chan *models.WsBookData) {
	if conn == nil {
		logrus.Error("ws conn is nil")
		panic("ws conn is nil")
	}
	defer conn.Close()
	pushMsg = make(chan *wsMsg)
	defer close(pushMsg)
	go func() {
		for run {
			msgType, msg, err := conn.ReadMessage()
			if err != nil {
				logrus.Errorf("read: %v", err)
				return
			}
			switch msgType {
			case websocket.BinaryMessage:
				logrus.Errorf("recv: BinaryMessage: %s", msg)
			case websocket.TextMessage:
				logrus.Debugf("recv: txt raw: %s", msg)
				bd, err := models.Parse(msg)
				if err != nil {
					logrus.Warnf("unkown data: %v %s", err, msg)
					break
				}
				bookData1, err := models.IsBookData(bd)
				if err != nil {
					logrus.Warnf("unkown response: %v %s", err, msg)
					break
				}
				bookData <- bookData1
			case websocket.PingMessage:
				logrus.Debugf("recv: PingMessage: %s", msg)
				// TODO send pong message
			case websocket.PongMessage:
				logrus.Debugf("recv: PongMessage: %s", msg)
				// TODO
			default:
				logrus.Errorf("recv[%d]: %s", msgType, msg)
			}
		}
	}()
	logrus.Info("started")
	for run {
		select {
		case <-done:
			run = false
			logrus.Infof("interrupted")
			err := conn.WriteMessage(websocket.CloseMessage,
				websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				logrus.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		case m := <-pushMsg:
			logrus.Debugf("WriteMessage[%d]: %+s", m.msgType, string(m.payload))
			err := conn.WriteMessage(m.msgType, m.payload)
			if err != nil {
				logrus.Errorf("unable to send message: %s, %v", m.payload, err)
				return
			}
			requests.Add(1)
		}
	}
}
