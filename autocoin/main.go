package main

import (
	"fmt"
	"time"
)

// accesskey
// YMu3ew4txZcTZvcDgQllX5YFqkQ3TvNmLiV7A08W

// secretkey
// XmmmMaYgxwxI8wUZdGwNfACIbxn8WZqqlMQVuyqH
const (
	accessKey = "YMu3ew4txZcTZvcDgQllX5YFqkQ3TvNmLiV7A08W"
	secretKey = "XmmmMaYgxwxI8wUZdGwNfACIbxn8WZqqlMQVuyqH"
)

var exit = make(chan bool)

func main() {
	u := NewUpbit(accessKey, secretKey)

	go implementStrategy(u)
	<-exit
	fmt.Println("Done")
}

func implementStrategy(u *Upbit) {
	count := 10
	for count > 0 {
		marketID := "KRW-BTC"
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
