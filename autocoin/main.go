package main

import "fmt"

// accesskey
// YMu3ew4txZcTZvcDgQllX5YFqkQ3TvNmLiV7A08W

// secretkey
// XmmmMaYgxwxI8wUZdGwNfACIbxn8WZqqlMQVuyqH
const (
	accessKey = "YMu3ew4txZcTZvcDgQllX5YFqkQ3TvNmLiV7A08W"
	secretKey = "XmmmMaYgxwxI8wUZdGwNfACIbxn8WZqqlMQVuyqH"
)

func main() {
	u := NewUpbit(accessKey, secretKey)
	marketID := "KRW-BTC"
	minuteCandles, remaining, e := u.GetMinuteCandles(marketID, "", "50", "10")
	if e != nil {
		fmt.Errorf("%s's GetMinuteCandles error : %s", marketID, e.Error())
		return
	}
	fmt.Println("GetMinuteCandles[remaining:%+v]", *remaining)
	for _, candle := range minuteCandles {
		fmt.Println("%+v", *candle)
	}
}
