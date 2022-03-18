package models

import (
	"reflect"
	"testing"
)

func d(t *testing.T, a interface{}) {
	t.Logf("got %v %+v", reflect.TypeOf(a), a)

}

func TestJson1(t *testing.T) {
	var err error
	var response *WsResponse
	responseT1 := `{"result":null,"id":12}`
	responseT2 := ` {"u":7658928,"s":"BTCUSDT","b":"38743.81000000","B":"0.02875000","a":"38743.82000000","A":"0.03936200"}`

	response, err = Parse([]byte(responseT1))
	if err != nil {
		t.Errorf("err %v", err)
	}
	if response.Id != 12 {
		t.Errorf("err %v", err)
	}
	d(t, response)

	response, err = Parse([]byte(responseT2))
	if err != nil {
		t.Errorf("err %v", err)
	}
	d(t, response)
	if response.WsBookData.Id != 7658928 {
		t.Errorf("err %v", err)
	}

}
