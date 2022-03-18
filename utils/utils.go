package utils

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

func StoI(str string, i int) int {
	o, err := strconv.Atoi(str)
	if err != nil {
		return i
	}
	return o
}

func ItoS(i interface{}) string {
	return fmt.Sprintf("%v", i)
}

func ItoI(i interface{}) int {
	iAreaId, ok := i.(int)
	if ok {
		return iAreaId
	}
	logrus.Infof("SHIT %v %v ", reflect.TypeOf(i), i)
	return 0
}

func StoDec(i string) *decimal.Decimal {
	out, _ := decimal.NewFromString(i)
	return &out
}
