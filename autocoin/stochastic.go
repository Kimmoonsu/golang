package main

import (
	"fmt"
	"server/autocoin/model/quotation"
)

const (
	K_DAY     = 5
	SLOWK_DAY = 3
	D_DAY     = 3
)

func GetStochastic(candle []*quotation.Candle) ([]float64, []float64) {
	slowK := getSlowK(getK(candle))
	return slowK, getSlowD(slowK)
}

func getK(candle []*quotation.Candle) []float64 {
	K := make([]float64, 50)
	for i := 0; i < 20; i++ {
		max := candle[i].HighPrice
		min := candle[i].LowPrice
		for j := i; j < i+K_DAY; j++ {
			highPrice := candle[j].HighPrice
			lowPrice := candle[j].LowPrice
			if max < highPrice {
				max = highPrice
			}
			if min > lowPrice {
				min = lowPrice
			}
		}

		K[i] = ((candle[i].TradePrice - min) / (max - min)) * 100
		fmt.Println("K : ", K[i])
	}
	return K
}

func getSlowK(K []float64) []float64 {
	slowK := make([]float64, 10)
	for i := 0; i < 10; i++ {
		for j := i; j < i+SLOWK_DAY; j++ {
			slowK[i] += K[j]
		}
		slowK[i] /= SLOWK_DAY
	}
	return slowK
}

func getSlowD(slowK []float64) []float64 {
	slowD := make([]float64, 3)
	for i := 0; i < 3; i++ {
		for j := i; j < i+D_DAY; j++ {
			slowD[i] += slowK[j]
		}
		slowD[i] /= D_DAY
	}
	return slowD
}
