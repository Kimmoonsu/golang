package main

import (
	"fmt"
	"server/autocoin/model/trade"
	"time"
)

// accesskey
// YMu3ew4txZcTZvcDgQllX5YFqkQ3TvNmLiV7A08W

// secretkey
// XmmmMaYgxwxI8wUZdGwNfACIbxn8WZqqlMQVuyqH
const (
	accessKey = "YMu3ew4txZcTZvcDgQllX5YFqkQ3TvNmLiV7A08W"
	secretKey = "XmmmMaYgxwxI8wUZdGwNfACIbxn8WZqqlMQVuyqH"
	marketID  = "KRW-BTC"
)

var exit = make(chan bool)

func main() {
	u := NewUpbit(accessKey, secretKey)
	purchaseOrder(u)
	// go implementStrategy(u)
	// <-exit
	// fmt.Println("Done")
}

func sellOrder(u *Upbit) {
	sellOrder, remaining, e := u.SellOrder(marketID, "0.00160384", "", trade.ORDER_TYPE_MARKET, "")
	if e != nil {
		fmt.Errorf("SellOrder error : %s", e.Error())
	} else {
		fmt.Errorf("SellOrder[remaining:%+v]\n%+v", *remaining, *sellOrder)
	}
}

func purchaseOrder(u *Upbit) {
	purchaseOrder, remaining, e := u.PurchaseOrder(marketID, "", "5000", trade.ORDER_TYPE_PRICE, "")
	if e != nil {
		fmt.Errorf("PurchaseOrder error : %s", e.Error())
	} else {
		fmt.Errorf("PurchaseOrder[remaining:%+v]\n%+v", *remaining, *purchaseOrder)
	}
}

func implementStrategy(u *Upbit) {
	count := 10
	for count > 0 {
		minuteCandles, _, e := u.GetMinuteCandles(marketID, "", "100", "10")
		if e != nil {
			fmt.Errorf("%s's GetMinuteCandles error : %s", marketID, e.Error())
			continue
		}
		slowK, slowD := GetStochastic(minuteCandles)
		fmt.Println("K[0] : ", slowK[0], " / D[0] : ", slowD[0])
		fmt.Println("K[1] : ", slowK[1], " / D[1] : ", slowD[1])
		time.Sleep(time.Second * 1)
		count--
	}
	exit <- true
}
