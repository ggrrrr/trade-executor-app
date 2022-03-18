package streams

import (
	"encoding/json"
	"fmt"

	"github.com/ggrrrr/trade-executor-app/binance/models"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

func Request(r *models.WsRequest) error {
	logrus.Infof("%+v", r)
	if conn == nil {
		logrus.Errorf("not connected")
		return fmt.Errorf("not connected")
	}

	payload, err := json.Marshal(r)
	if err != nil {
		logrus.Errorf("json: %v", err)
		return err
	}

	wsMsg := wsMsg{
		msgType: websocket.TextMessage,
		payload: payload,
	}

	logrus.Infof("wsMsg1wsMsg1wsMsg1wsMsg1wsMsg1wsMsg1: %+v", wsMsg)
	pushMsg <- &wsMsg
	logrus.Infof("wsMsg2: %+v", wsMsg)

	return nil
}
