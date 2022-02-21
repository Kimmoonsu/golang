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
	minuteCandles, _, e := u.GetMinuteCandles(marketID, "", "100", "10")
	if e != nil {
		fmt.Errorf("%s's GetMinuteCandles error : %s", marketID, e.Error())
		return
	}
	// fmt.Println("", *remaining)
	// for _, candle := range minuteCandles {
	// 	fmt.Println("high : ", candle.HighPrice)
	// }

	slowK, slowD := GetStochastic(minuteCandles)
	for index, k := range slowK {
		fmt.Println("", index, " : ", k)
	}

	for index, d := range slowD {
		fmt.Println("", index, " : ", d)
	}

}
